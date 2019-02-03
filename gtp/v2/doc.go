// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Package v2 provides the simple and painless handling of GTPv2-C protocol in pure Golang.
//
// NOTE: Working examples are available in example directory, which might be a better help.
//
// To creatie a Session as a client, use Dial(), AddHandler(), CreateSession(), and you can get *Conn, *Session and *Bearer.
//
// 1. Dial() to retrieve *v2.Conn
//
//   // give local/remote net.Addr, restart counter, channel to let background process pass the errors.
//   conn, err := v2.Dial(laddr, raddr, 0, errCh)
//   if err != nil {
//   	// ...
//   }
//
// 2. AddHandler() to register your own handler before creating session.
//
//   // write what you expect to do on receiving a message. Handlers should be added per message type.
//   // by default, Echo Request/Response and Version Not Supported Indication is handled automatically.
//   conn.AddHandler(
//   	// first param is the type of message. give number in uint8 or use v2.MsgTypeXXX.
//   	messages.MsgTypeCreateSessionResponse,
//   	// second param is the HandlerFunc to describe how you handle the message coming from peer.
//   	func(c *v2.Conn, senderAddr net.Addr, msg messages.Message) error {
//   		// GetSessionByTEID helps you get the relevant Session(=created when you run CreateSession()).
//   		session, err := c.GetSessionByTEID(msg.TEID())
//   		if err != nil {
//   			c.RemoveSession(session)
//   			return err
//   		}
//   		// GetDefaultBearer() helps you get the default bearer.
//   		// to get other bearers, use GetBearerByName("name"), or GetBearerByEBI(ebi).
//   		bearer := session.GetDefaultBearer()
//
//   		// assert type to refer to the struct field specific to the message.
//   		// in general, no need to check if it can be type-asserted, as long as the MessageType is
//   		// specified correctly in AddHandler().
//   		csRsp := msg.(*messages.CreateSessionResponse)
//
//   		// all struct fields(except Header) are typed as *ies.IE, and there are the helpers methods
//   		// to retrieve the value from an IE's payload.
//   		// it's important to confirm the IE is not nil, as the other endpoint does not necessarily
//   		// contain the IE you expect.
//   		if ie := csRsp.Cause; ie != nil {
//   			if cause := ie.Cause(); cause != v2.CauseRequestAccepted {
//   				// before returning on failure, RemoveSession() to delete if it's no longer used.
//   				c.RemoveSession(session)
//   				// some errors expected to be used so often is available in v2/errors.go.
//   				return &v2.ErrCauseNotOK{
//   					MsgType: csRsp.MessageTypeName(),
//   					Cause:   cause,
//   					Msg:     fmt.Sprintf("subscriber: %s", session.IMSI),
//   				}
//   			}
//   		} else {
//   			// if the missing IE is required to proceed, returns error.
//   			c.RemoveSession(session)
//   			return &v2.ErrRequiredIEMissing{Type: msg.MessageType()}
//   		}
//
//   		// do not forget to add TEID to Session by AddTEID() when you receive F-TEID.
//   		if ie := csRsp.SenderFTEIDC; ie != nil {
//   			session.AddTEID(ie.InterfaceType(), ie.TEID())
//   		} else {
//   			return &v2.ErrRequiredIEMissing{Type: ies.FullyQualifiedTEID}
//   		}
//
//   		// IEs inside grouped IE can be handled by ranging over ie.ChildIEs.
//   		// also, grouped IE has FindByType(), but it might be slower.
//   		if brCtxIE := csRsp.BearerContextsCreated; brCtxIE != nil {
//   			for _, ie := range brCtxIE.ChildIEs {
//   				switch ie.Type {
//   				case ies.EPSBearerID:
//   					bearer.EBI = ie.EPSBearerID()
//   				case ies.FullyQualifiedTEID:
//   					if ie.Instance() != 0 {
//   						continue
//   					}
//   					// do not forget to add TEID to Session by AddTEID() when you receive F-TEID.
//   					session.AddTEID(ie.InterfaceType(), ie.TEID())
//   				}
//   			}
//   		} else {
//   			return &v2.ErrRequiredIEMissing{Type: ies.BearerContext}
//   		}
//
//   		// if Session is ready, let's active it.
//   		if err := session.Activate(); err != nil {
//   			c.RemoveSession(session)
//   			return err
//   		}
//   	},
//   )
//
//   // default handlers can be overridden just by specifying its type and giving a HandlerFunc.
//   conn.AddHandler(
//   	messages.MsgTypeEchoResponse,
//   	func(c *v2.Conn, senderAddr net.Addr, msg messages.Message) error {
//   		log.Printf("Got %s from %s", msg.MessageTypeName(), senderAddr)
//   		// do something special for Echo Response.
//   	},
//   )
//
// 3. CreateSession() to start creating a Session.
//
//   // CreateSession() sends Create Session Request with given IEs, and stores information
//   // inside Session returned.
//   session, err := c.CreateSession(
//   	// put IEs required for your implementation here.
//   	// it is easier to use constructors in ies package.
//   	ies.NewIMSI("123451234567890"),
//   	// or, you can use ies.New() to create an IE without type-specific constructor.
//   	// put the type of IE, flags/instance, and payload as the parameters.
//   	ies.New(ies.ExtendedTraceInformation, 0x00, []byte{0xde, 0xad, 0xbe, 0xef}),
//   	// to set the instance to IE created with message-specific constructor, WithInstance()
//   	// may be your help.
//   	ies.NewIMSI("123451234567890").WithInstance(1), // no one wants to set instance to IMSI, though.
//   	// to be secure, TEID should be generated with random values, without conflicts in a Conn.
//   	// to achieve that, v2 provides NewFTEID() which returns F-TEID in *ies.IE.
//   	s11Conn.NewFTEID(v2.IFTypeS1UeNodeBGTPU, enbIP, ""),
//   )
//   if err != nil {
//   	// ...
//   }
//   // do not forget to add session to *Conn.
//   // do not Activate() it before you confirm the remote endpoint accepted the request.
//   c.AddSession(session)
//
// To wait for a Session to be created as a server, use ListenAndServe(), AddHandler(), and you can get *Conn, *Session, and *Bearer.
//
// 1. ListenAndServe() to retrieve *v2.Conn and start listening.
//
//   // give local net.Addr, restart counter, channel to let background process pass the errors.
//   conn, err := v2.ListenAndServe(laddr, 0, errCh)
//   if err != nil {
//   	// ...
//   }
//
// 2. AddHandler() to register your own handler in the same way as previous section.
//
// When adding handler for server, you should take the followings into account;
//
// * Session should be created by your own with NewSession(), and the subscriber/bearer information should be set properly(which is often in the request message).
//
// * Response with error should be sent before returning with failure.
package v2
