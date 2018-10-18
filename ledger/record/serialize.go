/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package record

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ugorji/go/codec"

	"github.com/insolar/insolar/core"
)

// hashableBytes exists just to allow []byte implements hash.Writer
type hashableBytes []byte

func (b hashableBytes) WriteHash(w io.Writer) {
	_, err := w.Write(b)
	if err != nil {
		panic(err)
	}
}

// Bytes2ID converts ID from byte representation to struct.
func Bytes2ID(b []byte) ID {
	return ID{
		Pulse: core.Bytes2PulseNumber(b[:core.PulseNumberSize]),
		Hash:  b[core.PulseNumberSize:],
	}
}

// Core2Reference converts commonly used reference to Ledger-specific.
func Core2Reference(cRef core.RecordRef) Reference {
	return Reference{
		Record: Bytes2ID(cRef[:core.RecordIDSize]),
		Domain: Bytes2ID(cRef[core.RecordIDSize:]),
	}
}

// ID2Bytes converts ID struct to it's byte representation.
func ID2Bytes(id ID) []byte {
	rec := core.GenRecordID(id.Pulse, id.Hash)
	return rec[:]
}

// record type ids for record types
// in use mostly for hashing and deserialization
// (we don't use iota for clarity and predictable ids,
// not depended on definition order)
const (
	// meta
	typeGenesis TypeID = 10
	typeChild   TypeID = 11

	// request
	typeCallRequest TypeID = 20

	// result
	typeType           TypeID = 30
	typeCode           TypeID = 31
	typeClassActivate  TypeID = 32
	typeClassAmend     TypeID = 33
	typeObjectActivate TypeID = 34
	typeObjectAmend    TypeID = 35
	typeDeactivate     TypeID = 36
)

// getRecordByTypeID returns Record interface with concrete record type under the hood.
// This is useful with deserialization cases.
func getRecordByTypeID(id TypeID) Record { // nolint: gocyclo
	switch id {
	// request records
	case typeCallRequest:
		return &CallRequest{}
	case typeClassActivate:
		return &ClassActivateRecord{}
	case typeObjectActivate:
		return &ObjectActivateRecord{}
	case typeCode:
		return &CodeRecord{}
	case typeClassAmend:
		return &ClassAmendRecord{}
	case typeDeactivate:
		return &DeactivationRecord{}
	case typeObjectAmend:
		return &ObjectAmendRecord{}
	case typeType:
		return &TypeRecord{}
	case typeChild:
		return &ChildRecord{}
	case typeGenesis:
		return &GenesisRecord{}
	default:
		panic(fmt.Errorf("unknown record type id %v", id))
	}
}

// SerializeType returns binary representation of provided type.
func SerializeType(id TypeID) []byte {
	buf := make([]byte, 4) // uint32
	binary.BigEndian.PutUint32(buf, uint32(id))
	return buf
}

// DeserializeType returns type from provided binary representation.
func DeserializeType(buf []byte) TypeID {
	return TypeID(binary.BigEndian.Uint32(buf))
}

// SerializeRecord returns binary representation of provided record.
func SerializeRecord(rec Record) []byte {
	typeBytes := SerializeType(rec.Type())
	buff := bytes.NewBuffer(typeBytes)
	enc := codec.NewEncoder(buff, &codec.CborHandle{})
	enc.MustEncode(rec)
	return buff.Bytes()
}

// DeserializeRecord returns record decoded from bytes.
func DeserializeRecord(buf []byte) Record {
	t := DeserializeType(buf[:4]) // uint32
	dec := codec.NewDecoderBytes(buf[4:], &codec.CborHandle{})
	rec := getRecordByTypeID(t)
	dec.MustDecode(&rec)
	return rec
}
