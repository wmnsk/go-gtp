// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package gtpv1

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/vishvananda/netlink"
	"github.com/wmnsk/go-gtp/gtpv1/ie"
	"github.com/wmnsk/go-gtp/gtpv1/message"
	v2ie "github.com/wmnsk/go-gtp/gtpv2/ie"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

type tpduSet struct {
	raddr   net.Addr
	teid    uint32
	seq     uint16
	payload []byte
}

type pktConn interface {
	// WriteToWithDSCPECN writes a packet with payload p to addr using the given DSCP/ECN value.
	// WriteToWithDSCPECN can be made to time out and return
	// an Error with Timeout() == true after a fixed time limit;
	// see SetDeadline and SetWriteDeadline.
	// On packet-oriented connections, write timeouts are rare.
	WriteToWithDSCPECN(p []byte, addr net.Addr, dscpecn int) (n int, err error)

	// File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished.
	// Closing c does not affect f, and closing f does not affect c.
	// The returned os.File's file descriptor is different from the connection's.
	// Attempting to change properties of the original using this duplicate may or may not have the desired effect.
	File() (f *os.File, err error)

	net.PacketConn
}

type pktConn4 struct {
	// mu is the mutex used before Writing to the PacketConn,
	// to be sure the right DSCP/ECN value
	// is applied before performing the Write.
	mu *sync.Mutex

	// udpConn is the UDPConn used as underlying transport
	udpConn *net.UDPConn

	*ipv4.PacketConn
}

// ReadFrom implements the io.ReaderFrom ReadFrom method.
func (pkt pktConn4) ReadFrom(b []byte) (n int, src net.Addr, err error) {
	n, _, src, err = pkt.PacketConn.ReadFrom(b)
	return n, src, err
}

// WriteTo implements the PacketConn WriteTo method.
func (pkt pktConn4) WriteTo(b []byte, dst net.Addr) (n int, err error) {
	return pkt.PacketConn.WriteTo(b, nil, dst)
}

// setDSCPECN sets the DSCP/ECN value used for next writes.
func (pkt pktConn4) setDSCPECN(dscpecs int) error {
	// With IPv4, DSCP/ECN is using the TOS field
	return pkt.SetTOS(dscpecs)
}

// DSCPECN returns the DSCP/ECN value.
func (pkt pktConn4) DSCPECN() (int, error) {
	// With IPv4, DSCP/ECN is using the TOS field
	return pkt.TOS()
}

// WriteToWithDSCPECN implements the pktConn WriteToWithDSCPECN method.
func (pkt pktConn4) WriteToWithDSCPECN(p []byte, addr net.Addr, dscpecn int) (n int, err error) {
	pkt.mu.Lock()
	defer pkt.mu.Unlock()
	oldDSCPECN, err := pkt.DSCPECN()
	if err != nil {
		return 0, err
	}
	err = pkt.setDSCPECN(dscpecn)
	if err != nil {
		return 0, err
	}
	defer func() {
		// set back DSCP/ECN for next write calls
		_ = pkt.setDSCPECN(oldDSCPECN)
	}()
	return pkt.WriteTo(p, addr)
}

// File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
// The returned os.File's file descriptor is different from the connection's.
// Attempting to change properties of the original using this duplicate may or may not have the desired effect.
func (pkt pktConn4) File() (f *os.File, err error) {
	return pkt.udpConn.File()
}

type pktConn6 struct {
	// mu is the mutex used before Writing to the PacketConn,
	// to be sure the right DSCP/ECN value
	// is applied before performing the Write.
	mu *sync.Mutex

	// udpConn is the UDPConn used as underlying transport.
	udpConn *net.UDPConn

	*ipv6.PacketConn
}

// ReadFrom implements the io.ReaderFrom ReadFrom method.
func (pkt pktConn6) ReadFrom(b []byte) (n int, src net.Addr, err error) {
	n, _, src, err = pkt.PacketConn.ReadFrom(b)
	return n, src, err
}

// WriteTo implements the PacketConn WriteTo method.
func (pkt pktConn6) WriteTo(b []byte, dst net.Addr) (n int, err error) {
	return pkt.PacketConn.WriteTo(b, nil, dst)
}

// setDSCPECN sets the DSCP/ECN value used for next writes.
func (pkt pktConn6) setDSCPECN(dscpecs int) error {
	// With IPv6, DSCP/ECN is using the Traffic Class field
	return pkt.SetTrafficClass(dscpecs)
}

// DSCPECN returns the DSCP/ECN value.
func (pkt pktConn6) DSCPECN() (int, error) {
	// With IPv6, DSCP/ECN is using the Traffic Class field
	return pkt.TrafficClass()
}

// WriteToWithDSCPECN implements the pktConn WriteToWithDSCPECN method.
func (pkt pktConn6) WriteToWithDSCPECN(p []byte, addr net.Addr, dscpecn int) (n int, err error) {
	pkt.mu.Lock()
	defer pkt.mu.Unlock()
	oldDSCPECN, err := pkt.DSCPECN()
	if err != nil {
		return 0, err
	}
	err = pkt.setDSCPECN(dscpecn)
	if err != nil {
		return 0, err
	}
	defer func() {
		// set back DSCP/ECN for next write calls
		_ = pkt.setDSCPECN(oldDSCPECN)
	}()
	return pkt.WriteTo(p, addr)
}

// File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
// The returned os.File's file descriptor is different from the connection's.
// Attempting to change properties of the original using this duplicate may or may not have the desired effect.
func (pkt pktConn6) File() (f *os.File, err error) {
	return pkt.udpConn.File()
}

// newPktConn creates a new pktConn initialized with a given local UDP address
func newPktConn(laddr net.Addr) (pktConn, error) {
	var err error
	pktC, err := net.ListenPacket(laddr.Network(), laddr.String())
	if err != nil {
		return nil, err
	}
	// check if UDPConn is over IPv4 or IPv6
	addr, err := net.ResolveUDPAddr(laddr.Network(), laddr.String())
	if err != nil {
		return nil, err
	}
	if addr.IP.To4() != nil {
		return pktConn4{
			mu:         &sync.Mutex{},
			udpConn:    pktC.(*net.UDPConn),
			PacketConn: ipv4.NewPacketConn(pktC),
		}, nil
	} else if addr.IP.To16() != nil {
		return pktConn6{
			mu:         &sync.Mutex{},
			udpConn:    pktC.(*net.UDPConn),
			PacketConn: ipv6.NewPacketConn(pktC),
		}, nil
	} else {
		return nil, fmt.Errorf("laddr must refer to an IP address")
	}
}

// UPlaneConn represents a U-Plane Connection of GTPv1.
type UPlaneConn struct {
	mu      sync.Mutex
	laddr   net.Addr
	pktConn pktConn
	*msgHandlerMap
	*iteiMap

	tpduCh  chan *tpduSet
	closeCh chan struct{}

	relayMap map[uint32]*peer

	errIndEnabled bool

	// for Linux kernel GTP with netlink
	KernelGTP
}

// KernelGTP consists of the Linux Kernel GTP-U related objects.
type KernelGTP struct {
	enabled  bool
	connFile *os.File
	Link     *netlink.GTP
}

// NewUPlaneConn creates a new UPlaneConn used for server. On client side, use DialUPlane instead.
func NewUPlaneConn(laddr net.Addr) *UPlaneConn {
	return &UPlaneConn{
		mu:            sync.Mutex{},
		msgHandlerMap: newDefaultMsgHandlerMap(),
		iteiMap:       newiteiMap(),
		laddr:         laddr,

		tpduCh:  make(chan *tpduSet),
		closeCh: make(chan struct{}),

		errIndEnabled: true,
	}
}

// DialUPlane sends Echo Request to raddr to check if the endpoint is alive and returns UPlaneConn.
func DialUPlane(ctx context.Context, laddr, raddr net.Addr) (*UPlaneConn, error) {
	u := &UPlaneConn{
		mu:            sync.Mutex{},
		msgHandlerMap: newDefaultMsgHandlerMap(),
		iteiMap:       newiteiMap(),
		laddr:         laddr,

		tpduCh:  make(chan *tpduSet),
		closeCh: make(chan struct{}),

		errIndEnabled: true,
	}

	// setup UDPConn first.
	var err error
	if u.pktConn == nil {
		u.pktConn, err = newPktConn(u.laddr)
		if err != nil {
			return nil, err
		}
	}

	// if no response coming within 5 seconds, returns error.
	if err := u.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		return nil, err
	}

	buf := make([]byte, 1600)
	for {
		select {
		case <-ctx.Done():
			return nil, nil
		default:
			// go forward
		}

		// send EchoRequest to raddr.
		if err := u.EchoRequest(raddr); err != nil {
			return nil, err
		}

		n, _, err := u.ReadFrom(buf)
		if err != nil {
			return nil, err
		}
		if err := u.SetReadDeadline(time.Time{}); err != nil {
			return nil, err
		}

		// decode incoming message and let it be handled by default handler funcs.
		msg, err := message.Parse(buf[:n])
		if err != nil {
			return nil, err
		}
		if _, ok := msg.(*message.EchoResponse); !ok {
			continue
		}

		break
	}

	go func() {
		if err := u.serve(ctx); err != nil {
			logf("fatal error on UPlaneConn %s: %s", u.LocalAddr(), err)
		}
	}()

	return u, nil
}

// ListenAndServe creates a new GTPv2-C *Conn and start serving.
// This blocks, and returns error only if it face the fatal one. Non-fatal errors are logged
// with logger. See SetLogger/EnableLogger/DisableLogger for handling of those logs.
func (u *UPlaneConn) ListenAndServe(ctx context.Context) error {
	if u.pktConn == nil {
		var err error
		u.mu.Lock()
		u.pktConn, err = newPktConn(u.laddr)
		u.mu.Unlock()
		if err != nil {
			return err
		}
	}
	return u.listenAndServe(ctx)
}

func (u *UPlaneConn) listenAndServe(ctx context.Context) error {
	// TODO: this func is left for future enhancement.
	return u.serve(ctx)
}

func (u *UPlaneConn) serve(ctx context.Context) error {
	go func() {
		select { // ctx is canceled or Close() is called
		case <-ctx.Done():
		case <-u.closed():
		}

		if u.KernelGTP.enabled {
			if err := u.KernelGTP.connFile.Close(); err != nil {
				logf("error closing GTPFile: %s", err)
			}
			if err := netlink.LinkDel(u.KernelGTP.Link); err != nil {
				logf("error deleting GTPLink: %s", err)
			}
		}

		// This doesn't finish for some reason when Kernel GTP is enabled.
		if u.pktConn != nil {
			if err := u.pktConn.Close(); err != nil {
				logf("error closing the underlying conn: %s", err)
			}
		}
	}()

	buf := make([]byte, 1500)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-u.closed():
			return nil
		default:
			// do nothing and go forward.
		}

		n, raddr, err := u.ReadFrom(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			// TODO: Use net.ErrClosed instead (available from Go 1.16).
			// https://github.com/golang/go/commit/e9ad52e46dee4b4f9c73ff44f44e1e234815800f
			if strings.Contains(err.Error(), "use of closed network connection") {
				return nil
			}
			return fmt.Errorf("error reading from UPlaneConn %s: %w", u.LocalAddr(), err)
		}

		raw := make([]byte, n)
		copy(raw, buf)
		go func() {
			// just forward T-PDU instead of passing it to reader if relayer is
			// configured and the message type is T-PDU.
			if len(u.relayMap) != 0 && raw[1] == message.MsgTypeTPDU {
				// ignore if the packet size is smaller than minimum header size
				if n < 11 {
					return
				}

				u.mu.Lock()
				peer, ok := u.relayMap[binary.BigEndian.Uint32(raw[4:8])]
				u.mu.Unlock()
				if !ok { // pass message to handler if TEID is unknown
					msg, err := message.Parse(raw[:n])
					if err != nil {
						return
					}

					if err := u.handleMessage(raddr, msg); err != nil {
						// should not stop serving with this error
						logf("error handling message on UPlaneConn %s: %v", u.LocalAddr(), err)
					}
					return
				}

				// just use original packet not to get it slow.
				binary.BigEndian.PutUint32(raw[4:8], peer.teid)
				if _, err := peer.srcConn.WriteToWithDSCPECN(raw[:n], peer.addr, 0); err != nil {
					// should not stop serving with this error
					logf("error sending on UPlaneConn %s: %v", u.LocalAddr(), err)
				}
				return
			}

			msg, err := message.Parse(raw[:n])
			if err != nil {
				logf("error parsing message on UPlaneConn %s: %v", u.LocalAddr(), err)
				return
			}

			if err := u.handleMessage(raddr, msg); err != nil {
				// should not stop serving with this error
				logf("error handling message on UPlaneConn %s: %v", u.LocalAddr(), err)
				return
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
//
// Note that valid GTP-U packets handled by Kernel can NOT be retrieved by this.
func (u *UPlaneConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	return u.pktConn.ReadFrom(p)
}

// ReadFromGTP reads a packet from the connection, copying the payload without
// GTP header into p. It returns the number of bytes copied into p, the return
// address that was on the packet, TEID in the GTP header.
//
// Note that valid GTP-U packets handled by Kernel can NOT be retrieved by this.
func (u *UPlaneConn) ReadFromGTP(p []byte) (n int, addr net.Addr, teid uint32, err error) {
	select {
	case <-u.closed():
		return
	case tpdu, ok := <-u.tpduCh:
		if !ok {
			err = ErrConnNotOpened
			return
		}
		n = copy(p, tpdu.payload)
		addr = tpdu.raddr
		teid = tpdu.teid
		return
	}
}

// WriteTo writes a packet with payload p to addr.
// WriteTo can be made to time out and return
// an Error with Timeout() == true after a fixed time limit;
// see SetDeadline and SetWriteDeadline.
// On packet-oriented connections, write timeouts are rare.
func (u *UPlaneConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	return u.pktConn.WriteToWithDSCPECN(p, addr, 0)
}

// WriteToWithDSCPECN writes a packet with payload p to addr using the given DSCP/ECN value.
// WriteToWithDSCPECN can be made to time out and return
// an Error with Timeout() == true after a fixed time limit;
// see SetDeadline and SetWriteDeadline.
// On packet-oriented connections, write timeouts are rare.
func (u *UPlaneConn) WriteToWithDSCPECN(p []byte, addr net.Addr, dscpecn int) (n int, err error) {
	return u.pktConn.WriteToWithDSCPECN(p, addr, dscpecn)
}

// WriteToGTP writes a packet with TEID and payload to addr.
func (u *UPlaneConn) WriteToGTP(teid uint32, p []byte, addr net.Addr) (n int, err error) {
	b, err := Encapsulate(teid, p).Marshal()
	if err != nil {
		return
	}

	if _, err = u.WriteTo(b, addr); err != nil {
		return
	}
	return len(b), nil
}

// closed would be used in multiple goroutines.
// never send struct{}{} to it; instead, use close(u.closeCh).
func (u *UPlaneConn) closed() <-chan struct{} {
	return u.closeCh
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (u *UPlaneConn) Close() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	close(u.closeCh)

	return nil
}

// LocalAddr returns the local network address.
func (u *UPlaneConn) LocalAddr() net.Addr {
	return u.pktConn.LocalAddr()
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
func (u *UPlaneConn) SetDeadline(t time.Time) error {
	return u.pktConn.SetDeadline(t)
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (u *UPlaneConn) SetReadDeadline(t time.Time) error {
	return u.pktConn.SetReadDeadline(t)
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (u *UPlaneConn) SetWriteDeadline(t time.Time) error {
	return u.pktConn.SetWriteDeadline(t)
}

// AddHandler adds a message handler to *UPlaneConn.
//
// By adding HandlerFuncs, *UPlaneConn (and *Session, *Bearer created by the *UPlaneConn) will handle
// the specified type of message with it's paired HandlerFunc when receiving.
// Messages without registered handlers are just ignored and discarded and the user will
// get ErrNoHandlersFound error.
//
// This should be performed just after creating *UPlaneConn, otherwise the user cannot retrieve
// any values, which is in most cases vital to continue working as a node, from the incoming
// message.
//
// HandlerFuncs for EchoResponse and ErrorIndication are registered by default.
// These HandlerFuncs can be overwritten by specifying message.MsgTypeEchoResponse and/or
// message.MsgTypeErrorIndication as msgType parameter.
func (u *UPlaneConn) AddHandler(msgType uint8, fn HandlerFunc) {
	u.msgHandlerMap.store(msgType, fn)
}

// AddHandlers adds multiple handler funcs at a time.
//
// See AddHandler for detailed usage.
func (u *UPlaneConn) AddHandlers(funcs map[uint8]HandlerFunc) {
	for msgType, fn := range funcs {
		u.msgHandlerMap.store(msgType, fn)
	}
}

func (u *UPlaneConn) handleMessage(senderAddr net.Addr, msg message.Message) error {
	handle, ok := u.msgHandlerMap.load(msg.MessageType())
	if !ok {
		return &HandlerNotFoundError{MsgType: msg.MessageTypeName()}
	}

	if err := handle(u, senderAddr, msg); err != nil {
		return fmt.Errorf("failed to handle %s: %w", msg.MessageTypeName(), err)
	}

	return nil
}

// EchoRequest sends a EchoRequest.
func (u *UPlaneConn) EchoRequest(raddr net.Addr) error {
	b, err := message.NewEchoRequest(0, ie.NewRecovery(0)).Marshal()
	if err != nil {
		return err
	}

	if _, err := u.WriteTo(b, raddr); err != nil {
		return err
	}
	return nil
}

// EchoResponse sends a EchoResponse.
func (u *UPlaneConn) EchoResponse(raddr net.Addr) error {
	b, err := message.NewEchoResponse(0, ie.NewRecovery(0)).Marshal()
	if err != nil {
		return err
	}

	if _, err := u.WriteTo(b, raddr); err != nil {
		return err
	}
	return nil
}

// ErrorIndication just sends ErrorIndication message.
func (u *UPlaneConn) ErrorIndication(raddr net.Addr, received message.Message) error {
	ip, _, err := net.SplitHostPort(u.LocalAddr().String())
	if err != nil {
		return err
	}

	errInd, err := message.NewErrorIndication(
		0, received.Sequence(),
		ie.NewTEIDDataI(received.TEID()),
		ie.NewGSNAddress(ip),
	).Marshal()
	if err != nil {
		return err
	}

	if _, err := u.WriteTo(errInd, raddr); err != nil {
		return err
	}
	return nil
}

// RespondTo sends a message(specified with "toBeSent" param) in response to
// a message(specified with "received" param).
//
// This is to make it easier to handle SequenceNumber.
func (u *UPlaneConn) RespondTo(raddr net.Addr, received, toBeSent message.Message) error {
	toBeSent.SetSequenceNumber(received.Sequence())
	b := make([]byte, toBeSent.MarshalLen())
	if err := toBeSent.MarshalTo(b); err != nil {
		return err
	}

	if _, err := u.WriteTo(b, raddr); err != nil {
		return err
	}
	return nil
}

// Restarts returns the number of restarts in uint8.
func (u *UPlaneConn) Restarts() uint8 {
	return 0
}

// NewFTEID creates a new GTPv2 F-TEID with random TEID value that is unique within UPlaneConn.
// To ensure the uniqueness, don't create in the other way if you once use this method.
// This is meant to be used for creating F-TEID IE for non-local interface type, such as
// the ones that are used in U-Plane. For local interface, use (*Conn).NewSenderFTEID instead.
//
// Note that in the case there's a lot of Session on the Conn, it may take a long
// time to find a new unique value.
//
// TODO: optimize performance...
func (u *UPlaneConn) NewFTEID(ifType uint8, v4, v6 string) (fteidIE *v2ie.IE) {
	var teid uint32
	for try := uint32(0); try < 0xffff; try++ {
		const logEvery = 0xff
		if try&logEvery == logEvery {
			logf("Generating NewSenderFTEID crossed tries:%d", try)
		}

		t := generateRandomUint32()
		if t == 0 {
			continue
		}

		// Try to mark TEID as taken. Fails if something exists
		if ok := u.iteiMap.tryStore(t, time.Now()); !ok {
			logf("TEID-U: %#08x has already been taken, trying to generate another one...", t)
			continue
		}

		teid = t
		break
	}

	if teid == 0 {
		return nil
	}
	return v2ie.NewFullyQualifiedTEID(ifType, teid, v4, v6)
}

func generateRandomUint32() uint32 {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return 0
	}

	return binary.BigEndian.Uint32(b)
}

type iteiMap struct {
	syncMap sync.Map
}

func newiteiMap() *iteiMap {
	return &iteiMap{}
}

func (t *iteiMap) tryStore(itei uint32, ts time.Time) bool {
	_, loaded := t.syncMap.LoadOrStore(itei, ts)
	return !loaded
}

func (t *iteiMap) delete(itei uint32) {
	t.syncMap.Delete(itei)
}

// EnableErrorIndication re-enables automatic sending of
// Error Indication to unknown messages, which is enabled by
// default.
//
// See also: DisableErrorIndication.
func (u *UPlaneConn) EnableErrorIndication() {
	u.mu.Lock()
	u.errIndEnabled = true
	u.mu.Unlock()
}

// DisableErrorIndication makes default T-PDU handler stop
// responding with Error Indication in case of receiving T-PDU
// with unknown TEID.
//
// When disabled, it passes the unhandled T-PDU to user who calls
// ReadFromGTP instead.
func (u *UPlaneConn) DisableErrorIndication() {
	u.mu.Lock()
	u.errIndEnabled = false
	u.mu.Unlock()
}
