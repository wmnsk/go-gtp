// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import (
	"net"
	"sync"
	"time"

	"github.com/wmnsk/go-gtp/v2/messages"
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

	seqChMap map[uint32]chan messages.Message

	// PeerAddr is a net.Addr of the peer of the Session.
	PeerAddr net.Addr

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
		teidMap:    newTeidMap(),
		bearerMap:  newBearerMap("default", &Bearer{QoSProfile: &QoSProfile{}}),
		Subscriber: sub,
		seqChMap:   map[uint32]chan messages.Message{},
	}

	return s
}

// Activate marks a Session active.
func (s *Session) Activate() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.IMSI == "" {
		return &ErrRequiredParameterMissing{"IMSI", "Session must have IMSI set"}
	}

	s.isActive = true
	return nil
}

// Deactivate marks a Session inactive.
func (s *Session) Deactivate() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.isActive = false
	return nil
}

// IsActive reports whether a Session is active or not.
func (s *Session) IsActive() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

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

// PassMessageTo passes the message (typically "triggerred message") to the session
// expecting to receive it.
//
// Note that it should be ensured that WaitMessage is called before calling this.
// Otherwise it may fail with ErrInvalidSequence.
func PassMessageTo(s *Session, msg messages.Message, timeout time.Duration) error {
	s.mu.Lock()
	msgCh, ok := s.seqChMap[msg.Sequence()]
	s.mu.Unlock()
	if !ok {
		return &ErrInvalidSequence{Seq: msg.Sequence()}
	}

	select {
	case msgCh <- msg:
		return nil
	case <-time.After(timeout):
		return ErrTimeout
	}
}

// WaitMessage waits for a message to come.
// Unless the user does not use PassMessage() func, this always fails with timeout.
func (s *Session) WaitMessage(seq uint32, timeout time.Duration) (messages.Message, error) {
	msgCh := make(chan messages.Message)
	s.mu.Lock()
	s.seqChMap[seq] = msgCh
	s.mu.Unlock()

	select {
	case msg := <-msgCh:
		return msg, nil
	case <-time.After(timeout):
		return nil, ErrTimeout
	}
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

// RemoveBearerByEBI removes a Bearer looked up by name.
func (s *Session) RemoveBearerByEBI(ebi uint8) {
	name, err := s.LookupBearerNameByEBI(ebi)
	if err != nil {
		return
	}
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

	return nil, &ErrNoBearerFound{IMSI: s.IMSI}
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
		return nil, &ErrNoBearerFound{IMSI: s.IMSI}

	}
	return bearer, nil
}

// LookupBearerNameByEBI looks up name of Bearer by EBI.
func (s *Session) LookupBearerNameByEBI(ebi uint8) (string, error) {
	var name string
	s.bearerMap.rangeWithFunc(func(n, br interface{}) bool {
		bearer := br.(*Bearer)
		if ebi == bearer.EBI {
			name = n.(string)
			return false
		}
		return true
	})

	if name == "" {
		return "", &ErrNoBearerFound{IMSI: s.IMSI}

	}
	return name, nil
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

// Bearers returns all the bearers registered in Session.
func (s *Session) Bearers() []*Bearer {
	s.mu.Lock()
	defer s.mu.Unlock()

	var bs []*Bearer
	s.bearerMap.rangeWithFunc(func(k, v interface{}) bool {
		bs = append(bs, v.(*Bearer))
		return true
	})

	return bs
}

// BearerCount returns the number of bearers registered in Session.
func (s *Session) BearerCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	var count int
	s.bearerMap.rangeWithFunc(func(k, v interface{}) bool {
		count++
		return true
	})

	return count
}
