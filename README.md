# go-gtp: GTP in Golang

Package gtp provides simple and painless handling of GTP(GPRS Tunneling Protocol), implemented in the Go Programming Language.

[![CircleCI](https://circleci.com/gh/wmnsk/go-gtp.svg?style=shield)](https://circleci.com/gh/wmnsk/go-gtp)
[![GoDoc](https://godoc.org/github.com/wmnsk/go-gtp?status.svg)](https://godoc.org/github.com/wmnsk/go-gtp)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/wmnsk/go-gtp/blob/master/LICENSE)

## Features

* Flexible enough to **control everything in GTP protocol**.;
  * For developing mobile core network nodes (see [examples](./examples)).
  * For developing testing tools like traffic simulators or fuzzers.
* Many **helpers kind to developers** provided, like session, bearer, and TEID associations.
* Easy handling of **multiple connections with fixed IP and Port** with UDP (or other `net.PacketConn`).
* ~~No platform-specific codes inside, so it **works almost everywhere Golang works**.~~ _Currently it works only on Linux and macOS since netlink support is introduced. I'll make them separated from the base to let it work even on Windows in the future._

## Getting Started

### Prerequisites

go-gtp supports Go Modules. Just run go mod tidy in your project's directory to collect the required packages automatically.

```
go mod tidy
```

_This project follows [the Release Policy of Go](https://golang.org/doc/devel/release.html#policy)._

### Running examples

Examples works as it is by `go build` and executing commands in the following way.

1. Open four terminals on a machine and start capturing on loopback interface.

2. Start P-GW on terminal #1 and #2
```shell-session
// on terminal #1
./pgw

// on terminal #2
./pgw -s5c 127.0.0.53 -s5u 127.0.0.5
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

_If you want to see fewer number of subscribers, please comment-out the `v2.Subscriber` definitions in `example/mme/main.go`._

### Developing by your own

Each version has `net.PacketConn`-like APIs and GTP-specific ones which is often version-specific.
The basic idea behind the current implementation is;

* `Dial` or `ListenAndServe` to create a connection(`Conn`) between nodes.
* Register handlers to the `Conn` for specific messages with `AddHandler`, which allows users to handle the messages coming from the remote endpoint as flexible as possible, with less pain.
* `CreateXXX` to create session or PDP context with arbitrary IEs given. Session/PDP context is structured and they also have some helpers like `AddTEID` to handle known TEID properly.

For the detailed usage of specific version, see README.md under each version's directory.

| Version | Details                      |
|---------|------------------------------|
| GTPv0   | [README.md](gtpv0/README.md) |
| GTPv1   | [README.md](gtpv1/README.md) |
| GTPv2   | [README.md](gtpv2/README.md) |

## Supported Features

Note that "supported" means that the package provides helpers which makes it easier to handle.
In other words, even if a message/IE is not marked as "Yes", you can make it work with some additional effort.

Your contribution is welcome to implement helpers for all the types, of course!

| Version           | Messages | IEs   | Networking (state machine)                           | Details                                                  |
|-------------------|----------|-------|------------------------------------------------------|----------------------------------------------------------|
| GTPv0             | 35.7%    | 81.8% | not implemented yet                                  | [Supported Features](gtpv0/README.md#supported-features) |
| GTPv1             | 26.6%    | 30.1% | v1-U is functional, <br> v1-C is not implemented yet | [Supported Features](gtpv1/README.md#supported-features) |
| GTPv2             | 41.0%    | 43.2% | almost functional                                    | [Supported Features](gtpv2/README.md#supported-features) |
| GTP' <br> (Prime) | N/A      | N/A   | N/A                                                  | _not planned_                                            |

## Disclaimer

This is still an experimental project. Any part of implementations(including exported APIs) may be changed before released as v1.0.0.

## Author(s)

Yoshiyuki Kurauchi ([Twitter](https://twitter.com/wmnskdmms) / [LinkedIn](https://www.linkedin.com/in/yoshiyuki-kurauchi/))

_I'm always open to welcome co-authors! Please feel free to talk to me._

## LICENSE

[MIT](https://github.com/wmnsk/go-gtp/blob/master/LICENSE)
