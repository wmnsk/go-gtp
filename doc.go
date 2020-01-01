// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Package gtp provides simple and painless handling of GTP(GPRS Tunneling Protocol),
// implemented in the Go Programming Language.
//
// Examples for specific node are available in examples directory, which can be  as it is
// in the following way.
// As for the detailed usage as a package, see v0/v1/v2 directory for what you can do with
// the current implementation.
//
// 1. Open four terminals on the same machine and start capturing on loopback interface.
//
// 2. Start P-GW on terminal #1 and #2
//   // on terminal #1
//   ./pgw
//
//   // on terminal #2
//   ./pgw -s5c 127.0.0.53:2123 -s5u 127.0.0.5:2152
//
// 3. Start S-GW on terminal #3
//
//   // on terminal #3
//   ./sgw
//
// 4. Start MME on terminal #4
//
//   // on terminal #4
//   ./mme
//
// 5. You will see the nodes exchanging Create Session and Modify Bearer on C-Plane, and ICMP Echo on U-Plane afterwards.
package gtp
