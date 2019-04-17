// Code generated by go generate; DO NOT EDIT.
// This file was generated by: go run gen/type.go

package object

import (
	"fmt"

	record "github.com/insolar/insolar/insolar/record"
)

func TypeFromRecord(generic record.VirtualRecord) TypeID {
	switch generic.(type) {
	case *GenesisRecord:
		return 100
	case *ChildRecord:
		return 101
	case *RequestRecord:
		return 200
	case *ResultRecord:
		return 300
	case *TypeRecord:
		return 301
	case *CodeRecord:
		return 302
	case *ActivateRecord:
		return 303
	case *AmendRecord:
		return 304
	case *DeactivationRecord:
		return 305
	default:
		panic(fmt.Sprintf("%T record is not registered", generic))
	}
}

func RecordFromType(i TypeID) record.VirtualRecord {
	switch i {
	case 100:
		return new(GenesisRecord).Init()
	case 101:
		return new(ChildRecord)
	case 200:
		return new(RequestRecord)
	case 300:
		return new(ResultRecord)
	case 301:
		return new(TypeRecord)
	case 302:
		return new(CodeRecord)
	case 303:
		return new(ActivateRecord)
	case 304:
		return new(AmendRecord)
	case 305:
		return new(DeactivationRecord)
	default:
		panic(fmt.Sprintf("identificator %d is not registered", i))
	}
}

func (i TypeID) String() string {
	switch i {
	case 100:
		return "GenesisRecord"
	case 101:
		return "ChildRecord"
	case 200:
		return "RequestRecord"
	case 300:
		return "ResultRecord"
	case 301:
		return "TypeRecord"
	case 302:
		return "CodeRecord"
	case 303:
		return "ActivateRecord"
	case 304:
		return "AmendRecord"
	case 305:
		return "DeactivationRecord"
	default:
		panic(fmt.Sprintf("identificator %d is not registered", i))
	}
}
