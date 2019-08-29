// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ledger/drop/drop.proto

package drop

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	github_com_insolar_insolar_insolar "github.com/insolar/insolar/insolar"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Drop struct {
	Polymorph              int32                                          `protobuf:"varint,16,opt,name=polymorph,proto3" json:"polymorph,omitempty"`
	Pulse                  github_com_insolar_insolar_insolar.PulseNumber `protobuf:"bytes,20,opt,name=Pulse,proto3,customtype=github.com/insolar/insolar/insolar.PulseNumber" json:"Pulse"`
	PrevHash               []byte                                         `protobuf:"bytes,21,opt,name=PrevHash,proto3" json:"PrevHash,omitempty"`
	Hash                   []byte                                         `protobuf:"bytes,22,opt,name=Hash,proto3" json:"Hash,omitempty"`
	DropSize               uint64                                         `protobuf:"varint,23,opt,name=DropSize,proto3" json:"DropSize,omitempty"`
	JetID                  github_com_insolar_insolar_insolar.JetID       `protobuf:"bytes,24,opt,name=JetID,proto3,customtype=github.com/insolar/insolar/insolar.JetID" json:"JetID"`
	SplitThresholdExceeded int64                                          `protobuf:"varint,25,opt,name=SplitThresholdExceeded,proto3" json:"SplitThresholdExceeded,omitempty"`
	Split                  bool                                           `protobuf:"varint,26,opt,name=Split,proto3" json:"Split,omitempty"`
}

func (m *Drop) Reset()      { *m = Drop{} }
func (*Drop) ProtoMessage() {}
func (*Drop) Descriptor() ([]byte, []int) {
	return fileDescriptor_f87624f7639ca597, []int{0}
}
func (m *Drop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Drop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Drop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Drop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Drop.Merge(m, src)
}
func (m *Drop) XXX_Size() int {
	return m.Size()
}
func (m *Drop) XXX_DiscardUnknown() {
	xxx_messageInfo_Drop.DiscardUnknown(m)
}

var xxx_messageInfo_Drop proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Drop)(nil), "drop.Drop")
}

func init() { proto.RegisterFile("ledger/drop/drop.proto", fileDescriptor_f87624f7639ca597) }

var fileDescriptor_f87624f7639ca597 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x4f, 0xb4, 0x95, 0x19, 0x3c, 0x48, 0x98, 0x33, 0x0e, 0xf9, 0x56, 0x3c, 0xf5, 0x62, 0x27,
	0x08, 0x43, 0x3c, 0x8e, 0x29, 0x2a, 0x22, 0xa3, 0xf3, 0x05, 0xd6, 0x35, 0xb6, 0x85, 0xce, 0x94,
	0xb4, 0x15, 0xf5, 0xe4, 0x23, 0xf8, 0x08, 0x1e, 0xf7, 0x28, 0x3b, 0xee, 0x38, 0x76, 0x18, 0xb6,
	0xbb, 0x78, 0xdc, 0x23, 0x48, 0x53, 0x99, 0x22, 0x08, 0x5e, 0x92, 0xdf, 0x9f, 0xfc, 0xbe, 0xfc,
	0xe0, 0x23, 0xb5, 0x90, 0xbb, 0x1e, 0x97, 0x4d, 0x57, 0x8a, 0x48, 0x1d, 0x56, 0x24, 0x45, 0x22,
	0xa8, 0x56, 0xe0, 0xfa, 0xa1, 0x17, 0x24, 0x7e, 0xea, 0x58, 0x03, 0x31, 0x6c, 0x7a, 0xc2, 0x13,
	0x4d, 0x65, 0x3a, 0xe9, 0x9d, 0x62, 0x8a, 0x28, 0x54, 0x86, 0x0e, 0x66, 0x6b, 0x44, 0xeb, 0x48,
	0x11, 0xd1, 0x7d, 0xb2, 0x19, 0x89, 0xf0, 0x69, 0x28, 0x64, 0xe4, 0xb3, 0x6d, 0x03, 0x9b, 0xba,
	0xfd, 0x2d, 0xd0, 0x6b, 0xa2, 0x77, 0xd3, 0x30, 0xe6, 0xac, 0x6a, 0x60, 0x73, 0xab, 0xdd, 0x1a,
	0xcf, 0x1b, 0x68, 0x36, 0x6f, 0x58, 0x3f, 0x3e, 0x0b, 0xee, 0x63, 0x11, 0xf6, 0xe5, 0xef, 0xdb,
	0x52, 0xb9, 0x9b, 0x74, 0xe8, 0x70, 0x69, 0x97, 0x43, 0x68, 0x9d, 0x54, 0xba, 0x92, 0x3f, 0x5c,
	0xf4, 0x63, 0x9f, 0xed, 0x14, 0x03, 0xed, 0x15, 0xa7, 0x94, 0x68, 0x4a, 0xaf, 0x29, 0x5d, 0xe1,
	0xe2, 0x7d, 0xd1, 0xb1, 0x17, 0x3c, 0x73, 0xb6, 0x6b, 0x60, 0x53, 0xb3, 0x57, 0x9c, 0x9e, 0x13,
	0xfd, 0x8a, 0x27, 0x97, 0x1d, 0xc6, 0x54, 0xb3, 0xa3, 0xaf, 0x66, 0xe6, 0x3f, 0x9a, 0xa9, 0x9c,
	0x5d, 0xc6, 0x69, 0x8b, 0xd4, 0x7a, 0x51, 0x18, 0x24, 0xb7, 0xbe, 0xe4, 0xb1, 0x2f, 0x42, 0xf7,
	0xec, 0x71, 0xc0, 0xb9, 0xcb, 0x5d, 0xb6, 0x67, 0x60, 0x73, 0xdd, 0xfe, 0xc3, 0xa5, 0x55, 0xa2,
	0x2b, 0x87, 0xd5, 0x0d, 0x6c, 0x56, 0xec, 0x92, 0x9c, 0x6a, 0xa3, 0xb7, 0x06, 0x6e, 0x9f, 0x8c,
	0x33, 0x40, 0x93, 0x0c, 0xd0, 0x34, 0x03, 0xb4, 0xcc, 0x00, 0xbf, 0xe4, 0x80, 0x47, 0x39, 0xe0,
	0x71, 0x0e, 0x78, 0x92, 0x03, 0x7e, 0xcf, 0x01, 0x7f, 0xe4, 0x80, 0x96, 0x39, 0xe0, 0xd7, 0x05,
	0xa0, 0xc9, 0x02, 0xd0, 0x74, 0x01, 0xc8, 0xd9, 0x50, 0xdb, 0x39, 0xfe, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0x15, 0x2a, 0x64, 0x83, 0xec, 0x01, 0x00, 0x00,
}

func (this *Drop) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Drop)
	if !ok {
		that2, ok := that.(Drop)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Polymorph != that1.Polymorph {
		return false
	}
	if !this.Pulse.Equal(that1.Pulse) {
		return false
	}
	if !bytes.Equal(this.PrevHash, that1.PrevHash) {
		return false
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	if this.DropSize != that1.DropSize {
		return false
	}
	if !this.JetID.Equal(that1.JetID) {
		return false
	}
	if this.SplitThresholdExceeded != that1.SplitThresholdExceeded {
		return false
	}
	if this.Split != that1.Split {
		return false
	}
	return true
}

type DropFace interface {
	Proto() github_com_gogo_protobuf_proto.Message
	GetPolymorph() int32
	GetPulse() github_com_insolar_insolar_insolar.PulseNumber
	GetPrevHash() []byte
	GetHash() []byte
	GetDropSize() uint64
	GetJetID() github_com_insolar_insolar_insolar.JetID
	GetSplitThresholdExceeded() int64
	GetSplit() bool
}

func (this *Drop) Proto() github_com_gogo_protobuf_proto.Message {
	return this
}

func (this *Drop) TestProto() github_com_gogo_protobuf_proto.Message {
	return NewDropFromFace(this)
}

func (this *Drop) GetPolymorph() int32 {
	return this.Polymorph
}

func (this *Drop) GetPulse() github_com_insolar_insolar_insolar.PulseNumber {
	return this.Pulse
}

func (this *Drop) GetPrevHash() []byte {
	return this.PrevHash
}

func (this *Drop) GetHash() []byte {
	return this.Hash
}

func (this *Drop) GetDropSize() uint64 {
	return this.DropSize
}

func (this *Drop) GetJetID() github_com_insolar_insolar_insolar.JetID {
	return this.JetID
}

func (this *Drop) GetSplitThresholdExceeded() int64 {
	return this.SplitThresholdExceeded
}

func (this *Drop) GetSplit() bool {
	return this.Split
}

func NewDropFromFace(that DropFace) *Drop {
	this := &Drop{}
	this.Polymorph = that.GetPolymorph()
	this.Pulse = that.GetPulse()
	this.PrevHash = that.GetPrevHash()
	this.Hash = that.GetHash()
	this.DropSize = that.GetDropSize()
	this.JetID = that.GetJetID()
	this.SplitThresholdExceeded = that.GetSplitThresholdExceeded()
	this.Split = that.GetSplit()
	return this
}

func (this *Drop) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&drop.Drop{")
	s = append(s, "Polymorph: "+fmt.Sprintf("%#v", this.Polymorph)+",\n")
	s = append(s, "Pulse: "+fmt.Sprintf("%#v", this.Pulse)+",\n")
	s = append(s, "PrevHash: "+fmt.Sprintf("%#v", this.PrevHash)+",\n")
	s = append(s, "Hash: "+fmt.Sprintf("%#v", this.Hash)+",\n")
	s = append(s, "DropSize: "+fmt.Sprintf("%#v", this.DropSize)+",\n")
	s = append(s, "JetID: "+fmt.Sprintf("%#v", this.JetID)+",\n")
	s = append(s, "SplitThresholdExceeded: "+fmt.Sprintf("%#v", this.SplitThresholdExceeded)+",\n")
	s = append(s, "Split: "+fmt.Sprintf("%#v", this.Split)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDrop(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Drop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Drop) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Polymorph != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintDrop(dAtA, i, uint64(m.Polymorph))
	}
	dAtA[i] = 0xa2
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintDrop(dAtA, i, uint64(m.Pulse.Size()))
	n1, err := m.Pulse.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.PrevHash) > 0 {
		dAtA[i] = 0xaa
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintDrop(dAtA, i, uint64(len(m.PrevHash)))
		i += copy(dAtA[i:], m.PrevHash)
	}
	if len(m.Hash) > 0 {
		dAtA[i] = 0xb2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintDrop(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if m.DropSize != 0 {
		dAtA[i] = 0xb8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintDrop(dAtA, i, uint64(m.DropSize))
	}
	dAtA[i] = 0xc2
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintDrop(dAtA, i, uint64(m.JetID.Size()))
	n2, err := m.JetID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.SplitThresholdExceeded != 0 {
		dAtA[i] = 0xc8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintDrop(dAtA, i, uint64(m.SplitThresholdExceeded))
	}
	if m.Split {
		dAtA[i] = 0xd0
		i++
		dAtA[i] = 0x1
		i++
		if m.Split {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintDrop(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Drop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Polymorph != 0 {
		n += 2 + sovDrop(uint64(m.Polymorph))
	}
	l = m.Pulse.Size()
	n += 2 + l + sovDrop(uint64(l))
	l = len(m.PrevHash)
	if l > 0 {
		n += 2 + l + sovDrop(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 2 + l + sovDrop(uint64(l))
	}
	if m.DropSize != 0 {
		n += 2 + sovDrop(uint64(m.DropSize))
	}
	l = m.JetID.Size()
	n += 2 + l + sovDrop(uint64(l))
	if m.SplitThresholdExceeded != 0 {
		n += 2 + sovDrop(uint64(m.SplitThresholdExceeded))
	}
	if m.Split {
		n += 3
	}
	return n
}

func sovDrop(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDrop(x uint64) (n int) {
	return sovDrop(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Drop) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Drop{`,
		`Polymorph:` + fmt.Sprintf("%v", this.Polymorph) + `,`,
		`Pulse:` + fmt.Sprintf("%v", this.Pulse) + `,`,
		`PrevHash:` + fmt.Sprintf("%v", this.PrevHash) + `,`,
		`Hash:` + fmt.Sprintf("%v", this.Hash) + `,`,
		`DropSize:` + fmt.Sprintf("%v", this.DropSize) + `,`,
		`JetID:` + fmt.Sprintf("%v", this.JetID) + `,`,
		`SplitThresholdExceeded:` + fmt.Sprintf("%v", this.SplitThresholdExceeded) + `,`,
		`Split:` + fmt.Sprintf("%v", this.Split) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDrop(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Drop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDrop
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Drop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Drop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Polymorph", wireType)
			}
			m.Polymorph = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Polymorph |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pulse", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pulse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevHash = append(m.PrevHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PrevHash == nil {
				m.PrevHash = []byte{}
			}
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 23:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DropSize", wireType)
			}
			m.DropSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DropSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 24:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JetID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.JetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 25:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SplitThresholdExceeded", wireType)
			}
			m.SplitThresholdExceeded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SplitThresholdExceeded |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 26:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Split", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Split = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDrop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDrop
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDrop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDrop(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDrop
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDrop
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDrop
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDrop
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDrop(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDrop
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDrop = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDrop   = fmt.Errorf("proto: integer overflow")
)