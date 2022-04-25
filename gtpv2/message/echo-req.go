// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/wmnsk/go-gtp/gtpv2/ie"
)

// EchoRequest is a EchoRequest Header and its IEs above.
type EchoRequest struct {
	*Header
	Recovery            *ie.IE
	SendingNodeFeatures *ie.IE
	PrivateExtension    *ie.IE
	AdditionalIEs       []*ie.IE
}

// NewEchoRequest creates a new EchoRequest.
func NewEchoRequest(seq uint32, ies ...*ie.IE) *EchoRequest {
	e := &EchoRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 0),
			MsgTypeEchoRequest, 0, seq, nil,
		),
	}

	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Recovery:
			e.Recovery = i
		case ie.NodeFeatures:
			e.SendingNodeFeatures = i
		case ie.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}

	e.SetLength()
	return e
}

// Marshal returns the byte sequence generated from a EchoRequest.
func (e *EchoRequest) Marshal() ([]byte, error) {
	b := make([]byte, e.MarshalLen())
	if err := e.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (e *EchoRequest) MarshalTo(b []byte) error {
	if e.Header.Payload != nil {
		e.Header.Payload = nil
	}
	e.Header.Payload = make([]byte, e.MarshalLen()-e.Header.MarshalLen())

	offset := 0
	if ie := e.Recovery; ie != nil {
		if err := ie.MarshalTo(e.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := e.SendingNodeFeatures; ie != nil {
		if err := ie.MarshalTo(e.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := e.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(e.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	for _, ie := range e.AdditionalIEs {
		if err := ie.MarshalTo(e.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	e.Header.SetLength()
	return e.Header.MarshalTo(b)
}

// ParseEchoRequest decodes a given byte sequence as a EchoRequest.
func ParseEchoRequest(b []byte) (*EchoRequest, error) {
	e := &EchoRequest{}
	if err := e.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return e, nil
}

// UnmarshalBinary decodes a given byte sequence as a EchoRequest.
func (e *EchoRequest) UnmarshalBinary(b []byte) error {
	var err error
	e.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(e.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ie.ParseMultiIEs(e.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ie.Recovery:
			e.Recovery = i
		case ie.NodeFeatures:
			e.SendingNodeFeatures = i
		case ie.PrivateExtension:
			e.PrivateExtension = i
		default:
			e.AdditionalIEs = append(e.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (e *EchoRequest) MarshalLen() int {
	l := e.Header.MarshalLen() - len(e.Header.Payload)

	if ie := e.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := e.SendingNodeFeatures; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := e.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range e.AdditionalIEs {
		l += ie.MarshalLen()
	}

	return l
}

// SetLength sets the length in Length field.
func (e *EchoRequest) SetLength() {
	e.Header.Length = uint16(e.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (e *EchoRequest) MessageTypeName() string {
	return "Echo Request"
}

// TEID returns the TEID in uint32.
func (e *EchoRequest) TEID() uint32 {
	return e.Header.teid()
}

//////////////////////////////////////////////////////////////////////////////

func (e *EchoRequest) MarshalTo1(Payload []byte, marshalLen int, debug bool) error {

	e.Header.Payload = make([]byte, marshalLen-e.Header.MarshalLen())

	itemVal := reflect.ValueOf(*e)
	var IElen, offset int64
	for i := 1; i < itemVal.NumField(); i++ { //loop over fields in msg
		fieldVal := itemVal.Field(i)        //get a field
		if fieldVal.Kind() == reflect.Ptr { //check that it's a pointer
			fieldVal1 := fieldVal.Elem() // This is only helpful if the field is not nil
			if debug {
				fmt.Println(i, " ==> fieldVal ", fieldVal.Kind(), "   ", fieldVal1.Kind())
			}
			if fieldVal1.Kind() != reflect.Invalid { //check that Kind is valid
				//fmt.Println(i, " ==> fieldVal1.Type().Name() ;", fieldVal1.Type().Name(), ";")

				MarshalLenMethod := fieldVal.MethodByName("MarshalLen")      //get method (i *IE) MarshalLen()
				MarshalLenResult := MarshalLenMethod.Call([]reflect.Value{}) //call method (i *IE) MarshalLen()
				IElen = MarshalLenResult[0].Int()

				MarshalToMethod := fieldVal.MethodByName("MarshalTo")                                                //get method (i *IE) MarshalTo()
				MarshalToResult := MarshalToMethod.Call([]reflect.Value{reflect.ValueOf(e.Header.Payload[offset:])}) //call method (i *IE) MarshalTo()
				if debug {
					fmt.Println(i, " MarshalToResult[0] ==> ", MarshalToResult[0], "IElen = ", IElen)
					fmt.Println(i, " ==> fieldVal1 ", fieldVal1.Kind(),
						"  reflect.TypeOf(fieldVal1) =  ", fieldVal1.Type().Name())
				}
				offset += IElen //update offset

			}
			//fmt.Println(i, " ==> fieldVal ", fieldVal.Kind(), "   ", fieldVal1.Kind())
		} else {
			if debug {
				fmt.Println(i, " ==> fieldVal ", fieldVal.Kind())
			}
		}
	}
	if debug {
		fmt.Println("MarshalTo(msg *EchoRequest) ==> msg.Header.Payload =  ", hex.Dump(e.Header.Payload))
	}
	e.Header.SetLength()
	return e.Header.MarshalTo(Payload)

}

func (e *EchoRequest) MarshalLen1(debug bool) int64 {
	if debug {
		fmt.Printf(" EchoRequest ==> len(msg.Header.Payload) =  %d \n", len(e.Header.Payload))
	}
	itemVal := reflect.ValueOf(*e)
	var Totlen, IElen int64
	for i := 0; i < itemVal.NumField(); i++ { //loop over fields in msg
		fieldVal := itemVal.Field(i)        //get a field
		if fieldVal.Kind() == reflect.Ptr { //check that it's a pointer
			fieldVal1 := fieldVal.Elem() // This is only helpful if the field is not nil
			if debug {
				fmt.Println(i, " ==> fieldVal ", fieldVal.Kind(), "   ", fieldVal1.Kind())
			}
			if fieldVal1.Kind() != reflect.Invalid { //check that Kind is valid
				//fmt.Println(i, " ==> fieldVal1.Type().Name() ;", fieldVal1.Type().Name(), ";")

				MarshalLenMethod := fieldVal.MethodByName("MarshalLen")      //get method (i *IE) MarshalLen()
				MarshalLenResult := MarshalLenMethod.Call([]reflect.Value{}) //call method (i *IE) MarshalLen()
				IElen = MarshalLenResult[0].Int()                            // get result from (i *IE) MarshalLen()
				Totlen += IElen                                              //add MarshalLen to total length
				if debug {
					fmt.Println(i, "IElen = ", IElen)
				}
			}
			//fmt.Println(i, " ==> fieldVal ", fieldVal.Kind(), "   ", fieldVal1.Kind())
		} else {
			if debug {
				fmt.Println(i, " ==> fieldVal ", fieldVal.Kind())
			}
		}
	}
	return Totlen
}
