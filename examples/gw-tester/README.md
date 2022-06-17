# GW Tester

A pseudo eNB and MME as a tester for S/P-GW.

## What's this?

![diagram](./docs/diagram.png)

It is a burden to use actual UE/eNB/MME just to test S/P-GW, isn't it?  
GW Tester emulates the minimal required behavior of surrounding nodes to perform quick and simple testing on S/P-GW.

A blog post by the author is available [here](https://wmnsk.com/posts/20200116_gw-tester/) for those who are interested in :)  
_NOTE: Some of the blog post's configurations or codes might be no longer relevant in the current version._

## How it works

### Authentication

Nothing!  
Subscribers defined in the `enb.yml` file can immediately attach and use the created sessions. MME accepts any subscribers without an authentication procedure.
Communication over the S1-MME interface is done with protobuf/gRPC instead of the S1AP protocol.

```
=== AD ===
Looking for a Go package for LTE authentication?
MILENAGE algorithm implementation is available :)

https://github.com/wmnsk/milenage
```

### Gateway Selection

No DNS lookup by TAI/APN is implemented.
MME just chooses gateways according to the mapping of source IP ranges and GW's IPs defined in `mme.yml`. This behavior might be changed in the future.

### Session Establishment

MME exchanges the real GTPv2 session establishment messages like Create/Modify/Delete Session with S-GW.

* IP address assignment  
Currently, we use the IP address that is defined in `enb.yml`, and the one passed by P-GW is ignored. This behavior might be changed in the future to be more practical.

* TEID allocation  
It can be both static and dynamic. Random TEID will be allocated by enb if `i_tei` in `enb.yml` is set to `0`. For outgoing TEID, the one that S-GW allocates will be used.

### U-Plane Data Injection

eNB forwards incoming traffic from UE or generates traffic by itself depending on the `type` in `enb.yml`.
GTP-U feature is based on [Linux Kernel GTP-U](https://www.kernel.org/doc/Documentation/networking/gtp.txt) with netlink.

| type     | behavior                                                                                                        |
| -------- | --------------------------------------------------------------------------------------------------------------- |
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
Functional S-GW and P-GW are also available in the same directory if you need them.

```shell-session
go get github.com/wmnsk/go-gtp/examples/gw-tester/enb
go get github.com/wmnsk/go-gtp/examples/gw-tester/mme
```

And run them with YAML configuration. See [Configuration](#configurations) section for details.

```shell-session
./mme
```

```shell-session
./enb
```

Then you'll see;

* MME starts sending GTPv2 Create Session Request to S-GW after it receives subscriber information from eNB.
* When sessions are successfully created on S/P-GW, eNB sets up GTP-U tunnels with S-GW.

After the successful creation of the sessions, you can inject packets externally or generate them on eNB.

## Configurations

Each node has a YAML file as a configuration.  
In general, config consists of the network information of local/remote nodes and some node-specific parameters.

### eNB

#### Global

These values are used to identify eNB. Some of them are just to be set inside the packets and not validated.

| config           | type of value | description                                                       |
| ---------------- | ------------- | ----------------------------------------------------------------- |
| `mcc`            | string        | MCC of eNB                                                        |
| `mnc`            | string        | MNC of eNB                                                        |
| `rat_type`       | uint8         | RAT Type (`6` for E-UTRAN)                                        |
| `tai`            | uint16        | TAI of eNB                                                        |
| `eci`            | uint32        | ECI of eNB                                                        |
| `mme_addr`       | string        | IP/Port of MME to dial, for S1-MME interface                      |
| `use_kernel_gtp` | bool          | Use Kernel GTP or not. U-Plane does not work when set to `false`. |
| `prom_addr`      | string        | IP/Port of MME to serve Prometheus                                |

#### Local Addresses

`local_addresses` are the IP addresses/ports to be bound on the local machine.

| config   | type of value | description                   |
| -------- | ------------- | ----------------------------- |
| `s1c_ip` | string        | local IP for S1-MME interface |
| `s1u_ip` | string        | local IP for S1-U interface   |

#### Subscribers

`subscribers` are the list of subscribers to attach.

| config               | type of value | description                                                                                                                                                                     |
| -------------------- | ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `imsi`               | string        | IMSI of the subscriber                                                                                                                                                          |
| `msisdn`             | string        | MSISDN of the subscriber                                                                                                                                                        |
| `imeisv`             | string        | IMEISV of the subscriber                                                                                                                                                        |
| `src_ip`             | string        | source IP of the subscriber (not assigned by P-GW)                                                                                                                              |
| `i_tei`              | uint32        | incoming TEID that S-GW should specify the subscriber                                                                                                                           |
| `type`               | string        | `external` or `http_get`. see [U-Plane Data Injection](#u-plane-data-injection)                                                                                                 |
| `euu_if_name`        | string        | name of the network interface on eUu side.</br>type=`external`: Used to receive traffic from external UE</br>type=`http_get`: Used as a source interface that `src_ip` is added |
| `http_url`           | string        | URL to HTTP GET by built-in traffic generator                                                                                                                                   |
| `reattach_on_reload` | bool          | whether to perform the attach procedure again on config reload                                                                                                                  |

### MME

#### Global

These values are used to identify MME. Some of them are just to be set inside the packets and not validated.

| config      | type of value | description                        |
| ----------- | ------------- | ---------------------------------- |
| `mcc`       | string        | MCC of MME                         |
| `mnc`       | string        | MNC of MME                         |
| `apn`       | string        | APN assigned to all the subscriber |
| `prom_addr` | string        | IP/Port of MME to serve Prometheus |

#### Local Addresses

`local_addresses` are the IP addresses/ports to be bound on the local machine.

| config     | type of value | description                        |
| ---------- | ------------- | ---------------------------------- |
| `s1c_addr` | string        | local IP/Port for S1-MME interface |
| `s11_ip`   | string        | local IP for S11 interface         |

#### Gateway IPs

IP addresses required to know/tell S-GW. This is typically done by DNS lookup with APN, but it's static for now.

| config       | type of value | description                  |
| ------------ | ------------- | ---------------------------- |
| `sgw_s11_ip` | string        | S-GW's IP for S11 interface  |
| `pgw_s5c_ip` | string        | P-GW's IP for S5-C interface |

### S-GW

#### Local Addresses

`local_addresses` are the IP addresses/ports to be bound on the local machine.

| config           | type of value | description                         |
| ---------------- | ------------- | ----------------------------------- |
| `s11_ip`         | string        | local IP for S11 interface          |
| `s1u_ip`         | string        | local IP for S1-U interface         |
| `s5c_ip`         | string        | local IP for S5-C interface         |
| `s5u_ip`         | string        | local IP for S5-U interface         |
| `use_kernel_gtp` | bool          | Use Kernel GTP or not.              |
| `prom_addr`      | string        | IP/Port of S-GW to serve Prometheus |

### P-GW

#### Global

| config           | type of value | description                                                                |
| ---------------- | ------------- | -------------------------------------------------------------------------- |
| `sgi_if_name`    | string        | name of the network interface on SGi side. Used to downlink route traffic. |
| `route_subnet`   | string        | IP subnet of UEs that should be routed properly.                           |
| `use_kernel_gtp` | bool          | Use Kernel GTP or not. U-Plane does not work when set to `false`.          |
| `prom_addr`      | string        | IP/Port of P-GW to serve Prometheus                                        |

#### Local Addresses

`local_addresses` are the IP addresses/ports to be bound on the local machine.

| config   | type of value | description                 |
| -------- | ------------- | --------------------------- |
| `s5c_ip` | string        | local IP for S5-C interface |
| `s5u_ip` | string        | local IP for S5-U interface |
| `sgi_ip` | string        | local IP for SGi interface  |

## Other Features

### Reloading config

The programs can handle `SIGHUP` to reload config without deleting sessions. Update YAML file and send `SIGHUP` to the process.

### Instrumentation

GW Tester nodes expose some metrics for Prometheus if `prom_addr` is given in each config. You can see the sample response from each node in [this Gist](https://gist.github.com/wmnsk/72f6d2d2450452090cd6351ffe63f660).  
I'm planning to add some more metrics like "success rate of HTTP probe", etc.

| Metrics           | Name                                  | Description                                   |
| ----------------- | ------------------------------------- | --------------------------------------------- |
| Active sessions   | `<node-name>_active_sessions`         | number of session established currently       |
| Active bearers    | `<node-name>_active_bearers`          | number of GTP-U tunnels established currently |
| Messages sent     | `<node-name>_messages_sent_total`     | number of messages sent by messagge type      |
| Messages received | `<node-name>_messages_received_total` | number of messages received by messagge type  |
