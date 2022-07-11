package message_test

import (
	"net"
)

var (
	mac1, _ = net.ParseMAC("12:34:56:78:90:01")
	mac2, _ = net.ParseMAC("12:34:56:78:90:02")
)
