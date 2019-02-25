# v1: GTPv1 in Golang

Package v1 provides the simple and painless handling of GTPv2-C protocol in pure Golang.

## Getting Started

This package is still under construction. The networking feature is only available for GTPv1-U. GTPv1-C feature would be available in the future.
See messages and ies directory for what you can do with the current implementation. 

### Creating a PDP Context as a client

_NOT IMPLEMENTED YET!_

### Waiting for a PDP Context to be created as a server

_NOT IMPLEMENTED YET!_

### Opening a U-Plane connection

Use `Dial()` or `ListenAndServe()` to retrieve `UPlaneConn`.The difference between the two functions is;

* `Dial()` sends Echo Request and returns `UPlaneConn` if it succeeds.

```go
// give local/remote net.Addr, restart counter, channel to let background process pass the errors.
uConn, err := v1.Dial(laddr, raddr, 0, errCh)
if err != nil {
    // ...
}
```

* `ListenAndServe()` just returns `UPlaneConn` without any validation.

```go
// give local net.Addr, restart counter, channel to let background process pass the errors.
uConn, err := v1.ListenAndServe(laddr, 0, errCh)
if err != nil {
    // ...
}
```

With `UPlaneConn`, you can `ReadFromGTP()` and `WriteToGTP()`, which gives you a easy handling of TEID and remote address.

* `ReadFromGTP()` reads from `UPlaneConn`, and returns the number of bytes copied into the given buffer(not including header), sender's net.Addr, incoming TEID set in GTP header, and error if occurred.

```go
buf := make([]byte, 1500)
n, raddr, teid, err := uConn.ReadFromGTP(buf)
if err != nil {
    // ...
}

fmt.Printf("%x", buf[:n]) // prints the payload encapsulated in the GTP header.
```

* `WriteToGTP()` writes the payload encapsulated with GTP header to the specified addr over `UPlaneConn`.

```go
// first return value is the number of bytes written.
if _, err := uConn.WriteToGTP(teid, payload, addr); err != nil {
    // ...
}
```

For SGSN/S-GW-ish nodes, this package provides a special feature: `Relay`. It relays a packet coming from a `UPlaneConn` to another.

```go
// s1Conn, s5Conn is UPlaneConn retrieved with Dial() or ListenAndServe().
relay := v1.NewRelay(s1Conn, s5Conn)

// associate incoming TEID on S1 with outgoing TEID and address on S5, and vice versa.
relay.AddPeer(s1TEIDIn, s5TEIDOut, s5Addr)
relay.AddPeer(s5TEIDIn, s1TEIDOut, s1Addr)

// make it start working by Run(), often good to work background.
// if no peer is registered, it just drops the packets.
go relay.Run()
defer relay.Close()
```

Note: _package v1 does provide encapsulation/decapsulation and some networking features, but it does not provide routing of the decapsulated packets, nor capturing IP layer and above on the specified interface. This is because such kind of operations cannot be done without platform-specific codes._

## Supported Features

### Messages

The following Messages marked with "Yes" are currently available with their own useful constructors.

_Even there are some missing Messages, you can create any kind of Message by using `messages.NewGeneric()`._

| ID        | Name                                        | Supported |
|-----------|---------------------------------------------|-----------|
| 0         | (Spare/Reserved)                            | -         |
| 1         | Echo Request                                | Yes       |
| 2         | Echo Response                               | Yes       |
| 3         | Version Not Supported                       | Yes       |
| 4         | Node Alive Request                          |           |
| 5         | Node Alive Response                         |           |
| 6         | Redirection Request                         |           |
| 7         | Redirection Response                        |           |
| 8-15      | (Spare/Reserved)                            | -         |
| 16        | Create PDP Context Request                  | Yes       |
| 17        | Create PDP Context Response                 | Yes       |
| 18        | Update PDP Context Request                  | Yes       |
| 19        | Update PDP Context Response                 | Yes       |
| 20        | Delete PDP Context Request                  | Yes       |
| 21        | Delete PDP Context Response                 | Yes       |
| 22        | Initiate PDP Context Activation Request     |           |
| 23        | Initiate PDP Context Activation Response    |           |
| 24-25     | (Spare/Reserved)                            | -         |
| 26        | Error Indication                            | Yes       |
| 27        | PDU Notification Request                    |           |
| 28        | PDU Notification Response                   |           |
| 29        | PDU Notification Reject Request             |           |
| 30        | PDU Notification Reject Response            |           |
| 31        | Supported Extension Headers Notification    |           |
| 32        | Send Routeing Information for GPRS Request  |           |
| 33        | Send Routeing Information for GPRS Response |           |
| 34        | Failure Report Request                      |           |
| 35        | Failure Report Response                     |           |
| 36        | Note MS GPRS Present Request                |           |
| 37        | Note MS GPRS Present Response               |           |
| 38-47     | (Spare/Reserved)                            | -         |
| 48        | Identification Request                      |           |
| 49        | Identification Response                     |           |
| 50        | SGSN Context Request                        |           |
| 51        | SGSN Context Response                       |           |
| 52        | SGSN Context Acknowledge                    |           |
| 53        | Forward Relocation Request                  |           |
| 54        | Forward Relocation Response                 |           |
| 55        | Forward Relocation Complete                 |           |
| 56        | Relocation Cancel Request                   |           |
| 57        | Relocation Cancel Response                  |           |
| 58        | Forward SRNS Context                        |           |
| 59        | Forward Relocation Complete Acknowledge     |           |
| 60        | Forward SRNS Context Acknowledge            |           |
| 61        | UE Registration Query Request               |           |
| 62        | UE Registration Query Response              |           |
| 63-69     | (Spare/Reserved)                            | -         |
| 70        | RAN Information Relay                       |           |
| 71-95     | (Spare/Reserved)                            | -         |
| 96        | MBMS Notification Request                   |           |
| 97        | MBMS Notification Response                  |           |
| 98        | MBMS Notification Reject Request            |           |
| 99        | MBMS Notification Reject Response           |           |
| 100       | Create MBMS Context Request                 |           |
| 101       | Create MBMS Context Response                |           |
| 102       | Update MBMS Context Request                 |           |
| 103       | Update MBMS Context Response                |           |
| 104       | Delete MBMS Context Request                 |           |
| 105       | Delete MBMS Context Response                |           |
| 106 - 111 | (Spare/Reserved)                            | -         |
| 112       | MBMS Registration Request                   |           |
| 113       | MBMS Registration Response                  |           |
| 114       | MBMS De-Registration Request                |           |
| 115       | MBMS De-Registration Response               |           |
| 116       | MBMS Session Start Request                  |           |
| 117       | MBMS Session Start Response                 |           |
| 118       | MBMS Session Stop Request                   |           |
| 119       | MBMS Session Stop Response                  |           |
| 120       | MBMS Session Update Request                 |           |
| 121       | MBMS Session Update Response                |           |
| 122-127   | (Spare/Reserved)                            | -         |
| 128       | MS Info Change Notification Request         |           |
| 129       | MS Info Change Notification Response        |           |
| 130-239   | (Spare/Reserved)                            | -         |
| 240       | Data Record Transfer Request                |           |
| 241       | Data Record Transfer Response               |           |
| 242-253   | (Spare/Reserved)                            | -         |
| 254       | End Marker                                  |           |
| 255       | G-PDU                                       | Yes       |

### Information Elements

The following Information Elements marked with "Yes" are currently supported with their own useful constructors.

_Even there are some missing IEs, you can create any kind of IEs by using `ies.New()` function or by initializing ies.IE directly._

| ID      | Name                                      | Supported |
|---------|-------------------------------------------|-----------|
| 0       | (Spare/Reserved)                          | -         |
| 1       | Cause                                     | Yes       |
| 2       | IMSI                                      | Yes       |
| 3       | Routeing Area Identity                    | Yes       |
| 4       | Temporary Logical Link Identity           |           |
| 5       | Packet TMSI                               | Yes       |
| 6       | (Spare/Reserved)                          | -         |
| 7       | (Spare/Reserved)                          | -         |
| 8       | Reordering Required                       | Yes       |
| 9       | Authentication Triplet                    | Yes       |
| 10      | (Spare/Reserved)                          | -         |
| 11      | MAP Cause                                 | Yes       |
| 12      | P-TMSI Signature                          | Yes       |
| 13      | MS Validated                              | Yes       |
| 14      | Recovery                                  | Yes       |
| 15      | Selection Mode                            | Yes       |
| 16      | TEID Data I                               | Yes       |
| 17      | TEID C-Plane                              | Yes       |
| 18      | TEID Data II                              | Yes       |
| 19      | Teardown Indication                       | Yes       |
| 20      | NSAPI                                     | Yes       |
| 21      | RANAP Cause                               | Yes       |
| 22      | RAB Context                               |           |
| 23      | Radio Priority SMS                        |           |
| 24      | Radio Priority                            |           |
| 25      | Packet Flow ID                            |           |
| 26      | Charging Characteristics                  |           |
| 27      | Trace Reference                           |           |
| 28      | Trace Type                                |           |
| 29      | MS Not Reachable Reason                   |           |
| 30-126  | (Spare/Reserved)                          | -         |
| 127     | Charging ID                               |           |
| 128     | End User Address                          | Yes       |
| 129     | MM Context                                |           |
| 130     | PDP Context                               |           |
| 131     | Access Point Name                         | Yes       |
| 132     | Protocol Configuration Options            | Yes       |
| 133     | GSN Address                               | Yes       |
| 134     | MSISDN                                    | Yes       |
| 135     | QoS Profile                               |           |
| 136     | Authentication Quintuplet                 | Yes       |
| 137     | Traffic Flow Template                     |           |
| 138     | Target Identification                     |           |
| 139     | UTRAN Transparent Container               |           |
| 140     | RAB Setup Information                     |           |
| 141     | Extension Header Type List                |           |
| 142     | Trigger Id                                |           |
| 143     | OMC Identity                              |           |
| 144     | RAN Transparent Container                 |           |
| 145     | PDP Context Prioritization                |           |
| 146     | Additional RAB Setup Information          |           |
| 147     | SGSN Number                               |           |
| 148     | Common Flags                              | Yes       |
| 149     | APN Restriction                           | Yes       |
| 150     | Radio Priority LCS                        |           |
| 151     | RAT Type                                  | Yes       |
| 152     | User Location Information                 | Yes       |
| 153     | MS Time Zone                              | Yes       |
| 154     | IMEISV                                    | Yes       |
| 155     | CAMEL Charging Information Container      |           |
| 156     | MBMS UE Context                           |           |
| 157     | Temporary Mobile Group Identity           |           |
| 158     | RIM Routing Address                       |           |
| 159     | MBMS Protocol Configuration Options       |           |
| 160     | MBMS Service Area                         |           |
| 161     | Source RNC PDCP Context Info              |           |
| 162     | Additional Trace Info                     |           |
| 163     | Hop Counter                               |           |
| 164     | Selected PLMN Id                          |           |
| 165     | MBMS Session Identifier                   |           |
| 166     | MBMS 2G/3G Indicator                      |           |
| 167     | Enhanced NSAPI                            |           |
| 168     | MBMS Session Duration                     |           |
| 169     | Additional MBMS Trace Info                |           |
| 170     | MBMS Session Repetition Number            |           |
| 171     | MBMS Time To Data Transfer                |           |
| 172     | (Spare/Reserved)                          | -         |
| 173     | BSS Container                             |           |
| 174     | Cell Identification                       |           |
| 175     | PDU Numbers                               |           |
| 176     | BSS GP Cause                              |           |
| 177     | Required MBMS Bearer Capabilities         |           |
| 178     | RIM Routing Address Discriminator         |           |
| 179     | List of Setup PFCs                        |           |
| 180     | PS Handover XID Parameters                |           |
| 181     | MS Info Change Reporting Action           |           |
| 182     | Direct Tunnel Flags                       |           |
| 183     | Correlation Id                            |           |
| 184     | Bearer Control Mode                       |           |
| 185     | MBMS Flow Identifier                      |           |
| 186     | MBMS IP Multicast Distribution            |           |
| 187     | MBMS Distribution Acknowledgement         |           |
| 188     | Reliable InterRAT Handover Info           |           |
| 189     | RFSP Index                                |           |
| 190     | Fully Qualified Domain Name               |           |
| 191     | Evolved Allocation Retention Priority I   |           |
| 192     | Evolved Allocation Retention Priority II  |           |
| 193     | Extended Common Flags                     |           |
| 194     | User CSG Information                      |           |
| 195     | CSG Information Reporting Action          |           |
| 196     | CSG ID                                    |           |
| 197     | CSG Membership Indication                 |           |
| 198     | Aggregate Maximum Bit Rate                |           |
| 199     | UE Network Capability                     |           |
| 200     | UE-AMBR                                   |           |
| 201     | APN-AMBR with NSAPI                       |           |
| 202     | GGSN Back-Off Time                        |           |
| 203     | Signalling Priority Indication            |           |
| 204     | Signalling Priority Indication with NSAPI |           |
| 205     | Higher Bitrates than 16Mbps Flag          |           |
| 206     | (Spare/Reserved)                          | -         |
| 207     | Additional MM Context for SRVCC           |           |
| 208     | Additional Flags for SRVCC                |           |
| 209     | STN-SR                                    |           |
| 210     | C-MSISDN                                  |           |
| 211     | Extended RANAP Cause                      |           |
| 212     | eNodeB ID                                 |           |
| 213     | Selection Mode with NSAPI                 |           |
| 214     | ULI Timestamp                             | Yes       |
| 215     | LHN Id with NSAPI                         |           |
| 216     | CN Operator Selection Entity              |           |
| 217     | UE Usage Type                             |           |
| 218     | Extended Common Flags II                  |           |
| 219     | Node Identifier                           |           |
| 220     | CIoT Optimizations Support Indication     |           |
| 221     | SCEF PDN Connection                       |           |
| 222     | IOV Updates Counter                       |           |
| 223-237 | (Spare/Reserved)                          | -         |
| 238     | Special IE Type for IE Type Extension     |           |
| 239-250 | (Spare/Reserved)                          | -         |
| 251     | Charging Gateway Address                  |           |
| 252-254 | (Spare/Reserved)                          | -         |
| 255     | Private Extension                         |           |
