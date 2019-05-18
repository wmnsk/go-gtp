// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import (
	"crypto/rand"
	"encoding/binary"
	"net"
	"sync"

	"github.com/wmnsk/go-gtp/v2/ies"
)

// Location is a subscriber's location.
type Location struct {
	MCC, MNC               string
	RATType                uint8
	LAC, CI, SAI, RAI, TAI uint16
	ECI, MeNBI, EMeNBI     uint32
}

// Subscriber is a subscriber that belongs to a GTPv2 session.
type Subscriber struct {
	IMSI, MSISDN, IMEI string
	*Location
}

// Session is a GTPv2 Session.
type Session struct {
	mu       sync.Mutex
	isActive bool
	*teidMap
	*bearerMap

	// PeerAddr is a net.Addr of the peer of the Session.
	PeerAddr net.Addr

	// Sequence is the last SequenceNumber used in the request.
	// This should be incremented when used manually by users.
	Sequence uint32

	// Subscriber is a Subscriber associated with the Session.
	*Subscriber
}

// NewSession creates a new Session with subscriber information.
//
// This is expected to be used by server-like nodes. Otherwise, use CreateSession(),
// which sends Create Session Request and returns a new Session.
func NewSession(peerAddr net.Addr, sub *Subscriber) *Session {
	s := &Session{
		mu:         sync.Mutex{},
		PeerAddr:   peerAddr,
		Subscriber: sub,
		teidMap:    newTeidMap(),
		bearerMap:  newBearerMap("default", &Bearer{QoSProfile: &QoSProfile{}}),
	}

	u32buf := make([]byte, 4)
	if _, err := rand.Read(u32buf); err != nil {
		u32buf = []byte{0x00, 0x00, 0x00, 0x00}
	}
	s.Sequence = binary.BigEndian.Uint32(u32buf)

	return s
}

// NewSessionWithNetlink creates a new Session with subscriber information, using
// Netlink-based GTP-U bearer with the version given.
//
// This is expected to be used by server-like nodes. Otherwise, use CreateSession(),
// which sends Create Session Request and returns a new Session.
func NewSessionWithNetlink(peerAddr net.Addr, sub *Subscriber, uVer uint32) *Session {
	s := &Session{
		mu:         sync.Mutex{},
		PeerAddr:   peerAddr,
		Subscriber: sub,
		teidMap:    newTeidMap(),
		bearerMap: newBearerMap(
			"default",
			NewNetlinkBearer(uVer, 0, "", &QoSProfile{}),
		),
	}

	u32buf := make([]byte, 4)
	if _, err := rand.Read(u32buf); err != nil {
		u32buf = []byte{0x00, 0x00, 0x00, 0x00}
	}
	s.Sequence = binary.BigEndian.Uint32(u32buf)

	return s
}

// CreateSession is an alias for (*Conn).CreateSessionRequest.
// See (*Conn).CreateSessionRequest for details.
func CreateSession(raddr net.Addr, c *Conn, ie ...*ies.IE) (*Session, error) {
	return c.CreateSession(raddr, ie...)
}

// DeleteSession is an alias for (*Conn).DeleteSessionRequest.
// See (*Conn).DeleteSessionRequest for details.
func DeleteSession(c *Conn, teid uint32, ie ...*ies.IE) error {
	return c.DeleteSession(teid, ie...)
}

// Delete sends a Delete Session Request toward the interface which
// is specified with c and ifType.
//
// By default, IEs on the Delete Session Request is only EBI of default
// bearer, but it can be overridden by giving EBI IE.
// Also, other IEs can be added by giving them as ie.
func (s *Session) Delete(c *Conn, ifType uint8, ie ...*ies.IE) error {
	// do nothing for non-active Session
	if !s.IsActive() {
		return nil
	}

	teid, err := s.GetTEID(ifType)
	if err != nil {
		return err
	}

	// send EBI of default bearer by default, but if the same type of
	// IE is given, the default one is replaced.
	ieToSend := []*ies.IE{ies.NewEPSBearerID(s.GetDefaultBearer().EBI)}
	for _, i := range ie {
		if i.Type == ies.EPSBearerID {
			ieToSend[0] = i
			continue
		}
		// other IEs given are just put regardless of their type.
		ieToSend = append(ieToSend, i)
	}

	return c.DeleteSession(teid, ieToSend...)
}

// ModifyBearer sends a Modify Bearer Request toward the interface which
// is specified with c and ifType.
//
// By default, IEs on the Modify Bearer Request is only EBI of default
// bearer, but it can be overridden by giving EBI IE.
// Also, other IEs can be added by giving them as ie.
func (s *Session) ModifyBearer(c *Conn, ifType uint8, ie ...*ies.IE) error {
	// do nothing for non-active Session
	if !s.IsActive() {
		return nil
	}

	teid, err := s.GetTEID(ifType)
	if err != nil {
		return err
	}

	return c.ModifyBearer(teid, ie...)
}

// Activate marks a Session active.
func (s *Session) Activate() error {
	if s.IMSI == "" {
		return &ErrRequiredParameterMissing{"IMSI", "Session must have IMSI set"}
	}

	s.mu.Lock()
	s.isActive = true
	s.mu.Unlock()
	return nil
}

// Deactivate marks a Session inactive.
func (s *Session) Deactivate() error {
	s.mu.Lock()
	s.isActive = false
	s.mu.Unlock()
	return nil
}

// IsActive reports whether a Session is active or not.
func (s *Session) IsActive() bool {
	return s.isActive
}

// AddTEID adds TEID to session with InterfaceType.
func (s *Session) AddTEID(ifType uint8, teid uint32) {
	s.teidMap.store(ifType, teid)
}

// GetTEID returns TEID associated with InterfaceType given.
func (s *Session) GetTEID(ifType uint8) (uint32, error) {
	if teid, ok := s.teidMap.load(ifType); ok {
		return teid, nil
	}
	return 0, ErrTEIDNotFound
}

// AddBearer adds a Bearer to Session with arbitrary name given.
//
// In the single-bearer environment it is not used, as a bearer named "default" is
// always available after created a Session.
func (s *Session) AddBearer(name string, br *Bearer) {
	s.bearerMap.store(name, br)
}

// RemoveBearer removes a Bearer looked up by name.
func (s *Session) RemoveBearer(name string) {
	s.bearerMap.delete(name)
}

// GetDefaultBearer returns the pointer to default bearer.
func (s *Session) GetDefaultBearer() *Bearer {
	// it is not expected that the default bearer cannot be found.
	bearer, ok := s.bearerMap.load("default")
	if !ok {
		return nil
	}

	return bearer
}

// SetDefaultBearer sets given bearer as the default bearer.
func (s *Session) SetDefaultBearer(bearer *Bearer) {
	// it is not expected that the default bearer cannot be found.
	s.bearerMap.store("default", bearer)
}

// LookupBearerByName looks up Bearer registered in Session by name.
func (s *Session) LookupBearerByName(name string) (*Bearer, error) {
	if br, ok := s.bearerMap.load(name); ok {
		return br, nil
	}

	return nil, ErrNoBearerFound
}

// LookupBearerByEBI looks up Bearer registered in Session by EBI.
func (s *Session) LookupBearerByEBI(ebi uint8) (*Bearer, error) {
	var bearer *Bearer
	s.bearerMap.rangeWithFunc(func(name, br interface{}) bool {
		b := br.(*Bearer)
		if ebi == b.EBI {
			bearer = b
			return false
		}
		return true
	})

	if bearer == nil {
		return nil, ErrNoBearerFound

	}
	return bearer, nil
}

// LookupEBIByName returns EBI associated with Name given.
//
// If no EBI found, it returns 0(invalid value for EBI).
func (s *Session) LookupEBIByName(name string) uint8 {
	if br, ok := s.bearerMap.load(name); ok {
		return br.EBI
	}

	return 0
}

// LookupEBIByTEID returns EBI associated with TEID given.
//
// If no EBI found, it returns 0(=invalid value for EBI).
func (s *Session) LookupEBIByTEID(teid uint32) uint8 {
	var ebi uint8
	s.bearerMap.rangeWithFunc(func(name, bearer interface{}) bool {
		br := bearer.(*Bearer)
		if br.IsNetlink() {
			if teid == br.PDP.ITEI || teid == br.PDP.OTEI {
				ebi = br.EBI
				return false
			}
		}
		if teid == br.teidIn || teid == br.teidOut {
			ebi = br.EBI
			return false
		}
		return true
	})

	return ebi
}

type teidMap struct {
	syncMap sync.Map
}

func newTeidMap() *teidMap {
	return &teidMap{}
}

func (t *teidMap) store(ifType uint8, teid uint32) {
	t.syncMap.Store(ifType, teid)
}

func (t *teidMap) load(ifType uint8) (uint32, bool) {
	teid, ok := t.syncMap.Load(ifType)
	if !ok {
		return 0, false
	}

	return teid.(uint32), true
}

func (t *teidMap) rangeWithFunc(fn func(ifType, teid interface{}) bool) {
	t.syncMap.Range(fn)
}

type bearerMap struct {
	syncMap sync.Map
}

func newBearerMap(name string, bearer *Bearer) *bearerMap {
	b := &bearerMap{}
	b.store(name, bearer)

	return b
}
func (b *bearerMap) store(name string, bearer *Bearer) {
	b.syncMap.Store(name, bearer)
}

func (b *bearerMap) load(name string) (*Bearer, bool) {
	bearer, ok := b.syncMap.Load(name)
	if !ok {
		return nil, false
	}

	return bearer.(*Bearer), true
}

func (b *bearerMap) delete(name string) {
	b.syncMap.Delete(name)
}

func (b *bearerMap) rangeWithFunc(fn func(name, bearer interface{}) bool) {
	b.syncMap.Range(fn)
}
