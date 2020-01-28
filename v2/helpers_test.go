package v2_test

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"testing"

	v2 "github.com/wmnsk/go-gtp/v2"
)

var testConn *v2.Conn
var sessions []*v2.Session
var dummyAddr net.Addr = &net.UDPAddr{IP: net.IP{0x00, 0x00, 0x00, 0x00}, Port: 2123}

func init() {
	testConn = v2.NewConn(dummyAddr, 0, v2.IFTypeS11MMEGTPC)
	sessions = []*v2.Session{
		v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: "001011234567891"}),
		v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: "001011234567892"}),
		v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: "001011234567893"}),
		v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: "001011234567894"}),
	}

	for i, sess := range sessions {
		_ = sess.Activate()
		sess.AddTEID(v2.IFTypeS11MMEGTPC, uint32(i+1))
		testConn.AddSession(sess)
	}
}

func TestSessionCount(t *testing.T) {
	if want, got := testConn.SessionCount(), len(sessions); want != got {
		t.Errorf("SessionCount is invalid. want: %d, got: %d", want, got)
	}
}

func TestGetSessionByIMSI_GetTEID(t *testing.T) {
	for i := 1; i <= testConn.SessionCount(); i++ {
		lastDigit := strconv.Itoa(i)
		sess, err := testConn.GetSessionByIMSI("00101123456789" + lastDigit)
		if err != nil {
			t.Fatal(err)
		}

		teid, err := sess.GetTEID(v2.IFTypeS11MMEGTPC)
		if err != nil {
			t.Fatal(err)
		}

		if teid != uint32(i) {
			t.Errorf("Got wrong TEID at %d, %d, %s", i, teid, sess.IMSI)
		}
	}
}

func BenchmarkAddSession(b *testing.B) {
	for k := 0.; k < 6; k++ {
		existingSessions := int(math.Pow(10, k))
		benchConn := v2.NewConn(dummyAddr, 0, v2.IFTypeS11MMEGTPC)
		for i := 0; i < existingSessions; i++ {
			imsi := fmt.Sprintf("%015d", i)
			sess := v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: imsi})
			sess.AddTEID(v2.IFTypeS11MMEGTPC, uint32(i+1))
			benchConn.AddSession(sess)
		}

		session := v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: "001011234567891"})
		session.AddTEID(v2.IFTypeS11MMEGTPC, 0xdead)
		b.Run(fmt.Sprintf("%d", existingSessions), func(b *testing.B) {
			for i := 1; i <= b.N; i++ {
				benchConn.AddSession(session)
			}
		})
	}
}

func TestGetSessionByTEID(t *testing.T) {
	for i := 1; i <= testConn.SessionCount(); i++ {
		sess, err := testConn.GetSessionByTEID(uint32(i), dummyAddr)
		if err != nil {
			t.Fatal(err)
		}

		lastDigit := strconv.Itoa(i)
		if string(sess.IMSI[14]) != lastDigit {
			t.Errorf("Got wrong session at %d, %s", i, sess.IMSI)
		}
	}
}

func TestGetIMSIByTEID(t *testing.T) {
	for i := 1; i <= testConn.SessionCount(); i++ {
		imsi, err := testConn.GetIMSIByTEID(uint32(i), dummyAddr)
		if err != nil {
			t.Fatal(err)
		}

		lastDigit := strconv.Itoa(i)
		if string(imsi[14]) != lastDigit {
			t.Errorf("Got wrong IMSI at %d, %s", i, imsi)
		}
	}
}

func TestRemoveSession(t *testing.T) {
	conn := *testConn // copy testConn
	conn.RemoveSession(sessions[0])

	if conn.SessionCount() != len(sessions)-1 {
		t.Errorf("Session not removed expectedly: %d, %v", conn.SessionCount(), conn.Sessions())
	}

	for i := 2; i <= testConn.SessionCount(); i++ {
		sess, err := testConn.GetSessionByTEID(uint32(i), dummyAddr)
		if err != nil {
			t.Fatal(err)
		}

		lastDigit := strconv.Itoa(i)
		if string(sess.IMSI[14]) != lastDigit {
			t.Errorf("Got wrong session at %d, %s", i, sess.IMSI)
		}
	}
}

func TestRemoveSessionByIMSI(t *testing.T) {
	conn := *testConn // copy testConn
	conn.RemoveSessionByIMSI("001011234567891")

	if conn.SessionCount() != len(sessions)-1 {
		t.Errorf("Session not removed expectedly: %d, %v", conn.SessionCount(), conn.Sessions())
	}

	for i := 2; i <= testConn.SessionCount(); i++ {
		sess, err := testConn.GetSessionByTEID(uint32(i), dummyAddr)
		if err != nil {
			t.Fatal(err)
		}

		lastDigit := strconv.Itoa(i)
		if string(sess.IMSI[14]) != lastDigit {
			t.Errorf("Got wrong session at %d, %s", i, sess.IMSI)
		}
	}
}
