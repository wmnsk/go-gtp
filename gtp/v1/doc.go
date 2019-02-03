// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Package v1 provides the simple and painless handling of GTPv1-C and GTPv1-U protocol in pure Golang.
//
// This package is still under construction. The networking fearure is only available for GTPv1-U.
// GTPv1-C feature would be available in the future.
// See messages and ies directory for what you can do with the current implementation.
//
// To open a U-Plane connection, use Dial()` or `ListenAndServe()` to retrieve `UPlaneConn`.The difference between the two functions is;
//
// Dial() sends Echo Request and returns UPlaneConn if it succeeds.
//
//   // give local/remote net.Addr, restart counter, channel to let background process pass the errors.
//   uConn, err := v1.Dial(laddr, raddr, 0, errCh)
//   if err != nil {
//   	// ...
//   }
//
// ListenAndServe() just returns UPlaneConn without any validation.
//
//   // give local net.Addr, restart counter, channel to let background process pass the errors.
//   uConn, err := v1.ListenAndServe(laddr, 0, errCh)
//   if err != nil {
//   	// ...
//   }
//
// With UPlaneConn, you can ReadFromGTP() and WriteToGTP(), which gives you a easy handling of TEID and remote address.
//
// ReadFromGTP() reads from UPlaneConn, and returns the number of bytes copied into the given buffer(not including header), sender's net.Addr, incoming TEID set in GTP header, and error if occurred.
//
//   buf := make([]byte, 1500)
//   n, raddr, teid, err := uConn.ReadFromGTP(buf)
//   if err != nil {
//   	// ...
//   }
//
//   fmt.Printf("%x", buf[:n]) // prints the payload encapsulated in the GTP header.
//
// WriteToGTP() writes the payload encapsulated with GTP header to the specified addr over UPlaneConn.
//
//   // first return value is the number of bytes written.
//   if _, err := uConn.WriteToGTP(teid, payload, addr); err != nil {
//   	// ...
//   }
//
// For SGSN/S-GW-ish nodes, this package provides a special feature: Relay. It relays a packet coming from a UPlaneConn to another.
//
//   // s1Conn, s5Conn is UPlaneConn retrieved with Dial() or ListenAndServe().
//   relay := v1.NewRelay(s1Conn, s5Conn)
//
//   // associate incoming TEID on S1 with outgoing TEID and address on S5, and vice versa.
//   relay.AddPeer(s1TEIDIn, s5TEIDOut, s5Addr)
//   relay.AddPeer(s5TEIDIn, s1TEIDOut, s1Addr)
//
//   // make it start working by Run(), often good to work background.
//   // if no peer is registered, it just drops the packets.
//   go relay.Run()
//   defer relay.Close()
//
// Note: package v1 does provide encapsulation/decapsulation and some networking features,
// but it does not provide routing of the decapsulated packets, nor capturing IP layer and above on the specified interface. This is because such kind of operations cannot be done without platform-specific codes.
package v1
