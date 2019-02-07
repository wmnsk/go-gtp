# go-gtp: GTP in Golang

Package gtp provides simple and painless handling of GTP(GPRS Tunneling Protocol), implemented in the Go Programming Language.

[![CircleCI](https://circleci.com/gh/wmnsk/go-gtp.svg?style=svg&circle-token=ee1cf4324ad327802bb152dcb43e97cb4e984656)](https://circleci.com/gh/wmnsk/go-gtp)
[![GoDoc](https://godoc.org/github.com/wmnsk/go-gtp?status.svg)](https://godoc.org/github.com/wmnsk/go-gtp)
[![Go Report Card](https://goreportcard.com/badge/github.com/wmnsk/go-gtp)](https://goreportcard.com/report/github.com/wmnsk/go-gtp)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/wmnsk/go-gtp/blob/master/LICENSE)

## Features

* No platform-specific codes inside, so it **works almost everywhere Golang works**.
* Flexible enough to **control everything in GTP protocol**.;
  * For developing mobile core network nodes (see [examples](./gtp/examples)).
  * For developing testing tools like traffic simulators or fuzzers.
* Many **helpers kind to developers** provided, like session, bearer, and TEID associations.
* Easy handling of **multiple connections with fixed IP and Port** with UDP (or other `net.PacketConn`).

## Getting Started

### Prerequisites

The following packages should be installed before starting.

```shell-session
go get -u github.com/pkg/errors
go get -u github.com/google/go-cmp/cmp
go get -u github.com/pascaldekloe/goe/verify
```

### Running examples

Examples works as it is by `go build` and executing commands in the following way.

1. Open four terminals on a machine and start capturing on loopback interface.

2. Start P-GW on terminal #1 and #2
```shell-session
// on terminal #1
./pgw

// on terminal #2
./pgw -s5c 127.0.0.53:2123 -s5u 127.0.0.5:2152
```

3. Start S-GW on terminal #3

```shell-session
// on terminal #3
./sgw
```

4. Start MME on terminal #4

```shell-session
// on terminal #4
./mme
```

5. You will see the nodes exchanging Create Session and Modify Bearer on C-Plane, and ICMP Echo on U-Plane afterwards.

### Developing by your own

Each version has `net.PacketConn`-like APIs and GTP-specific ones which is often version-specific.
The basic idea behind the current implementation is;

* `Dial()` or `ListenAndServe()` to create a connection(`Conn`) between nodes.
* Register handlers to the `Conn` for specific messages with `AddHandler()`, which allows users to handle the messages coming from the remote endpoint as flexible as possible, with less pain.
* `CreateXXX()` to create session or PDP context with arbitrary IEs given. Session/PDP context is structured and they also have some helpers like `AddTEID()` to handle known TEID properly.

For the detailed usage of specific version, see README.md under each version's directory.

| Version | Details                       |
|---------|-------------------------------|
| GTPv0   | [README.md](gtp/v0/README.md) |
| GTPv1   | [README.md](gtp/v1/README.md) |
| GTPv2   | [README.md](gtp/v2/README.md) |

## Supported Features

Note that "supported" means that the package provides helpers which makes it easier to handle.
In other words, even if a message/IE is not marked as "Yes", you can make it work with some additional effort.

Your contribution is welcome to implement helpers for all the types, of course!

| Version           | Messages | IEs   | Networking (state machine)                           | Details                                                   |
|-------------------|----------|-------|------------------------------------------------------|-----------------------------------------------------------|
| GTPv0             | 35.7%    | 81.8% | not implemented yet                                  | [Supported Features](gtp/v0/README.md#supported-features) |
| GTPv1             | 24.1%    | 21.8% | v1-U is functional, <br> v1-C is not implemented yet | [Supported Features](gtp/v1/README.md#supported-features) |
| GTPv2             | 32.0%    | 43.2% | almost functional                                    | [Supported Features](gtp/v2/README.md#supported-features) |
| GTP' <br> (Prime) | N/A      | N/A   | N/A                                                  | _not planned_                                             |

## Disclaimer

This is still an experimental project. Any part of implementations(including exported APIs) may be changed before released as v1.0.0.

## Author(s)

Yoshiyuki Kurauchi ([Twitter](https://twitter.com/wmnskdmms) / [LinkedIn](https://www.linkedin.com/in/yoshiyuki-kurauchi/))

_I'm always open to welcome co-authors! Please feel free to talk to me._

## LICENSE

[MIT](https://github.com/wmnsk/go-gtp/blob/master/LICENSE)
