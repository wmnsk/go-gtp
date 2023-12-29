# go-gtp: GTP in Go

Package gtp provides simple and painless handling of GTP (GPRS Tunneling Protocol) in the Go programming language.

![CI status](https://github.com/wmnsk/go-gtp/actions/workflows/go.yml/badge.svg)
[![golangci-lint](https://github.com/wmnsk/go-gtp/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/wmnsk/go-gtp/actions/workflows/golangci-lint.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/wmnsk/go-gtp.svg)](https://pkg.go.dev/github.com/wmnsk/go-gtp)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/wmnsk/go-gtp/blob/master/LICENSE)

## Project Status

This project is still EXPERIMENTAL.  
Any part of the implementations (including exported APIs) may be changed before released as v1.0.0.

## Features

* Flexible enough to control everything in the GTP protocol, making it suitable for developing mobile core network nodes or testing tools for them.
* Provides many helpers that are kind to developers, such as session, bearer, and TEID management.
* Makes it easy to handle multiple connections with fixed IP and Port with UDP (or other `net.PacketConn`).
* Currently works only on Linux and macOS since netlink support is introduced. However, the plan is to make it work on Windows in the future.

## Getting Started

### Prerequisites

go-gtp supports Go Modules. Run `go mod tidy` in your project's directory to collect the required packages automatically.

```
go mod tidy
```

_This project follows [the Release Policy of Go](https://golang.org/doc/devel/release.html#policy)._

### Running examples

#### End-to-end

We have a set of tools called GW Tester at [`examples/gw-tester`](./examples/gw-tester). See the [document](./examples/gw-tester/README.md) for how it works and how to run it. It is also used for the integration test of this project. [Workflow setting](./.github/workflows/go.yml) may help you understand it as well.

#### Individual node

[The examples](examples/) work as it is by `go build` and executing commands in the following way.

*Note for macOS users*: before running any go service, make sure to execute `./mac_local_host_enabler.sh` you will find at [examples/utils](examples/utils).

1. Open four terminals on a machine and start capturing on the loopback interface.

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

_The "mme" is not an MME per se. In addition to S11 interface, it also mocks UEs and an eNB to establish sessions and send packets on U-Plane._

## Developing with go-gtp

This section briefly describes how to develop your own GTP node with go-gtp.
For the detailed usage of a specific version, see README.md under each version's directory.

| Version | Details                      |
| ------- | ---------------------------- |
| GTPv0   | [README.md](gtpv0/README.md) |
| GTPv1   | [README.md](gtpv1/README.md) |
| GTPv2   | [README.md](gtpv2/README.md) |

### Establishing a connection between nodes

Each version has `net.PacketConn`-like APIs and GTP-specific ones, which are often version-specific.
The basic idea behind the current implementation is;

* `Dial` or `ListenAndServe` to create a connection (`Conn`) between nodes.
* Register handlers to the `Conn` for specific messages with `AddHandler`, allowing users to handle the messages coming from the remote endpoint as flexible as possible, with less pain.
* `CreateXXX` to create session or PDP context with arbitrary IEs given. Session/PDP context is structured, and they also have some helpers like `AddTEID` to handle known TEID properly.

### Handling messages and IEs

#### Messages

All the messages implement the same interface: `message.Message`, and have their own structs named `<MessageName>`, which can be created by `New<MessageName>` with given `ie.IE`s. `message.Message` can be sent on top of `Conn` with `SendMessageTo`, or can be serialized into `[]byte` with `Marshal`.

To parse the message from `[]byte`, use `message.Parse`. The parsed message will be one of the structs that implement `message.Message`, and you can type-assert it to the corresponding struct to access the fields which are a set of `ie.IE`s.

#### IEs

All the IEs are of the same type: `ie.IE` (not an interface). An IE can be created either with `New<IEName>`, with `ie.New`, or with `ie.New<TypeIE>`. The latter two are useful when you want to create an IE with an unsupported type or our constructor does not work well for you.

To parse the IE from `[]byte`, use `ie.Parse` (note that `message.Parse` parses all the IEs on a message - you don't need to call `ie.Parse` when you're handling IEs on a message). The value of the parsed `ie.IE` can be retrieved with `<IEName>`, `ValueAs<IEType>`. Some of the complicated IEs have their own struct named `<IEName>Fields` to get the values by accessing the fields.

For grouped IEs, accesing the `ChildIEs` field and iterating over the list of IEs contained is the most efficient way in most cases. Though there are the methods to get the specific IE value from the list (e.g., `BearerFlags` can be called upon `BearerContext` IE), they are not recommended since they always parse the whole list of IEs again.

## Supported Features

Note that "supported" means that the package provides helpers that make it easier to handle.
In other words, even if a message/IE is not marked as "Yes", you can make it work with some additional effort.

Your contribution is welcome to implement helpers for all the types, of course!

| Version           | Messages | IEs  | Networking (state machine)                  | Details                                                  |
| ----------------- | -------- | ---- | ------------------------------------------- | -------------------------------------------------------- |
| GTPv0             | ~35%     | ~80% | not implemented yet                         | [gtpv0/README](gtpv0/README.md#supported-features) |
| GTPv1             | ~25%     | ~30% | v1-U: functional <br> v1-C: not implemented | [gtpv1/README](gtpv1/README.md#supported-features) |
| GTPv2             | ~40.0%   | ~45% | functional                                  | [gtpv2/README](gtpv2/README.md#supported-features) |
| GTP' <br> (Prime) | N/A      | N/A  | N/A                                         | _not planned_                                            |

_You may also be interested in the sibling project [go-pfcp](https://github.com/wmnsk/go-pfcp) which is a PFCP implementation in Go._

## Contributing

With the design goal of being flexible and easy to use, go-gtp is still in the early stage of development. Any contribution is welcome! Please feel free to open an issue or a pull request.

Please don't forget to run tests once you are done with your changes. Additionally, running the fuzz test is recommended to make sure that the implementation is robust enough.

```shell-session
go test ./...
go test -fuzz .
```

*Note for macOS users*: the first time you run any test, make sure to execute `./mac_local_host_enabler.sh` you will find at [examples/utils](examples/utils). You will have to run the script again after each reboot.

## Authors

[Yoshiyuki Kurauchi](https://wmnsk.com/) and [contributors](https://github.com/wmnsk/go-gtp/graphs/contributors).

## LICENSE

[MIT](https://github.com/wmnsk/go-gtp/blob/master/LICENSE)
