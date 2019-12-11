package v2_test

import (
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
