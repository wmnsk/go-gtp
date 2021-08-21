# v1: GTPv1 in Golang

Package v1 provides simple and painless handling of GTPv1-C and GTPv1-U protocols in pure Golang.

## Getting Started

This package is still under construction. The networking feature is only available for GTPv1-U. GTPv1-C feature would be available in the future.  
See message and ie directory for what you can do with the current implementation. 

### Creating a PDP Context as a client

_NOT IMPLEMENTED YET!_

### Waiting for a PDP Context to be created as a server

_NOT IMPLEMENTED YET!_

### Opening a U-Plane connection

Retrieve `UPlaneConn` first, using `DialUPlane` (for client) or `NewUPlaneConn` (for server). 

#### Client

`DialUPlane` sends Echo Request and returns `UPlaneConn` if it succeeds.
If you don't need Echo, see Server section. As GTP is UDP-based connection, and there are no session management on `UPlaneConn`, the behavior of `Dial` and `ListenAndServe` is not quite different.

```go
uConn, err := v1.Dial(ctx, laddr, raddr)
if err != nil {
	// ...
}
defer uConn.Close()
```

#### Server

Retrieve `UPlaneConn` with `NewUPlaneConn`, and `ListenAndServe` to start listening.

```go
uConn := v1.NewUPlaneConn(laddr)
if err != nil {
	// ...
}
defer uConn.Close()

// This blocks, and returns an error when it's fatal.
if err := uConn.ListenAndServe(ctx); err != nil {
	// ...
}
```

### Manupulating `UPlaneConn`

With `UPlaneConn`, you can add and delete tunnels, and manipulate device directly.

#### Using Linux Kernel GTP-U

Linux Kernel GTP-U is quite performant and easy to handle, but it requires root privilege, and of course it works only on Linux. So it is disabled by default. To get started, enable it first. Note that it cannot be disabled while the program is working.

```go
if err := uConn.EnableKernelGTP("gtp0", v1.roleSGSN); err != nil {
	// ...
}
```

Then, when the bearer information is ready, use `AddTunnel` or `AddTunnelOverride` to add a tunnel.  
The latter one deletes the existing tunnel with the same IP and/or incoming TEID before creating a tunnel,
while the former fails if there's any duplication.

```go
// add a tunnel by giving GTP peer's IP, subscriber's IP,
if err := uConn.AddTunnelOverride(
	net.ParseIP("10.10.10.10"), // GTP peer's IP
	net.ParseIP("1.1.1.1"),     // subscriber's IP
	0x55667788,                 // outgoing TEID
	0x11223344,                 // incoming TEID
); err != nil {
	// ...
}
```

When the tunnel is no longer necessary, use `DelTunnelByITEI` or `DelTunnelByMSAddress` to delete it.  
Or, by `Close`-ing the `UPlaneConn`, all the tunnels associated will the cleared.

```go
// delete a tunnel by giving an incoming TEID.
if err := uConn.DelTunnelByITEI(0x11223344); err != nil {
	// ...
}

// delete a tunnel by giving an IP address assigned to a subscriber.
if err := uConn.DelTunnelByMSAddress(net.ParseIP("1.1.1.1")); err != nil {
	// ...
}
```

The packets NOT forwarded by the Kernel can be handled automatically by giving a handler to `UPlaneConn`.  
Handlers for T-PDU, Echo Request/Response, and Error Indication are registered by default, but you can override them using `AddHandler`.

```go
uConn.AddHandler(message.MsgTypeEchoRequest, func(c v1.Conn, senderAddr net.Addr, msg message.Message) error {
	// do anything you want for Echo Request here.
	// errors returned here are passed to `errCh` that is given to UPlaneConn at the beginning.
	return nil
})
```

If the tunnel with appropriate IP or TEID is not found for a T-PDU packet, Kernel sends it to userland. You can manipulate it with `ReadFromGTP`.

```go
buf := make([]byte, 1500)

// the 3rd returned value is TEID in GTPv1-U Header.
n, raddr, teid, err := uConn.ReadFromGTP(buf)
if err != nil {
	// ...
}

fmt.Printf("%x", buf[:n]) // prints only the payload, no GTP header included.
```

Also, you can send any payload by using `WriteToGTP`. It writes the given payload with GTP header to the specified addr over `UPlaneConn`.

```go
// first return value is the number of bytes written.
if _, err := uConn.WriteToGTP(teid, payload, addr); err != nil {
	// ...
}
```

#### Using userland GTP-U

**Note:** _package v1 does provide the encapsulation/decapsulation and some networking features, but it does NOT provide routing of the decapsulated packets, nor capturing IP layer and above on the specified interface. This is because such kind of operations cannot be done without platform-specific codes._

You can use to `ReadFromGTP` read the packets coming into uConn. This does not work for the packets which are handled by `RelayTo`.

```go
buf := make([]byte, 1500)
n, raddr, teid, err := uConn.ReadFromGTP(buf)
if err != nil {
	// ...
}

fmt.Printf("%x", buf[:n]) // prints the payload encapsulated in the GTP header.
```

Also, you can send any payload by using `WriteToGTP`. It writes the given payload with GTP header to the specified addr over `UPlaneConn`.

```go
// first return value is the number of bytes written.
if _, err := uConn.WriteToGTP(teid, payload, addr); err != nil {
	// ...
}
```

Especially or SGSN/S-GW-ish nodes(=have multiple GTP tunnels and its raison d'être is just to forward traffic right to left/left to right) we provide a method to swap TEID and forward T-PDU packets automatically and efficiently.  
By using `RelayTo`, the `UPlaneConn` automatically handles the T-PDU packet in background with the least cost. Note that it's performed on the userland and thus it's not so performant.

```go
// this is the example for S-GW that completed establishing a session and ready to forward U-Plane packets.
s1uConn.RelayTo(s5uConn, s1usgwTEID, s5uBearer.OutgoingTEID, s5uBearer.RemoteAddress)
s5uConn.RelayTo(s1uConn, s5usgwTEID, s1uBearer.OutgoingTEID, s1uBearer.RemoteAddress)
```

### Handling Extension Headers

`AddExtensionHeaders` adds ExtensionHeader(s) to the Header of a Message, set the E flag, and checks if the types given are consistent (error will be returned if not).

```go
msg := message.NewTPDU(0x11223344, []byte{0xde, 0xad, 0xbe, 0xef})
if err := msg.AddExtensionHeaders(
	// We don't support construction of the specific type of an ExtensionHeader.
	// The second parameter should be the serialized bytes of contents.
	message.NewExtensionHeader(
		message.ExtHeaderTypeUDPPort,
		[]byte{0x22, 0xb8},
		message.ExtHeaderTypePDUSessionContainer,
	),
	message.NewExtensionHeader(
		message.ExtHeaderTypePDUSessionContainer,
		[]byte{0x00, 0xc2},
		message.ExtHeaderTypeNoMoreExtensionHeaders,
	),
); err != nil {
	// ...
}
```

ExtensionHeaders decoded or added are stored in `ExtensionHeaders` field in the Header, which can be accessed like this.

```go
// no need to write msg.Header.ExtensionHeaders, as the Header is embedded in messages.
for _, eh := range msg.ExtensionHeaders {
	log.Println(eh.Type)     // ExtensionHeader type has its own Type while it's not actually included in a packet. 
	log.Println(eh.Content)  // We do not support decoding of each type of content yet. Decode them on your own.
	log.Println(eh.NextType) // Don't sort the slice - it ruins the packet, or even cause a panic.
}
```

When you are directly manipulating a Header for some reason, `WithExtensionHeaders` would help you simplify your operation.
Be sure not to call it on a Message, as it returns `*Header`, not a `Message` interface.

```go
header := message.NewHeader(
	0x30, // no need to set E flag here - With... method will do that instead.
	message.MsgTypeEchoRequest,
	0xdeadbeef,
	0x00,
	[]byte{0xde, 0xad, 0xbe, 0xef},
).WithExtensionHeaders(
	message.NewExtensionHeader(
		message.ExtHeaderTypeUDPPort,
		[]byte{0x22, 0xb8},
		message.ExtHeaderTypePDUSessionContainer,
	),
	message.NewExtensionHeader(
		message.ExtHeaderTypePDUSessionContainer,
		[]byte{0x00, 0xc2},
		message.ExtHeaderTypeNoMoreExtensionHeaders,
	),
)
```

## Supported Features

### Messages

The following Messages marked with "Yes" are currently available with their own useful constructors.

_Even there are some missing Messages, you can create any kind of Message by using `message.NewGeneric`._

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
| 31        | Supported Extension Headers Notification    | Yes       |
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

_Even there are some missing IEs, you can create any kind of IEs by using `ie.New` function or by initializing ie.IE directly._

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
| 127     | Charging ID                               | Yes       |
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
| 141     | Extension Header Type List                | Yes       |
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
