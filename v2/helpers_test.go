package v2_test

import (
	"fmt"
	"net"
	"strconv"
	"testing"

	v2 "github.com/wmnsk/go-gtp/v2"
)

var testConn *v2.Conn
var dummyAddr net.Addr = &net.UDPAddr{IP: net.IP{0x00, 0x00, 0x00, 0x00}, Port: 2123}

func init() {
	testConn = &v2.Conn{
		Sessions: []*v2.Session{
			v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: "001011234567891"}),
			v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: "001011234567892"}),
			v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: "001011234567893"}),
			v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: "001011234567894"}),
		},
	}

	for i, sess := range testConn.Sessions {
		_ = sess.Activate()
		sess.AddTEID(v2.IFTypeS11MMEGTPC, uint32(i+1))
		testConn.AddSession(sess)
	}
}

func TestGetSessionByIMSI_GetTEID(t *testing.T) {
	for i := 1; i <= len(testConn.Sessions); i++ {
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

func benchmarkAddSession(numExisitingSessions int, b *testing.B) {
	benchConn := &v2.Conn{Sessions: []*v2.Session{}}
	for i := 0; i < numExisitingSessions; i++ {
		imsi := fmt.Sprintf("%015d", i)
		benchConn.AddSession(v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: imsi}))
	}
	b.ResetTimer()
	for i := 1; i <= b.N; i++ {
		benchConn.AddSession(v2.NewSession(dummyAddr, &v2.Subscriber{IMSI: "001011234567891"}))
	}
}

func BenchmarkAddSessionExist0(b *testing.B)    { benchmarkAddSession(0, b) }
func BenchmarkAddSessionExist100(b *testing.B)  { benchmarkAddSession(1e2, b) }
func BenchmarkAddSessionExist1K(b *testing.B)   { benchmarkAddSession(1e3, b) }
func BenchmarkAddSessionExist10K(b *testing.B)  { benchmarkAddSession(1e4, b) }
func BenchmarkAddSessionExist100K(b *testing.B) { benchmarkAddSession(1e5, b) }
func BenchmarkAddSessionExist1M(b *testing.B)   { benchmarkAddSession(1e6, b) }

func TestGetSessionByTEID(t *testing.T) {
	for i := 1; i <= len(testConn.Sessions); i++ {
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
	for i := 1; i <= len(testConn.Sessions); i++ {
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
	conn.RemoveSession(testConn.Sessions[0])

	if conn.SessionCount() != len(testConn.Sessions)-1 {
		t.Errorf("Session not removed expectedly: %d, %v", conn.SessionCount(), conn.Sessions)
	}

	for i := 1; i <= len(testConn.Sessions); i++ {
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

	if conn.SessionCount() != len(testConn.Sessions)-1 {
		t.Errorf("Session not removed expectedly: %d, %v", conn.SessionCount(), conn.Sessions)
	}

	for i := 1; i <= len(testConn.Sessions); i++ {
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

func TestSessionCount(t *testing.T) {
	if want, got := testConn.SessionCount(), len(testConn.Sessions); want != got {
		t.Errorf("SessionCount is invalid. want: %d, got: %d", want, got)
	}
}
