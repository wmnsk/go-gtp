// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import (
	"crypto/rand"
	"encoding/binary"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
)

// Conn is a GTPv2-C connection.
type Conn struct {
	mu      sync.Mutex
	pktConn net.PacketConn

	validationEnabled bool

	closeCh chan struct{}
	errCh   chan error

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

	// Sessions is a set of sessions exists on the Conn with automatically-assigned IDs.
	Sessions []*Session
}

// NewConn creates a new Conn over existing net.PacketConn.
//
// This is for special situation that the user already have a net.PacketConn to be used for
// GTPv2-C connection. Otherwise, Dial() or ListenAndServe() should be used to create a Conn.
func NewConn(pktConn net.PacketConn, raddr net.Addr, counter uint8, errCh chan error) (*Conn, error) {
	c := &Conn{
		mu:                sync.Mutex{},
		pktConn:           pktConn,
		validationEnabled: true,
		closeCh:           make(chan struct{}),
		errCh:             errCh,
		msgHandlerMap:     defaultHandlerMap,
		sequence:          0,
		RestartCounter:    counter,
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
	msg, err := messages.Decode(buf[:n])
	if err != nil {
		return nil, err
	}
	if err := c.handleMessage(raddr, msg); err != nil {
		return nil, err
	}

	go c.serve()
	return c, nil
}

// Dial just exchanges the GTPv2 Echo and returns *Conn.
//
// Dial does not actually Dial() remote address so that the *Conn can be used with
// multiple source/destination address.
// The difference between Dial() and ListenAndServe() is just a presence of Echo
// exchange before returning *Conn.
//
// The hbInfo is *HeartBeatinfo. If nil is given, heartbeat will be disabled.
//
// The errCh given should be monitored continuously after retrieving *Conn.
// Otherwise the background process may get stuck.
func Dial(laddr, raddr net.Addr, counter uint8, errCh chan error) (*Conn, error) {
	c := &Conn{
		mu:                sync.Mutex{},
		validationEnabled: true,
		closeCh:           make(chan struct{}),
		errCh:             errCh,
		msgHandlerMap:     defaultHandlerMap,
		sequence:          0,
		RestartCounter:    counter,
	}

	// setup underlying connection first.
	// don't use Dial(), as it binds src/dst IP:Port and it makes it harder to
	// handle multiple connections.
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
	msg, err := messages.Decode(buf[:n])
	if err != nil {
		return nil, err
	}
	if err := c.handleMessage(raddr, msg); err != nil {
		return nil, err
	}

	go c.serve()
	return c, nil
}

// ListenAndServe creates a new GTPv2-C *Conn and start serving.
//
// The errCh given should be monitored continuously after retrieving *Conn.
// Otherwise the background process may get stuck.
func ListenAndServe(laddr net.Addr, counter uint8, errCh chan error) (*Conn, error) {
	c := &Conn{
		mu:                sync.Mutex{},
		validationEnabled: true,
		closeCh:           make(chan struct{}),
		errCh:             errCh,
		msgHandlerMap:     defaultHandlerMap,
		sequence:          0,
		RestartCounter:    counter,
	}

	var err error
	c.pktConn, err = net.ListenPacket(laddr.Network(), laddr.String())
	if err != nil {
		return nil, err
	}

	go c.serve()
	return c, nil
}

func (c *Conn) closed() <-chan struct{} {
	return c.closeCh
}

func (c *Conn) serve() {
	buf := make([]byte, 1600)
	for {
		select {
		case <-c.closed():
			return
		default:
			// do nothing and go forward.
		}

		n, raddr, err := c.pktConn.ReadFrom(buf)
		if err != nil {
			continue
		}

		raw := make([]byte, n)
		copy(raw, buf)
		go func() {
			msg, err := messages.Decode(raw)
			if err != nil {
				return
			}

			if err := c.handleMessage(raddr, msg); err != nil {
				c.errCh <- err
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

	// triggers error in blocking Read() / Write() immediately.
	if err := c.pktConn.SetDeadline(time.Now().Add(1 * time.Millisecond)); err != nil {
		return err
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

// AddHandler adds a message handler to *Conn.
//
// By adding HandlerFuncs, *Conn (and *Session, *Bearer created by the *Conn) will handle
// the specified type of message with it's paired HandlerFunc when receiving.
// Messages without registered handlers are just ignored and discarded and the user will
// get ErrNoHandlersFound error.
//
// This should be performed just after creating *Conn, otherwise the user cannot retrieve
// any values, which is in most cases vital to continue working as a node, from the incoming
// messages.
//
// HandlerFuncs for EchoResponse and VersionNotSupportedIndication are registered by default.
// These HandlerFuncs can be overwritten by specifying messages.MsgTypeEchoResponse and/or
// messages.MsgTypeVersionNotSupportedIndication as msgType parameter.
func (c *Conn) AddHandler(msgType uint8, fn HandlerFunc) {
	c.msgHandlerMap.store(msgType, fn)
}

// AddHandlers adds multiple handler funcs at a time.
//
// See AddHandler for detailed usage.
func (c *Conn) AddHandlers(funcs map[uint8]HandlerFunc) {
	for msgType, fn := range funcs {
		c.msgHandlerMap.store(msgType, fn)
	}
}

func (c *Conn) handleMessage(senderAddr net.Addr, msg messages.Message) error {
	if c.validationEnabled {
		if err := c.validate(senderAddr, msg); err != nil {
			return err
		}
	}

	handle, ok := c.msgHandlerMap.load(msg.MessageType())
	if !ok {
		return &ErrNoHandlersFound{MsgType: msg.MessageTypeName()}
	}
	if err := handle(c, senderAddr, msg); err != nil {
		c.errCh <- err
	}

	return nil
}

// EnableValidation turns on automatic validation of incoming messages.
// This is expected to be used only after DisableValidation() is used, as the validation
// is enabled by default.
func (c *Conn) EnableValidation() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.validationEnabled = true
}

// DisableValidation turns off automatic validation of incoming messages.
// It is not recommended to use this except the node is in debugging mode.
func (c *Conn) DisableValidation() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.validationEnabled = false
}

func (c *Conn) validate(senderAddr net.Addr, msg messages.Message) error {
	// check GTP version
	if msg.Version() != 2 {
		if err := c.VersionNotSupportedIndication(senderAddr, msg); err != nil {
			return err
		}
	}

	// check if TEID is known or not
	if teid := msg.TEID(); teid != 0 {
		if _, err := c.GetSessionByTEID(teid); err != nil {
			return &ErrInvalidTEID{TEID: teid}
		}
	}
	return nil
}

// SendMessageTo sends a message to specified addr.
// By using this function instead of WriteTo, package sets the Sequence Number
// properly and returns the one used to send the message.
func (c *Conn) SendMessageTo(msg messages.Message, addr net.Addr) (uint32, error) {
	seq := c.IncSequence()
	msg.SetSequenceNumber(seq)

	payload, err := messages.Serialize(msg)
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

// EchoResponse sends a EchoResponse to the EchoRequest.
func (c *Conn) EchoResponse(raddr net.Addr, req messages.Message) error {
	res := messages.NewEchoResponse(0, ies.NewRecovery(c.RestartCounter))

	if err := c.RespondTo(raddr, req, res); err != nil {
		return err
	}
	return nil
}

// VersionNotSupportedIndication just sends VersionNotSupportedIndication message.
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
// By creating a Session with this method, a Bearer named "default" is also created
// to be used as default bearer. The default bearer can be retrieved by using
// (*Session) GetDefaultBearer() or (*Session) LookupBearerByName("default").
//
// Note that this method doesn't care IEs given are sufficient or not, as the required IE
// varies much depending on the context Create Session Request is used.
func (c *Conn) CreateSession(raddr net.Addr, ie ...*ies.IE) (*Session, uint32, error) {
	// retrieve values from IEs given.
	sess := NewSession(raddr, &Subscriber{Location: &Location{}})
	br := sess.GetDefaultBearer()
	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			sess.IMSI = i.IMSI()
		case ies.MSISDN:
			sess.MSISDN = i.MSISDN()
		case ies.MobileEquipmentIdentity:
			sess.IMEI = i.MobileEquipmentIdentity()
		case ies.ServingNetwork:
			sess.MCC = i.MCC()
			sess.MNC = i.MNC()
		case ies.AccessPointName:
			br.APN = i.AccessPointName()
		case ies.RATType:
			sess.RATType = i.RATType()
		case ies.FullyQualifiedTEID:
			sess.AddTEID(i.InterfaceType(), i.TEID())
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				for _, child := range i.ChildIEs {
					switch child.Type {
					case ies.EPSBearerID:
						br.EBI = child.EPSBearerID()
					case ies.BearerQoS:
						br.PL = child.PriorityLevel()
						br.QCI = child.QCILabel()
						br.PCI = child.PreemptionCapability()
						br.PVI = child.PreemptionVulnerability()
						br.MBRUL = child.MBRForUplink()
						br.MBRDL = child.MBRForDownlink()
						br.GBRUL = child.GBRForUplink()
						br.GBRDL = child.GBRForUplink()
					case ies.FullyQualifiedTEID:
						sess.AddTEID(i.InterfaceType(), i.TEID())
					case ies.BearerTFT:
						// XXX - do nothing for BearerTFT?
					}
				}
			case 1:
				// XXX - do nothing for BearerContextsToBeRemoved?
			}
		}
	}

	// set IEs into CreateSessionRequest.
	msg := messages.NewCreateSessionRequest(0, 0, ie...)

	seq, err := c.SendMessageTo(msg, raddr)
	if err != nil {
		return nil, 0, err
	}
	return sess, seq, nil
}

// DeleteSession sends a DeleteSessionRequest with TEID and IEs given..
func (c *Conn) DeleteSession(teid uint32, ie ...*ies.IE) (uint32, error) {
	sess, err := c.GetSessionByTEID(teid)
	if err != nil {
		return 0, err
	}

	msg := messages.NewDeleteSessionRequest(teid, 0, ie...)

	seq, err := c.SendMessageTo(msg, sess.PeerAddr)
	if err != nil {
		return 0, err
	}
	return seq, nil
}

// ModifyBearer sends a ModifyBearerRequest with TEID and IEs given..
func (c *Conn) ModifyBearer(teid uint32, ie ...*ies.IE) (uint32, error) {
	sess, err := c.GetSessionByTEID(teid)
	if err != nil {
		return 0, err
	}

	msg := messages.NewModifyBearerRequest(teid, 0, ie...)

	seq, err := c.SendMessageTo(msg, sess.PeerAddr)
	if err != nil {
		return 0, err
	}
	return seq, nil
}

// DeleteBearer sends a DeleteBearerRequest TEID and with IEs given.
func (c *Conn) DeleteBearer(teid uint32, ie ...*ies.IE) (uint32, error) {
	sess, err := c.GetSessionByTEID(teid)
	if err != nil {
		return 0, err
	}

	msg := messages.NewDeleteBearerRequest(teid, 0, ie...)

	seq, err := c.SendMessageTo(msg, sess.PeerAddr)
	if err != nil {
		return 0, err
	}
	return seq, nil
}

// RespondTo sends a message(specified with "toBeSent" param) in response to
// a message(specified with "received" param).
//
// This is to make it easier to handle SequenceNumber.
func (c *Conn) RespondTo(raddr net.Addr, received, toBeSent messages.Message) error {
	toBeSent.SetSequenceNumber(received.Sequence())
	b := make([]byte, toBeSent.Len())

	if err := toBeSent.SerializeTo(b); err != nil {
		return err
	}

	if _, err := c.WriteTo(b, raddr); err != nil {
		return err
	}
	return nil
}

// GetSessionByTEID returns the current session looked up by InterfaceType and TEID of the message.
func (c *Conn) GetSessionByTEID(teid uint32) (*Session, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var session *Session
	for _, sess := range c.Sessions {
		sess.teidMap.rangeWithFunc(func(i, t interface{}) bool {
			if teid == t {
				session = sess
				return false
			}
			return true
		})
		if session != nil {
			return session, nil
		}
	}

	return nil, &ErrInvalidTEID{TEID: teid}
}

// GetSessionByIMSI returns the current session looked up by IMSI.
func (c *Conn) GetSessionByIMSI(imsi string) (*Session, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, sess := range c.Sessions {
		if imsi == sess.IMSI {
			return sess, nil
		}
	}

	return nil, &ErrUnknownIMSI{IMSI: imsi}
}

// GetIMSIByTEID returns IMSI associated with TEID.
func (c *Conn) GetIMSIByTEID(teid uint32) (string, error) {
	sess, err := c.GetSessionByTEID(teid)
	if err != nil {
		return "", err
	}

	return sess.IMSI, nil
}

// AddSession adds a session to c.Sessions.
// If the session given already exists, this removes the old one.
func (c *Conn) AddSession(session *Session) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// TODO: any smarter way?
	if len(c.Sessions) == 0 {
		c.Sessions = []*Session{session}
		return
	}

	var (
		newSessions []*Session
		exists      bool
	)
	for _, oldSession := range c.Sessions {
		if session.IMSI == oldSession.IMSI {
			exists = true
			newSessions = append(newSessions, session)
			continue
		}
		newSessions = append(newSessions, oldSession)
	}
	if !exists {
		newSessions = append(newSessions, session)
	}

	c.Sessions = newSessions
}

// RemoveSession removes a session from c.Session.
func (c *Conn) RemoveSession(session *Session) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var newSessions []*Session
	for _, sess := range c.Sessions {
		if session.IMSI == sess.IMSI {
			continue
		}
		newSessions = append(newSessions, sess)
	}

	c.Sessions = newSessions
}

// NewFTEID creates a new F-TEID with random TEID value that is different from existing one.
// If there's a lot of Session on the Conn, it may take a long time to find unique one.
func (c *Conn) NewFTEID(ifType uint8, v4, v6 string) (fteidIE *ies.IE) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var teids []uint32
	for _, sess := range c.Sessions {
		if teid, ok := sess.teidMap.load(ifType); ok {
			teids = append(teids, teid)
		}
	}

	return ies.NewFullyQualifiedTEID(ifType, generateUniqueUint32(teids), v4, v6)
}

func generateUniqueUint32(vals []uint32) uint32 {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return 0
	}

	generated := binary.BigEndian.Uint32(b)
	for _, existing := range vals {
		if generated == existing {
			return generateUniqueUint32(vals)
		}
	}

	return generated
}

// SessionCount returns the number of sessions registered in Conn.
//
// This may have impact on performance in case of large number of Session exists.
func (c *Conn) SessionCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	var count int
	for _, sess := range c.Sessions {
		if sess.IsActive() {
			count++
		}
	}
	return count
}

// BearerCount returns the number of bearers registered in Conn.
//
// This may have impact on performance in case of large number of Session and
// Bearer exist.
func (c *Conn) BearerCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	var count int
	for _, sess := range c.Sessions {
		if sess.IsActive() {
			count += sess.BearerCount()
		}
	}
	return count
}
