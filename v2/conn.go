// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"io"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
)

// Conn represents a GTPv2-C connection.
//
// Conn provides the automatic handling of message by adding handlers to it with
// AddHandler(s). See AddHandler for detailed usage.
//
// Conn also provides the functions to manage Sessions/Bearers that works over the
// connection(=between a node to another).
// See the docs of CreateSession, AddSession, DeleteSession methods for details.
type Conn struct {
	mu      sync.Mutex
	laddr   net.Addr
	pktConn net.PacketConn
	ifTypes []uint8
	*imsiSessionMap
	*teidSessionMap

	validationEnabled bool

	closeCh chan struct{}
	*msgHandlerMap

	// sequence is the last SequenceNumber used in the request.
	//
	// TS29.274 7.6  Reliable Delivery of Signalling Messages;
	// The Sequence Number shall be unique for each outstanding Initial message sourced
	// from the same IP/UDP endpoint(=Conn).
	sequence uint32

	// RestartCounter is the RestartCounter value in Recovery IE, which represents how many
	// times the GTPv2-C endpoint is restarted.
	RestartCounter uint8
}

// NewConn creates a new Conn used for server. On client side, use Dial instead.
func NewConn(laddr net.Addr, counter uint8, ifTypes ...uint8) *Conn {
	return &Conn{
		mu:                sync.Mutex{},
		laddr:             laddr,
		ifTypes:           ifTypes,
		imsiSessionMap:    newimsiSessionMap(),
		teidSessionMap:    newteidSessionMap(),
		validationEnabled: true,
		closeCh:           make(chan struct{}),
		msgHandlerMap:     defaultHandlerMap,
		sequence:          0,
		RestartCounter:    counter,
	}
}

// Dial sends Echo Request to raddr to check if the endpoint is alive and returns Conn.
//
// It does not bind the raddr to the underlying connection, which enables a Conn to
// send to/receive from multiple peers with single laddr.
//
// If Echo exchange is unnecessary, use NewConn and ListenAndServe instead.
func Dial(ctx context.Context, laddr, raddr net.Addr, counter uint8, ifTypes ...uint8) (*Conn, error) {
	c := &Conn{
		mu:                sync.Mutex{},
		ifTypes:           ifTypes,
		imsiSessionMap:    newimsiSessionMap(),
		teidSessionMap:    newteidSessionMap(),
		validationEnabled: true,
		closeCh:           make(chan struct{}),
		msgHandlerMap:     defaultHandlerMap,
		sequence:          0,
		RestartCounter:    counter,
	}

	// setup underlying connection first.
	// not using net.Dial, as it binds src/dst IP:Port, which makes it harder to
	// handle multiple connections with a Conn.
	var err error
	c.pktConn, err = net.ListenPacket(raddr.Network(), laddr.String())
	if err != nil {
		return nil, err
	}

	// send EchoRequest to raddr.
	if _, err := c.EchoRequest(raddr); err != nil {
		return nil, err
	}

	buf := make([]byte, 1600)

	// if no response coming within 3 seconds, returns error without retrying.
	if err := c.pktConn.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
		return nil, err
	}
	n, raddr, err := c.pktConn.ReadFrom(buf)
	if err != nil {
		return nil, err
	}
	if err := c.pktConn.SetReadDeadline(time.Time{}); err != nil {
		return nil, err
	}

	// decode incoming message and let it be handled by default handler funcs.
	msg, err := messages.Parse(buf[:n])
	if err != nil {
		return nil, err
	}
	if err := c.handleMessage(raddr, msg); err != nil {
		return nil, err
	}

	go func() {
		if err := c.serve(ctx); err != nil {
			logf("fatal error on Conn %s: %s", c.LocalAddr(), err)
		}
	}()
	return c, nil
}

// ListenAndServe creates a new GTPv2-C Conn and start serving background.
func (c *Conn) ListenAndServe(ctx context.Context) error {
	var err error
	c.pktConn, err = net.ListenPacket(c.laddr.Network(), c.laddr.String())
	if err != nil {
		return err
	}

	return c.listenAndServe(ctx)
}

func (c *Conn) listenAndServe(ctx context.Context) error {
	// TODO: this func is left for future enhancement.
	return c.serve(ctx)
}

func (c *Conn) closed() <-chan struct{} {
	return c.closeCh
}

func (c *Conn) serve(ctx context.Context) error {
	buf := make([]byte, 1500)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-c.closed():
			return nil
		default:
			// do nothing and go forward.
		}

		n, raddr, err := c.pktConn.ReadFrom(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return errors.Errorf("error reading from Conn %s: %s", c.LocalAddr(), err)
		}

		raw := make([]byte, n)
		copy(raw, buf)
		go func() {
			msg, err := messages.Parse(raw)
			if err != nil {
				logf("error parsing the message: %v, %x", err, raw)
				return
			}

			if err := c.handleMessage(raddr, msg); err != nil {
				logf("error handling message on Conn %s: %s", c.LocalAddr(), err)
			}
		}()
	}
}

// ReadFrom reads a packet from the connection,
// copying the payload into p. It returns the number of
// bytes copied into p and the return address that
// was on the packet.
// It returns the number of bytes read (0 <= n <= len(p))
// and any error encountered. Callers should always process
// the n > 0 bytes returned before considering the error err.
// ReadFrom can be made to time out and return
// an Error with Timeout() == true after a fixed time limit;
// see SetDeadline and SetReadDeadline.
func (c *Conn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	return c.pktConn.ReadFrom(p)
}

// WriteTo writes a packet with payload p to addr.
// WriteTo can be made to time out and return
// an Error with Timeout() == true after a fixed time limit;
// see SetDeadline and SetWriteDeadline.
// On packet-oriented connections, write timeouts are rare.
func (c *Conn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	return c.pktConn.WriteTo(p, addr)
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (c *Conn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.msgHandlerMap = defaultHandlerMap
	c.RestartCounter = 0
	close(c.closeCh)

	if err := c.pktConn.Close(); err != nil {
		logf("error closing the underlying conn: %s", err)
	}
	return nil
}

// LocalAddr returns the local network address.
func (c *Conn) LocalAddr() net.Addr {
	return c.pktConn.LocalAddr()
}

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail with a timeout (see type Error) instead of
// blocking. The deadline applies to all future and pending
// I/O, not just the immediately following call to Read or
// Write. After a deadline has been exceeded, the connection
// can be refreshed by setting a deadline in the future.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
func (c *Conn) SetDeadline(t time.Time) error {
	return c.pktConn.SetDeadline(t)
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.pktConn.SetReadDeadline(t)
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.pktConn.SetWriteDeadline(t)
}

// AddHandler adds a message handler to Conn.
//
// By adding HandlerFunc, Conn(and Session, Bearer created over the Conn) will handle
// the specified type of message with it's paired HandlerFunc when receiving.
// Messages without registered handlers are just ignored and logged.
//
// This should be performed just after creating Conn, otherwise the user cannot retrieve
// any values, which is in most cases vital to continue working as a node, from the incoming
// messages.
//
// The error returned from handler is just logged. Any important due should be done inside
// the HandlerFunc before returning. This behavior might change in the future.
//
// HandlerFunc for EchoResponse and VersionNotSupportedIndication are registered by default.
// These HandlerFunc can be overridden by specifying messages.MsgTypeEchoResponse and/or
// messages.MsgTypeVersionNotSupportedIndication as msgType parameter.
func (c *Conn) AddHandler(msgType uint8, fn HandlerFunc) {
	c.msgHandlerMap.store(msgType, fn)
}

// AddHandlers adds multiple handler funcs at a time, using a map.
// The key of the map is message type of the GTPv2-C messages. You can use MsgTypeFooBar
// constants defined in this package as well as any raw uint8 values.
//
// See AddHandler for how the given handlers behave.
func (c *Conn) AddHandlers(funcs map[uint8]HandlerFunc) {
	for msgType, fn := range funcs {
		c.msgHandlerMap.store(msgType, fn)
	}
}

func (c *Conn) handleMessage(senderAddr net.Addr, msg messages.Message) error {
	if c.validationEnabled {
		if err := c.validate(senderAddr, msg); err != nil {
			logf("failed to validate a message: %s", err)
		}
	}

	handle, ok := c.msgHandlerMap.load(msg.MessageType())
	if !ok {
		logf("%v", &HandlerNotFoundError{MsgType: msg.MessageTypeName()})
	}
	if err := handle(c, senderAddr, msg); err != nil {
		logf("failed to handle message %s: %s", msg, err)
	}

	return nil
}

// EnableValidation turns on automatic validation of incoming messages.
// This is expected to be used only after DisableValidation() is used, as the validation
// is enabled by default.
//
// Conn checks if;
//
//  GTP Version is 2
//  TEID is known to Conn
//
// Even the validation is failed, it does not return error to user. Instead, it just logs
// and discards the packets so that the HandlerFunc won't get the invalid message.
// Extra validations should be done in HandlerFunc.
func (c *Conn) EnableValidation() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.validationEnabled = true
}

// DisableValidation turns off automatic validation of incoming messages.
// It is not recommended to use this except the node is in debugging mode.
//
// See EnableValidation for what are validated.
func (c *Conn) DisableValidation() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.validationEnabled = false
}

func (c *Conn) validate(senderAddr net.Addr, msg messages.Message) error {
	// check GTP version
	if msg.Version() != 2 {
		if err := c.VersionNotSupportedIndication(senderAddr, msg); err != nil {
			return errors.Errorf("failed to respond with VersionNotSupportedIndication: %s", err)
		}
		return errors.Errorf("received an invalid version(%d) of message: %v", msg.Version(), msg)
	}

	// check if TEID is known or not
	if teid := msg.TEID(); teid != 0 {
		if _, err := c.GetSessionByTEID(teid, senderAddr); err != nil {
			return &InvalidTEIDError{TEID: teid}
		}
	}
	return nil
}

// SendMessageTo sends a message to addr.
// Unlike WriteTo, it sets the Sequence Number properly and returns the one used in the message.
func (c *Conn) SendMessageTo(msg messages.Message, addr net.Addr) (uint32, error) {
	seq := c.IncSequence()
	msg.SetSequenceNumber(seq)

	payload, err := messages.Marshal(msg)
	if err != nil {
		seq = c.DecSequence()
		return seq, errors.Wrapf(err, "failed to send %T", msg)
	}

	if _, err := c.WriteTo(payload, addr); err != nil {
		seq = c.DecSequence()
		return seq, errors.Wrapf(err, "failed to send %T", msg)
	}
	return seq, nil
}

// IncSequence increments the SequenceNumber associated with Conn.
func (c *Conn) IncSequence() uint32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sequence++

	// SequenceNumber is 3-octet long
	if c.sequence > 0xffffff {
		c.sequence = 0
	}

	return c.sequence
}

// DecSequence decrements the SequenceNumber associated with Conn.
func (c *Conn) DecSequence() uint32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sequence--

	return c.sequence
}

// SequenceNumber returns the current(=last used) SequenceNumber associated with Conn.
func (c *Conn) SequenceNumber() uint32 {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.sequence
}

// EchoRequest sends a EchoRequest.
func (c *Conn) EchoRequest(raddr net.Addr) (uint32, error) {
	msg := messages.NewEchoRequest(0, ies.NewRecovery(c.RestartCounter))

	seq, err := c.SendMessageTo(msg, raddr)
	if err != nil {
		return 0, err
	}
	return seq, nil
}

// EchoResponse sends a EchoResponse in response to the EchoRequest.
func (c *Conn) EchoResponse(raddr net.Addr, req messages.Message) error {
	res := messages.NewEchoResponse(0, ies.NewRecovery(c.RestartCounter))

	if err := c.RespondTo(raddr, req, res); err != nil {
		return err
	}
	return nil
}

// VersionNotSupportedIndication sends VersionNotSupportedIndication message
// in response to any kind of message.Message.
func (c *Conn) VersionNotSupportedIndication(raddr net.Addr, req messages.Message) error {
	res := messages.NewVersionNotSupportedIndication(0, req.Sequence())

	if err := c.RespondTo(raddr, req, res); err != nil {
		return err
	}
	return nil
}

// CreateSession sends a CreateSessionRequest and stores information given with IE
// in the Session returned.
//
// After using this method, users don't need to call AddSession with the session
// returned.
//
// By creating a Session with this method, the values in IEs given, such as TEID in F-TEID
// are stored with "best effort". See the source code to see what kind information is
// handled automatically in this method.
//
// Also, a Bearer named "default" is also created to be used as default bearer.
// The default bearer can be retrieved by using GetDefaultBearer() or LookupBearerByName("default").
//
// Note that this method doesn't care IEs given are sufficient or not, as the required IE
// varies much depending on the context in which the Create Session Request is used.
// In other words, any kind of IE can be put on the Create Session Request message using
// this method.
func (c *Conn) CreateSession(raddr net.Addr, ie ...*ies.IE) (*Session, uint32, error) {
	// retrieve values from IEs given.
	sess := NewSession(raddr, &Subscriber{Location: &Location{}})
	br := sess.GetDefaultBearer()
	var err error
	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			sess.IMSI, err = i.IMSI()
			if err != nil {
				return nil, 0, err
			}
		case ies.MSISDN:
			sess.MSISDN, err = i.MSISDN()
			if err != nil {
				return nil, 0, err
			}
		case ies.MobileEquipmentIdentity:
			sess.IMEI, err = i.MobileEquipmentIdentity()
			if err != nil {
				return nil, 0, err
			}
		case ies.ServingNetwork:
			sess.MCC, err = i.MCC()
			if err != nil {
				return nil, 0, err
			}
			sess.MNC, err = i.MNC()
			if err != nil {
				return nil, 0, err
			}
		case ies.AccessPointName:
			br.APN, err = i.AccessPointName()
			if err != nil {
				return nil, 0, err
			}
		case ies.RATType:
			sess.RATType, err = i.RATType()
			if err != nil {
				return nil, 0, err
			}
		case ies.FullyQualifiedTEID:
			it, err := i.InterfaceType()
			if err != nil {
				return nil, 0, err
			}
			teid, err := i.TEID()
			if err != nil {
				return nil, 0, err
			}
			sess.AddTEID(it, teid)
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				for _, child := range i.ChildIEs {
					switch child.Type {
					case ies.EPSBearerID:
						br.EBI, err = child.EPSBearerID()
						if err != nil {
							return nil, 0, err
						}
					case ies.BearerQoS:
						br.PL, err = child.PriorityLevel()
						if err != nil {
							return nil, 0, err
						}
						br.QCI, err = child.QCILabel()
						if err != nil {
							return nil, 0, err
						}
						br.PCI = child.PreemptionCapability()
						br.PVI = child.PreemptionVulnerability()

						br.MBRUL, err = child.MBRForUplink()
						if err != nil {
							return nil, 0, err
						}
						br.MBRDL, err = child.MBRForDownlink()
						if err != nil {
							return nil, 0, err
						}
						br.GBRUL, err = child.GBRForUplink()
						if err != nil {
							return nil, 0, err
						}
						br.GBRDL, err = child.GBRForUplink()
						if err != nil {
							return nil, 0, err
						}
					case ies.FullyQualifiedTEID:
						it, err := child.InterfaceType()
						if err != nil {
							return nil, 0, err
						}
						teid, err := child.TEID()
						if err != nil {
							return nil, 0, err
						}
						sess.AddTEID(it, teid)
					case ies.BearerTFT:
						// XXX - do nothing for BearerTFT?
					}
				}
			case 1:
				// XXX - do nothing for BearerContextsToBeRemoved?
			}
		}
	}

	// add session here to registry in case of the response comes super fast.
	c.AddSession(sess)

	// set IEs into CreateSessionRequest.
	msg := messages.NewCreateSessionRequest(0, 0, ie...)

	seq, err := c.SendMessageTo(msg, raddr)
	if err != nil {
		return nil, 0, err
	}
	return sess, seq, nil
}

// DeleteSession sends a DeleteSessionRequest with TEID and IEs given.
func (c *Conn) DeleteSession(teid uint32, raddr net.Addr, ie ...*ies.IE) (uint32, error) {
	sess, err := c.GetSessionByTEID(teid, raddr)
	if err != nil {
		return 0, err
	}

	msg := messages.NewDeleteSessionRequest(teid, 0, ie...)

	seq, err := c.SendMessageTo(msg, sess.peerAddr)
	if err != nil {
		return 0, err
	}
	return seq, nil
}

// ModifyBearer sends a ModifyBearerRequest with TEID and IEs given..
func (c *Conn) ModifyBearer(teid uint32, raddr net.Addr, ie ...*ies.IE) (uint32, error) {
	sess, err := c.GetSessionByTEID(teid, raddr)
	if err != nil {
		return 0, err
	}

	msg := messages.NewModifyBearerRequest(teid, 0, ie...)

	seq, err := c.SendMessageTo(msg, sess.peerAddr)
	if err != nil {
		return 0, err
	}
	return seq, nil
}

// DeleteBearer sends a DeleteBearerRequest TEID and with IEs given.
func (c *Conn) DeleteBearer(teid uint32, raddr net.Addr, ie ...*ies.IE) (uint32, error) {
	sess, err := c.GetSessionByTEID(teid, raddr)
	if err != nil {
		return 0, err
	}

	msg := messages.NewDeleteBearerRequest(teid, 0, ie...)

	seq, err := c.SendMessageTo(msg, sess.peerAddr)
	if err != nil {
		return 0, err
	}
	return seq, nil
}

// RespondTo sends a message(specified with "toBeSent" param) in response to a message
// (specified with "received" param).
//
// This exists to make it easier to handle SequenceNumber.
func (c *Conn) RespondTo(raddr net.Addr, received, toBeSent messages.Message) error {
	toBeSent.SetSequenceNumber(received.Sequence())
	b := make([]byte, toBeSent.MarshalLen())

	if err := toBeSent.MarshalTo(b); err != nil {
		return err
	}

	if _, err := c.WriteTo(b, raddr); err != nil {
		return err
	}
	return nil
}

// GetSessionByTEID returns Session looked up by TEID and sender of the message.
func (c *Conn) GetSessionByTEID(teid uint32, peer net.Addr) (*Session, error) {
	session, ok := c.teidSessionMap.load(teid)
	if !ok {
		return nil, &InvalidTEIDError{TEID: teid}
	}
	if peer.String() != session.peerAddrString {
		return nil, &InvalidTEIDError{TEID: teid}
	}
	return session, nil
}

// GetSessionByIMSI returns Session looked up by IMSI.
func (c *Conn) GetSessionByIMSI(imsi string) (*Session, error) {
	if session, ok := c.imsiSessionMap.load(imsi); ok {
		return session, nil
	}
	return nil, &UnknownIMSIError{IMSI: imsi}
}

// GetIMSIByTEID returns IMSI associated with TEID and the peer node.
func (c *Conn) GetIMSIByTEID(teid uint32, peer net.Addr) (string, error) {
	sess, err := c.GetSessionByTEID(teid, peer)
	if err != nil {
		return "", err
	}

	return sess.IMSI, nil
}

// AddSession adds a session imsiSessionMap
func (c *Conn) AddSession(session *Session) {
	c.imsiSessionMap.store(session.IMSI, session)
	for _, ifType := range c.ifTypes {
		teid, ok := session.teidMap.load(ifType)
		if !ok {
			logf("Error: Missing TEID %d for IFType %d in session for IMSI %s",
				teid, ifType, session.IMSI)
			continue
		}
		c.teidSessionMap.store(teid, session)
	}
}

// RemoveSession removes a session from c.Session.
// The Session is identified by IMSI.
func (c *Conn) RemoveSession(session *Session) {
	c.imsiSessionMap.delete(session.IMSI)
	for _, ifType := range c.ifTypes {
		teid, ok := session.teidMap.load(ifType)
		if !ok {
			logf("Error: Missing TEID %d for IFType %d in session for IMSI %s",
				teid, ifType, session.IMSI)
			continue
		}
		c.teidSessionMap.delete(teid)
	}
}

// RemoveSessionByIMSI removes a session looked up by IMSI.
func (c *Conn) RemoveSessionByIMSI(imsi string) {
	sess, ok := c.imsiSessionMap.load(imsi)
	if !ok {
		return
	}
	c.RemoveSession(sess)
}

// NewFTEID creates a new F-TEID with random TEID value that is unique within Conn.
// To ensure the uniqueness, don't create in the other way if you once use this method.
//
// Note that in the case there's a lot of Session on the Conn, it may take a long
// time to find a new unique value.
//
// TODO: optimize performance...
func (c *Conn) NewFTEID(ifType uint8, v4, v6 string) (fteidIE *ies.IE) {
	var teid uint32
	for try := uint32(0); try < 0xffff; try++ {
		const logEvery = 0xff
		if try&logEvery == logEvery {
			logf("Generating NewFTEID crossed tries:%d", try)
		}

		t := generateRandomUint32()
		if t == 0 {
			continue
		}

		// Try to mark TEID as taken. Fails if something exists
		if ok := c.teidSessionMap.tryStore(t, nil); !ok {
			continue
		}

		teid = t
		break
	}

	if teid == 0 {
		return nil
	}
	return ies.NewFullyQualifiedTEID(ifType, teid, v4, v6)
}

func generateRandomUint32() uint32 {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return 0
	}

	return binary.BigEndian.Uint32(b)
}

// Bearers returns all the bearers registered in Session.
func (c *Conn) Sessions() []*Session {
	var ss []*Session
	c.imsiSessionMap.rangeWithFunc(func(k, v interface{}) bool {
		ss = append(ss, v.(*Session))
		return true
	})

	return ss
}

// SessionCount returns the number of sessions registered in Conn.
//
// This may have some impact on performance in case of large number of Session exists.
func (c *Conn) SessionCount() int {
	var count int
	c.imsiSessionMap.rangeWithFunc(func(k, v interface{}) bool {
		sess := v.(*Session)
		if sess.IsActive() {
			count++
		}
		return true
	})

	return count
}

// BearerCount returns the number of bearers registered in Conn.
//
// This may have some impact on performance in case of large number of Session and Bearer exist.
func (c *Conn) BearerCount() int {
	var count int
	c.imsiSessionMap.rangeWithFunc(func(k, v interface{}) bool {
		sess := v.(*Session)
		if sess.IsActive() {
			count += sess.BearerCount()
		}
		return true
	})

	return count
}

type imsiSessionMap struct {
	syncMap sync.Map
}

func newimsiSessionMap() *imsiSessionMap {
	return &imsiSessionMap{}
}

func (i *imsiSessionMap) store(imsi string, session *Session) {
	i.syncMap.Store(imsi, session)
}

func (i *imsiSessionMap) load(imsi string) (*Session, bool) {
	session, ok := i.syncMap.Load(imsi)
	if ok && session != nil {
		return session.(*Session), true
	}
	return nil, ok
}

func (i *imsiSessionMap) delete(imsi string) {
	i.syncMap.Delete(imsi)
}

func (i *imsiSessionMap) rangeWithFunc(fn func(imsi, session interface{}) bool) {
	i.syncMap.Range(fn)
}

type teidSessionMap struct {
	syncMap sync.Map
}

func newteidSessionMap() *teidSessionMap {
	return &teidSessionMap{}
}

func (t *teidSessionMap) store(teid uint32, session *Session) {
	t.syncMap.Store(teid, session)
}

func (t *teidSessionMap) tryStore(teid uint32, session *Session) bool {
	_, loaded := t.syncMap.LoadOrStore(teid, session)
	return !loaded
}

func (t *teidSessionMap) load(teid uint32) (*Session, bool) {
	session, ok := t.syncMap.Load(teid)
	if ok && session != nil {
		return session.(*Session), true
	}
	return nil, ok
}

func (t *teidSessionMap) delete(teid uint32) {
	t.syncMap.Delete(teid)
}
