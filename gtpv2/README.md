# gtpv2: GTPv2 in Golang

Package v2 provides simple and painless handling of GTPv2-C protocol in pure Golang.

## Getting Started

_Working examples are available in [example](../examples) directory, which might be the better instruction for developers._

### Opening a connection

#### Client

`Dial` opens a connection between the specified peer by confirming the peer is alive by Echo Request/Response exchange. If you don't need Echo, see Server section.

```go
// Note that the conn is not bound to raddr. to let a Conn to be able to communicate with multiple peers.
// Interface type is required here to let Conn know which is the local interface type.
conn, err := gtpv2.Dial(ctx, laddr, raddr, gtpv2.IFTypeS11MMEGTPC, 0)
if err != nil {
    // ...
}
```

#### Server 

Retrieve `Conn` with `NewConn`, and `ListenAndServe` to start listening.

```go
// Interface type is required here to let Conn know which is the local interface type.
srvConn := gtpv2.NewConn(srvAddr, gtpv2.IFTypeS11MMEGTPC, 0)
if err := srvConn.ListenAndServe(ctx); err != nil {
    // ...
}
```

### Handling incoming messages

Prepare functions that comform to [`HandlerFunc`](https://pkg.go.dev/github.com/wmnsk/go-gtp/gtpv2#Conn.AddHandler), and register them to `Conn` with `AddHandler`. This should be done as soon as you get `Conn` not to miss the incoming messages.

`HandlerFunc` is to handle the incoming messages by message type. See [example](../examples) for how it is like.  
Also consider using `AddHandlers` when you have many `HandlerFunc`s.

```go
// write what you expect to do on receiving a message. Handlers should be added per message type.
// by default, Echo Request/Response and Version Not Supported Indication is handled automatically.
conn.AddHandler(
    // first param is the type of message. give number in uint8 or use gtpgtpv2.MsgTypeXXX.
    messages.MsgTypeCreateSessionResponse,
    // second param is the HandlerFunc to describe how you handle the message coming from peer.
    func(c *gtpv2.Conn, senderAddr net.Addr, msg messages.Message) error {
        // Do what you want with CreateSessionResponse message here.
        // see examples directly for functional examples
    },
)
```

### Manipulating sessions

With `Conn`, you can create, modify, delete GTPv2-C sessions and bearers with the built-in methods.

#### Session creation

`CreateSession` creates a Session, while storing values given as IEs and creating a default bearer.
In the following call, for example, IMSI and TEID for S1-U eNB is stored in the created `Session`.

```go
session, err := c.CreateSession(
    // put IEs required for your implementation here.
    // it is easier to use constructors in ie package.
    ie.NewIMSI("123451234567890"),
    
    // or, you can use ie.New() to create an IE without type-specific constructor.
    // put the type of IE, flags/instance, and payload as the parameters.
    ie.New(ie.ExtendedTraceInformation, 0x00, []byte{0xde, 0xad, 0xbe, 0xef}),
    
    // to set the instance to IE created with message-specific constructor, WithInstance()
    // may be your help.
    ie.NewIMSI("123451234567890").WithInstance(1), // no one wants to set instance to IMSI, though.

    // don't forget to contain the Sender F-TEID, as it is used to distinguish the incoming
    // message by Conn.
    //
    // to be secure, TEID should be generated with random values, without conflicts in a Conn.
    // to achieve that, gtpv2 provides NewFTEID() which returns F-TEID in *ie.IE.
    s11Conn.NewFTEID(gtpv2.IFTypeS11MMEGTPC, mmeIP, ""),
)
if err != nil {
    // ...
}
```

#### Session deletion / Bearer modification

`DeleteSession` and `ModifyBearer` methods are provided to send each message as easy as possible.
Unlike `CreateSession`, they don't manipulate the Session information automatically.

### Opening a U-Plane connection

_See [v1/README.md](../gtpv1/README.md#opening-a-u-plane-connection)._

## Supported Features

The following Messages marked with "Yes" are currently available with their own useful constructors.

_Even there are some missing Messages, you can create any kind of Message by using `messages.NewGeneric()`._

### Messages

| ID      | Name                                            | Supported |
|---------|-------------------------------------------------|-----------|
| 0       | (Spare/Reserved)                                | -         |
| 1       | Echo Request                                    | Yes       |
| 2       | Echo Response                                   | Yes       |
| 3       | Version Not Supported Indication                | Yes       |
| 4-16    | (Spare/Reserved)                                | -         |
| 17-24   | (Spare/Reserved)                                | -         |
| 25-31   | (Spare/Reserved)                                | -         |
| 32      | Create Session Request                          | Yes       |
| 33      | Create Session Response                         | Yes       |
| 34      | Modify Bearer Request                           | Yes       |
| 35      | Modify Bearer Response                          | Yes       |
| 36      | Delete Session Request                          | Yes       |
| 37      | Delete Session Response                         | Yes       |
| 38      | Change Notification Request                     |           |
| 39      | Change Notification Response                    |           |
| 40      | Remote UE Report Notification                   |           |
| 41      | Remote UE Report Acknowledge                    |           |
| 42-63   | (Spare/Reserved)                                | -         |
| 64      | Modify Bearer Command                           | Yes       |
| 65      | Modify Bearer Failure Indication                | Yes       |
| 66      | Delete Bearer Command                           | Yes       |
| 67      | Delete Bearer Failure Indication                | Yes       |
| 68      | Bearer Resource Command                         |           |
| 69      | Bearer Resource Failure Indication              |           |
| 70      | Downlink Data Notification Failure Indication   | Yes       |
| 71      | Trace Session Activation                        |           |
| 72      | Trace Session Deactivation                      |           |
| 73      | Stop Paging Indication                          | Yes       |
| 74-94   | (Spare/Reserved)                                | -         |
| 95      | Create Bearer Request                           | Yes       |
| 96      | Create Bearer Response                          | Yes       |
| 97      | Update Bearer Request                           | Yes       |
| 98      | Update Bearer Response                          | Yes       |
| 99      | Delete Bearer Request                           | Yes       |
| 100     | Delete Bearer Response                          | Yes       |
| 101     | Delete PDN Connection Set Request               | Yes       |
| 102     | Delete PDN Connection Set Response              | Yes       |
| 103     | PGW Downlink Triggering Notification            |           |
| 104     | PGW Downlink Triggering Acknowledge             |           |
| 105-127 | (Spare/Reserved)                                | -         |
| 128     | Identification Request                          |           |
| 129     | Identification Response                         |           |
| 130     | Context Request                                 | Yes       |
| 131     | Context Response                                | Yes       |
| 132     | Context Acknowledge                             | Yes       |
| 133     | Forward Relocation Request                      |           |
| 134     | Forward Relocation Response                     |           |
| 135     | Forward Relocation Complete Notification        |           |
| 136     | Forward Relocation Complete Acknowledge         |           |
| 137     | Forward Access Context Notification             |           |
| 138     | Forward Access Context Acknowledge              |           |
| 139     | Relocation Cancel Request                       |           |
| 140     | Relocation Cancel Response                      |           |
| 141     | Configuration Transfer Tunnel                   |           |
| 142-148 | (Spare/Reserved)                                | -         |
| 149     | Detach Notification                             | Yes       |
| 150     | Detach Acknowledge                              | Yes       |
| 151     | CS Paging Indication                            |           |
| 152     | RAN Information Relay                           |           |
| 153     | Alert MME Notification                          |           |
| 154     | Alert MME Acknowledge                           |           |
| 155     | UE Activity Notification                        |           |
| 156     | UE Activity Acknowledge                         |           |
| 157     | ISR Status Indication                           |           |
| 158     | UE Registration Query Request                   |           |
| 159     | UE Registration Query Response                  |           |
| 160     | Create Forwarding Tunnel Request                |           |
| 161     | Create Forwarding Tunnel Response               |           |
| 162     | Suspend Notification                            | Yes       |
| 163     | Suspend Acknowledge                             | Yes       |
| 164     | Resume Notification                             | Yes       |
| 165     | Resume Acknowledge                              | Yes       |
| 166     | Create Indirect Data Forwarding Tunnel Request  |           |
| 167     | Create Indirect Data Forwarding Tunnel Response |           |
| 168     | Delete Indirect Data Forwarding Tunnel Request  |           |
| 169     | Delete Indirect Data Forwarding Tunnel Response |           |
| 170     | Release Access Bearers Request                  | Yes       |
| 171     | Release Access Bearers Response                 | Yes       |
| 172-175 | (Spare/Reserved)                                | -         |
| 176     | Downlink Data Notification                      | Yes       |
| 177     | Downlink Data Notification Acknowledge          | Yes       |
| 178     | (Spare/Reserved)                                | -         |
| 179     | PGW Restart Notification                        | Yes       |
| 180     | PGW Restart Notification Acknowledge            | Yes       |
| 181-199 | (Spare/Reserved)                                | -         |
| 200     | Update PDN Connection Set Request               | Yes       |
| 201     | Update PDN Connection Set Response              | Yes       |
| 202-210 | (Spare/Reserved)                                | -         |
| 211     | Modify Access Bearers Request                   | Yes       |
| 212     | Modify Access Bearers Response                  | Yes       |
| 213-230 | (Spare/Reserved)                                | -         |
| 231     | MBMS Session Start Request                      |           |
| 232     | MBMS Session Start Response                     |           |
| 233     | MBMS Session Update Request                     |           |
| 234     | MBMS Session Update Response                    |           |
| 235     | MBMS Session Stop Request                       |           |
| 236     | MBMS Session Stop Response                      |           |
| 237-239 | (Spare/Reserved)                                | -         |
| 240-247 | (Spare/Reserved)                                | -         |
| 248-255 | (Spare/Reserved)                                | -         |

### Information Elements

The following Information Elements marked with "Yes" are currently available with their own useful constructors.

_Even there are some missing IEs, you can create any kind of IEs by using `ie.New()` function or by initializing ie.IE directly._

| ID      | Name                                                           | Supported |
|---------|----------------------------------------------------------------|-----------|
| 0       | (Spare/Reserved)                                               | -         |
| 1       | International Mobile Subscriber Identity (IMSI)                | Yes       |
| 2       | Cause                                                          | Yes       |
| 3       | Recovery (Restart Counter)                                     | Yes       |
| 4-34    | (Spare/Reserved)                                               | -         |
| 35-50   | (Spare/Reserved)                                               | -         |
| 51      | STN-SR                                                         |           |
| 52-70   | (Spare/Reserved)                                               | -         |
| 71      | Access Point Name (APN)                                        | Yes       |
| 72      | Aggregate Maximum Bit Rate (AMBR)                              | Yes       |
| 73      | EPS Bearer ID (EBI)                                            | Yes       |
| 74      | IP Address                                                     | Yes       |
| 75      | Mobile Equipment Identity (MEI)                                | Yes       |
| 76      | MSISDN                                                         | Yes       |
| 77      | Indication                                                     | Yes       |
| 78      | Protocol Configuration Options (PCO)                           | Yes       |
| 79      | PDN Address Allocation (PAA)                                   | Yes       |
| 80      | Bearer Level Quality of Service (Bearer QoS)                   | Yes       |
| 81      | Flow Quality of Service (Flow QoS)                             | Yes       |
| 82      | RAT Type                                                       | Yes       |
| 83      | Serving Network                                                | Yes       |
| 84      | EPS Bearer Level Traffic Flow Template (Bearer TFT)            | Yes       |
| 85      | Traffic Aggregation Description (TAD)                          |           |
| 86      | User Location Information (ULI)                                | Yes       |
| 87      | Fully Qualified Tunnel Endpoint Identifier (F-TEID)            | Yes       |
| 88      | TMSI                                                           | Yes       |
| 89      | Global CN-Id                                                   | Yes       |
| 90      | S103 PDN Data Forwarding Info (S103PDF)                        | Yes       |
| 91      | S1-U Data Forwarding Info (S1UDF)                              | Yes       |
| 92      | Delay Value                                                    | Yes       |
| 93      | Bearer Context                                                 | Yes       |
| 94      | Charging ID                                                    | Yes       |
| 95      | Charging Characteristics                                       | Yes       |
| 96      | Trace Information                                              |           |
| 97      | Bearer Flags                                                   | Yes       |
| 98      | (Spare/Reserved)                                               | -         |
| 99      | PDN Type                                                       | Yes       |
| 100     | Procedure Transaction ID                                       | Yes       |
| 101     | (Spare/Reserved)                                               | -         |
| 102     | (Spare/Reserved)                                               | -         |
| 103     | MM Context (GSM Key and Triplets)                              |           |
| 104     | MM Context (UMTS Key, Used Cipher and Quintuplets)             |           |
| 105     | MM Context (GSM Key, Used Cipher and Quintuplets)              |           |
| 106     | MM Context (UMTS Key and Quintuplets)                          |           |
| 107     | MM Context (EPS Security Context, Quadruplets and Quintuplets) |           |
| 108     | MM Context (UMTS Key, Quadruplets and Quintuplets)             |           |
| 109     | PDN Connection                                                 |           |
| 110     | PDU Numbers                                                    |           |
| 111     | Packet TMSI                                                    | Yes       |
| 112     | P-TMSI Signature                                               | Yes       |
| 113     | Hop Counter                                                    | Yes       |
| 114     | UE Time Zone                                                   | Yes       |
| 115     | Trace Reference                                                | Yes       |
| 116     | Complete Request Message                                       |           |
| 117     | GUTI                                                           | Yes       |
| 118     | F-Container                                                    |           |
| 119     | F-Cause                                                        |           |
| 120     | PLMN ID                                                        | Yes       |
| 121     | Target Identification                                          |           |
| 122     | (Spare/Reserved)                                               | -         |
| 123     | Packet Flow ID                                                 |           |
| 124     | RAB Context                                                    |           |
| 125     | Source RNC PDCP Context Info                                   |           |
| 126     | Port Number                                                    | Yes       |
| 127     | APN Restriction                                                | Yes       |
| 128     | Selection Mode                                                 | Yes       |
| 129     | Source Identification                                          |           |
| 130     | (Spare/Reserved)                                               | -         |
| 131     | Change Reporting Action                                        |           |
| 132     | Fully Qualified PDN Connection Set Identifier (FQ-CSID)        | Yes       |
| 133     | Channel Needed                                                 |           |
| 134     | eMLPP Priority                                                 |           |
| 135     | Node Type                                                      | Yes       |
| 136     | Fully Qualified Domain Name (FQDN)                             | Yes       |
| 137     | Transaction Identifier (TI)                                    |           |
| 138     | MBMS Session Duration                                          |           |
| 139     | MBMS Service Area                                              |           |
| 140     | MBMS Session Identifier                                        |           |
| 141     | MBMS Flow Identifier                                           |           |
| 142     | MBMS IP Multicast Distribution                                 |           |
| 143     | MBMS Distribution Acknowledge                                  |           |
| 144     | RFSP Index                                                     |           |
| 145     | User CSG Information (UCI)                                     | Yes       |
| 146     | CSG Information Reporting Action                               |           |
| 147     | CSG ID                                                         | Yes       |
| 148     | CSG Membership Indication (CMI)                                | Yes       |
| 149     | Service Indicator                                              | Yes       |
| 150     | Detach Type                                                    | Yes       |
| 151     | Local Distinguished Name (LDN)                                 | Yes       |
| 152     | Node Features                                                  | Yes       |
| 153     | MBMS Time to Data Transfer                                     |           |
| 154     | Throttling                                                     | Yes       |
| 155     | Allocation/Retention Priority (ARP)                            | Yes       |
| 156     | EPC Timer                                                      | Yes       |
| 157     | Signalling Priority Indication                                 |           |
| 158     | Temporary Mobile Group Identity (TMGI)                         |           |
| 159     | Additional MM context for SRVCC                                |           |
| 160     | Additional flags for SRVCC                                     |           |
| 161     | (Spare/Reserved)                                               | -         |
| 162     | MDT Configuration                                              |           |
| 163     | Additional Protocol Configuration Options (APCO)               |           |
| 164     | Absolute Time of MBMS Data Transfer                            |           |
| 165     | H(e)NB Information Reporting                                   |           |
| 166     | IPv4 Configuration Parameters (IP4CP)                          |           |
| 167     | Change to Report Flags                                         |           |
| 168     | Action Indication                                              |           |
| 169     | TWAN Identifier                                                |           |
| 170     | ULI Timestamp                                                  | Yes       |
| 171     | MBMS Flags                                                     |           |
| 172     | RAN/NAS Cause                                                  | Yes       |
| 173     | CN Operator Selection Entity                                   |           |
| 174     | Trusted WLAN Mode Indication                                   |           |
| 175     | Node Number                                                    |           |
| 176     | Node Identifier                                                |           |
| 177     | Presence Reporting Area Action                                 |           |
| 178     | Presence Reporting Area Information                            |           |
| 179     | TWAN Identifier Timestamp                                      |           |
| 180     | Overload Control Information                                   |           |
| 181     | Load Control Information                                       |           |
| 182     | Metric                                                         |           |
| 183     | Sequence Number                                                |           |
| 184     | APN and Relative Capacity                                      |           |
| 185     | WLAN Offloadability Indication                                 |           |
| 186     | Paging and Service Information                                 | Yes       |
| 187     | Integer Number                                                 | Yes       |
| 188     | Millisecond Time Stamp                                         |           |
| 189     | Monitoring Event Information                                   |           |
| 190     | ECGI List                                                      |           |
| 191     | Remote UE Context                                              |           |
| 192     | Remote User ID                                                 |           |
| 193     | Remote UE IP information                                       |           |
| 194     | CIoT Optimizations Support Indication                          |           |
| 195     | SCEF PDN Connection                                            |           |
| 196     | Header Compression Configuration                               |           |
| 197     | Extended Protocol Configuration Options (ePCO)                 |           |
| 198     | Serving PLMN Rate Control                                      |           |
| 199     | Counter                                                        |           |
| 200     | Mapped UE Usage Type                                           |           |
| 201     | Secondary RAT Usage Data Report                                |           |
| 202     | UP Function Selection Indication Flags                         |           |
| 203     | Maximum Packet Loss Rate                                       |           |
| 204     | APN Rate Control Status                                        |           |
| 205     | Extended Trace Information                                     |           |
| 206     | Monitoring Event Extension Information                         |           |
| 207     | Additional RRM Policy Index                                    |           |
| 208     | V2X Context                                                    |           |
| 209     | PC5 QoS Parameters                                             |           |
| 210     | Services Authorized                                            |           |
| 211     | Bit Rate                                                       |           |
| 212     | PC5 QoS Flow                                                   |           |
| 213-253 | (Spare/Reserved)                                               | -         |
| 254     | (Spare/Reserved)                                               | -         |
| 255     | Private Extension                                              | Yes       |
