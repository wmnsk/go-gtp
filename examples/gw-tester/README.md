# GW Tester

A pseudo eNB and MME as a tester for S/P-GW.  
This is a work in collaboration with CNCF Telecom User Group's CNF Testbed project.

## What's this?

It is a burden to use actual UE/eNB/MME just to test S/P-GW.
GW Tester emulates the minimal required behavior of surrounding nodes, which enables us to perform a quick and simple testing on S/P-GW.

## How it works

### Authentication

Nothing!  
Subscribers defined in [`enb.yml`]() file can immediately attach and use the created sessions. MME accepts any subscribers without authentication procedure.
Communication over S1-MME interface is done with protobuf/gRPC instead of S1AP protocol.

```
=== AD ===
Looking for library for LTE authentication?
MILENAGE algorithm implementation in Golang is available :)

https://github.com/wmnsk/milenage
```

### Gateway Selection

No DNS lookup by TAI/APN is implemented.
MME just chooses gateways according to the mapping of source IP ranges and GW's IPs defined in [`mme.yml`](). This behavior might be changed in the future.

### Session Establishment

MME exchanges the real GTPv2 session establishment messages like Create/Modify/Delete Session with S-GW.

* IP address assignment  
Currently we use the IP address that is defined in [`enb.yml`]() and the one passed by P-GW is ignored. This behavior might be changed in the future to be more practical.

* TEID allocation  
It can be both static and dynamic. Random TEID will be allocated by enb if `i_tei` in [`enb.yml`]() is set to `0`. For outgoing TEID, the one that is allocated by S-GW will be used.

### U-Plane Data Injection

eNB forwards incoming traffic from UE or generates traffic by itself depending on the `type` in [`enb.yml`]().
GTP-U feature is based on [Linux Kernel GTP-U](https://www.kernel.org/doc/Documentation/networking/gtp.txt) with netlink.

| type     | behavior                                                                                                        |
|----------|-----------------------------------------------------------------------------------------------------------------|
| external | eNB encapsulates and forwards the incoming packets from `src_ip` toward the specified interface(`euu_if_name`). |
| http_get | eNB starts sending HTTP GET to the specified URL.                                                               |
| ...      | _(other types might be implemented in the future!)_                                                             |

## Getting Started

### Prerequisites

* Linux (kernel >= 4.12) with root privilege
* `net.ip_forward` enabled
* ports opened: 36412/TCP, 2123/UDP, 2152/UDP

### Run testers

Just `go get` eNB and MME.

```shell-session
go get github.com/wmnsk/go-gtp/examples/gw-tester/enb
go get github.com/wmnsk/go-gtp/examples/gw-tester/mme
```

And run them with options you like. 

```shell-session
./mme -h
...
```

```shell-session
./enb -h
...
```

Then you'll see;

* MME starts sending GTPv2 Create Session Request to S-GW after it receives subscriber information from eNB.
* When session is successfully created on S/P-GW, eNB sets up GTP-U tunnels with S-GW.

## Features

### Automatic recovery

_(WIP)_

### Instrumentation

As a tester, it is important to monitor what's happening.
GW Tester implements direct instrumentation with Prometheus, which is available at `http://<addr:port>/metrics`.  

