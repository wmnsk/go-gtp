# v0: GTPv0 in Golang

Package v0 provides simple and painless handling of GTPv0 protocol in pure Golang.

## Getting Started

This package is still under construction.
See the source codes for what you can do with the current implementation. 

### Creating a PDP Context as a client

_NOT IMPLEMENTED YET!_

### Waiting for a PDP Context to be created as a server

_NOT IMPLEMENTED YET!_

### Opening a U-Plane connection

_NOT IMPLEMENTED YET!_

## Supported Features

The following Messages marked with "Yes" are currently available with their own useful constructors.

_Even there are some missing Messages, you can create any kind of Message by using `message.NewGeneric()`._

### Messages

| ID      | Name                                        | Supported |
|---------|---------------------------------------------|-----------|
| 0       | (Spare/Reserved)                            | -         |
| 1       | Echo Request                                | Yes       |
| 2       | Echo Response                               | Yes       |
| 3       | Version Not Supported                       |           |
| 4       | Node Alive Request                          |           |
| 5       | Node Alive Response                         |           |
| 6       | Redirection Request                         |           |
| 7       | Redirection Response                        |           |
| 8-15    | (Spare/Reserved)                            | -         |
| 16      | Create PDP Context Request                  | Yes       |
| 17      | Create PDP Context Response                 | Yes       |
| 18      | Update PDP Context Request                  | Yes       |
| 19      | Update PDP Context Response                 | Yes       |
| 20      | Delete PDP Context Request                  | Yes       |
| 21      | Delete PDP Context Response                 | Yes       |
| 22      | Create AA PDP Context Request               |           |
| 23      | Create AA PDP Context Response              |           |
| 24      | Delete AA PDP Context Request               |           |
| 25      | Delete AA PDP Context Response              |           |
| 26      | Error Indication                            |           |
| 27      | PDU Notification Request                    |           |
| 28      | PDU Notification Response                   |           |
| 29      | PDU Notification Reject Request             |           |
| 30      | PDU Notification Reject Response            |           |
| 31      | (Spare/Reserved)                            | -         |
| 32      | Send Routeing Information for GPRS Request  |           |
| 33      | Send Routeing Information for GPRS Response |           |
| 34      | Failure Report Request                      |           |
| 35      | Failure Report Response                     |           |
| 36      | Note MS GPRS Present Request                |           |
| 37      | Note MS GPRS Present Response               |           |
| 38-47   | (Spare/Reserved)                            | -         |
| 48      | Identification Request                      |           |
| 49      | Identification Response                     |           |
| 50      | SGSN Context Request                        |           |
| 51      | SGSN Context Response                       |           |
| 52      | SGSN Context Acknowledge                    |           |
| 53-239  | (Spare/Reserved)                            | -         |
| 240     | Data Record Transfer Request                |           |
| 241     | Data Record Transfer Response               |           |
| 242-254 | (Spare/Reserved)                            | -         |
| 255     | T-PDU                                       | Yes       |

### Information Elements

The following Information Elements marked with "Yes" are currently available with their own useful constructors.

_Even there are some missing IEs, you can create any kind of IEs by using `ie.New()` function or by initializing ie.IE directly._

| ID      | Name                                   | Supported |
|---------|----------------------------------------|-----------|
| 0       | (Spare/Reserved)                       | -         |
| 1       | Cause                                  | Yes       |
| 2       | IMSI                                   | Yes       |
| 3       | Routeing Area Identity (RAI)           | Yes       |
| 4       | Temporary Logical Link Identity (TLLI) | Yes       |
| 5       | Packet TMSI (P-TMSI)                   | Yes       |
| 6       | Quality of Service (QoS) Profile       | Yes       |
| 7       | (Spare/Reserved)                       | -         |
| 8       | Reordering Required                    | Yes       |
| 9       | Authentication Triplet                 |           |
| 10      | (Spare/Reserved)                       | -         |
| 11      | MAP Cause                              |           |
| 12      | P-TMSI Signature                       | Yes       |
| 13      | MS Validated                           |           |
| 14      | Recovery                               | Yes       |
| 15      | Selection mode                         | Yes       |
| 16      | Flow Label Data I                      | Yes       |
| 17      | Flow Label Signalling                  | Yes       |
| 18      | Flow Label Data II                     | Yes       |
| 19      | MS Not Reachable Reason                | Yes       |
| 20-126  | (Spare/Reserved)                       | -         |
| 127     | Charging ID                            | Yes       |
| 128     | End User Address                       | Yes       |
| 129     | MM Context                             |           |
| 130     | PDP Context                            |           |
| 131     | Access Point Name                      | Yes       |
| 132     | Protocol Configuration Options         |           |
| 133     | GSN Address                            | Yes       |
| 134     | MSISDN                                 | Yes       |
| 135-250 | (Spare/Reserved)                       | -         |
| 251     | Charging Gateway Address               | Yes       |
| 252-254 | (Spare/Reserved)                       | -         |
| 255     | Private Extension                      | Yes       |
