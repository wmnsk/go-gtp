# v2: GTPv2 in Golang

Package v2 provides the simple and painless handling of GTPv2-C protocol in pure Golang.

## Getting Started

_Working examples are available in [example](../examples) directory, which might be better instruction._

### Creating a Session as a client

Use `Dial()`, `AddHandler()`, `CreateSession()`, and you can get `*Conn`, `*Session` and `*Bearer`.

1. `Dial()` to retrieve *v2.Conn


```go
// give local/remote net.Addr, restart counter, channel to let background process pass the errors.
conn, err := v2.Dial(laddr, raddr, 0, errCh)
if err != nil {
    // ...
}
```

2. `AddHandler()` to register your own handler before creating session.

```go
// write what you expect to do on receiving a message. Handlers should be added per message type.
// by default, Echo Request/Response and Version Not Supported Indication is handled automatically.
conn.AddHandler(
    // first param is the type of message. give number in uint8 or use v2.MsgTypeXXX.
    messages.MsgTypeCreateSessionResponse,
    // second param is the HandlerFunc to describe how you handle the message coming from peer.
    func(c *v2.Conn, senderAddr net.Addr, msg messages.Message) error {
        // GetSessionByTEID helps you get the relevant Session(=created when you run CreateSession()).
        session, err := c.GetSessionByTEID(msg.TEID())
        if err != nil {
            c.RemoveSession(session)
            return err
        }
        // GetDefaultBearer() helps you get the default bearer.
        // to get other bearers, use GetBearerByName("name"), or GetBearerByEBI(ebi).
        bearer := session.GetDefaultBearer()

        // assert type to refer to the struct field specific to the message.
        // in general, no need to check if it can be type-asserted, as long as the MessageType is
        // specified correctly in AddHandler().
        csRsp := msg.(*messages.CreateSessionResponse)

        // all struct fields(except Header) are typed as *ies.IE, and there are the helpers methods
        // to retrieve the value from an IE's payload.
        // it's important to confirm the IE is not nil, as the other endpoint does not necessarily
        // contain the IE you expect.
        if ie := csRsp.Cause; ie != nil {
            if cause := ie.Cause(); cause != v2.CauseRequestAccepted {
                // before returning on failure, RemoveSession() to delete if it's no longer used.
                c.RemoveSession(session)
                // some errors expected to be used so often is available in v2/errors.go.
                return &v2.ErrCauseNotOK{
                    MsgType: csRsp.MessageTypeName(),
                    Cause:   cause,
                    Msg:     fmt.Sprintf("subscriber: %s", session.IMSI),
                }
            }
        } else {
            // if the missing IE is required to proceed, returns error.
            c.RemoveSession(session)
            return &v2.ErrRequiredIEMissing{Type: msg.MessageType()}
        }

        // do not forget to add TEID to Session by AddTEID() when you receive F-TEID.
        if ie := csRsp.SenderFTEIDC; ie != nil {
            session.AddTEID(ie.InterfaceType(), ie.TEID())
        } else {
            return &v2.ErrRequiredIEMissing{Type: ies.FullyQualifiedTEID}
        }
        
        // IEs inside grouped IE can be handled by ranging over ie.ChildIEs.
        // also, grouped IE has FindByType(), but it might be slower.
        if brCtxIE := csRsp.BearerContextsCreated; brCtxIE != nil {
            for _, ie := range brCtxIE.ChildIEs {
                switch ie.Type {
                case ies.EPSBearerID:
                    bearer.EBI = ie.EPSBearerID()
                case ies.FullyQualifiedTEID:
                    if ie.Instance() != 0 {
                        continue
                    }
                    // do not forget to add TEID to Session by AddTEID() when you receive F-TEID.
                    session.AddTEID(ie.InterfaceType(), ie.TEID())
                }    
            }
        } else {
            return &v2.ErrRequiredIEMissing{Type: ies.BearerContext}
        }
        
        // if Session is ready, let's active it.
        if err := session.Activate(); err != nil {
            c.RemoveSession(session)
            return err
        }

    },
)

// default handlers can be overridden just by specifying its type and giving a HandlerFunc.
conn.AddHandler(
    messages.MsgTypeEchoResponse,
    func(c *v2.Conn, senderAddr net.Addr, msg messages.Message) error {
        log.Printf("Got %s from %s", msg.MessageTypeName(), senderAddr)
        // do something special for Echo Response.
    },
)
```

3. `CreateSession()` to start creating a Session.

```go
// CreateSession() sends Create Session Request with given IEs, and stores information
// inside Session returned.
session, err := c.CreateSession(
    // put IEs required for your implementation here.
    // it is easier to use constructors in ies package.
    ies.NewIMSI("123451234567890"),
    // or, you can use ies.New() to create an IE without type-specific constructor.
    // put the type of IE, flags/instance, and payload as the parameters.
    ies.New(ies.ExtendedTraceInformation, 0x00, []byte{0xde, 0xad, 0xbe, 0xef}),
    // to set the instance to IE created with message-specific constructor, WithInstance()
    // may be your help.
    ies.NewIMSI("123451234567890").WithInstance(1), // no one wants to set instance to IMSI, though.
    // to be secure, TEID should be generated with random values, without conflicts in a Conn.
    // to achieve that, v2 provides NewFTEID() which returns F-TEID in *ies.IE.
    s11Conn.NewFTEID(v2.IFTypeS1UeNodeBGTPU, enbIP, ""),
)
if err != nil {
    // ...
}
// do not forget to add session to *Conn.
// do not Activate() it before you confirm the remote endpoint accepted the request.
c.AddSession(session)
```

### Waiting for a Session to be created as a server

Use `ListenAndServe()`, `AddHandler()`, and you can get `*Conn`, `*Session`, and `*Bearer`.

1. `ListenAndServe()` to retrieve *v2.Conn and start listening.

```go
// give local net.Addr, restart counter, channel to let background process pass the errors.
conn, err := v2.ListenAndServe(laddr, 0, errCh)
if err != nil {
    // ...
}
```

2. `AddHandler()` to register your own handler in the same way as previous section.

When adding handler for server, the following things should be taken into account;

* `Session` should be created by your own with `NewSession()`, and the subscriber/bearer information should be set properly(which is often in the request message).

* Response with error should be sent before returning with failure.

### Opening a U-Plane connection

_See [v1/README.md](../v1/README.md)._

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
| 64      | Modify Bearer Command                           |           |
| 65      | Modify Bearer Failure Indication                |           |
| 66      | Delete Bearer Command                           |           |
| 67      | Delete Bearer Failure Indication                |           |
| 68      | Bearer Resource Command                         |           |
| 69      | Bearer Resource Failure Indication              |           |
| 70      | Downlink Data Notification Failure Indication   |           |
| 71      | Trace Session Activation                        |           |
| 72      | Trace Session Deactivation                      |           |
| 73      | Stop Paging Indication                          |           |
| 74-94   | (Spare/Reserved)                                | -         |
| 95      | Create Bearer Request                           |           |
| 96      | Create Bearer Response                          |           |
| 97      | Update Bearer Request                           |           |
| 98      | Update Bearer Response                          |           |
| 99      | Delete Bearer Request                           | Yes       |
| 100     | Delete Bearer Response                          | Yes       |
| 101     | Delete PDN Connection Set Request               |           |
| 102     | Delete PDN Connection Set Response              |           |
| 103     | PGW Downlink Triggering Notification            |           |
| 104     | PGW Downlink Triggering Acknowledge             |           |
| 105-127 | (Spare/Reserved)                                | -         |
| 128     | Identification Request                          |           |
| 129     | Identification Response                         |           |
| 130     | Context Request                                 |           |
| 131     | Context Response                                |           |
| 132     | Context Acknowledge                             |           |
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
| 149     | Detach Notification                             |           |
| 150     | Detach Acknowledge                              |           |
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
| 162     | Suspend Notification                            |           |
| 163     | Suspend Acknowledge                             |           |
| 164     | Resume Notification                             |           |
| 165     | Resume Acknowledge                              |           |
| 166     | Create Indirect Data Forwarding Tunnel Request  |           |
| 167     | Create Indirect Data Forwarding Tunnel Response |           |
| 168     | Delete Indirect Data Forwarding Tunnel Request  |           |
| 169     | Delete Indirect Data Forwarding Tunnel Response |           |
| 170     | Release Access Bearers Request                  |           |
| 171     | Release Access Bearers Response                 |           |
| 172-175 | (Spare/Reserved)                                | -         |
| 176     | Downlink Data Notification                      |           |
| 177     | Downlink Data Notification Acknowledge          |           |
| 178     | (Spare/Reserved)                                | -         |
| 179     | PGW Restart Notification                        |           |
| 180     | PGW Restart Notification Acknowledge            |           |
| 181-199 | (Spare/Reserved)                                | -         |
| 200     | Update PDN Connection Set Request               |           |
| 201     | Update PDN Connection Set Response              |           |
| 202-210 | (Spare/Reserved)                                | -         |
| 211     | Modify Access Bearers Request                   |           |
| 212     | Modify Access Bearers Response                  |           |
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

_Even there are some missing IEs, you can create any kind of IEs by using `ies.New()` function or by initializing ies.IE directly._

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
| 84      | EPS Bearer Level Traffic Flow Template (Bearer TFT)            |           |
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
| 111     | P-TMSI                                                         |           |
| 112     | P-TMSI Signature                                               |           |
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
| 152     | Node Features                                                  |           |
| 153     | MBMS Time to Data Transfer                                     |           |
| 154     | Throttling                                                     |           |
| 155     | Allocation/Retention Priority (ARP)                            |           |
| 156     | EPC Timer                                                      |           |
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
| 186     | Paging and Service Information                                 |           |
| 187     | Integer Number                                                 |           |
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
| 206-253 | (Spare/Reserved)                                               | -         |
| 254     | (Spare/Reserved)                                               | -         |
| 255     | Private Extension                                              | Yes       |
