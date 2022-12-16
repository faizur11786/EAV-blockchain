// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: belshare/eav/attribute.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Attribute struct {
	Guid     string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EntityId string `protobuf:"bytes,3,opt,name=entityId,proto3" json:"entityId,omitempty"`
}

func (m *Attribute) Reset()         { *m = Attribute{} }
func (m *Attribute) String() string { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()    {}
func (*Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c14089b01901b88, []int{0}
}
func (m *Attribute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Attribute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribute.Merge(m, src)
}
func (m *Attribute) XXX_Size() int {
	return m.Size()
}
func (m *Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Attribute proto.InternalMessageInfo

func (m *Attribute) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *Attribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Attribute) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func init() {
	proto.RegisterType((*Attribute)(nil), "belshare.eav.Attribute")
}

func init() { proto.RegisterFile("belshare/eav/attribute.proto", fileDescriptor_3c14089b01901b88) }

var fileDescriptor_3c14089b01901b88 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x4a, 0xcd, 0x29,
	0xce, 0x48, 0x2c, 0x4a, 0xd5, 0x4f, 0x4d, 0x2c, 0xd3, 0x4f, 0x2c, 0x29, 0x29, 0xca, 0x4c, 0x2a,
	0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0xc9, 0xea, 0xa5, 0x26, 0x96,
	0x29, 0xf9, 0x73, 0x71, 0x3a, 0xc2, 0x14, 0x08, 0x09, 0x71, 0xb1, 0xa4, 0x97, 0x66, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20, 0xb1, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26,
	0x88, 0x18, 0x88, 0x2d, 0x24, 0xc5, 0xc5, 0x91, 0x9a, 0x57, 0x92, 0x59, 0x52, 0xe9, 0x99, 0x22,
	0xc1, 0x0c, 0x16, 0x87, 0xf3, 0x9d, 0xf4, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e,
	0x21, 0x4a, 0x24, 0x29, 0x35, 0x27, 0x18, 0xec, 0xac, 0x0a, 0xb0, 0xc3, 0x4a, 0x2a, 0x0b, 0x52,
	0x8b, 0x93, 0xd8, 0xc0, 0xae, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x06, 0x77, 0xbb, 0x0e,
	0xb5, 0x00, 0x00, 0x00,
}

func (m *Attribute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attribute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Attribute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EntityId) > 0 {
		i -= len(m.EntityId)
		copy(dAtA[i:], m.EntityId)
		i = encodeVarintAttribute(dAtA, i, uint64(len(m.EntityId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAttribute(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Guid) > 0 {
		i -= len(m.Guid)
		copy(dAtA[i:], m.Guid)
		i = encodeVarintAttribute(dAtA, i, uint64(len(m.Guid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAttribute(dAtA []byte, offset int, v uint64) int {
	offset -= sovAttribute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Attribute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Guid)
	if l > 0 {
		n += 1 + l + sovAttribute(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAttribute(uint64(l))
	}
	l = len(m.EntityId)
	if l > 0 {
		n += 1 + l + sovAttribute(uint64(l))
	}
	return n
}

func sovAttribute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttribute(x uint64) (n int) {
	return sovAttribute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Attribute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttribute
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
			return fmt.Errorf("proto: Attribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Guid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttribute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAttribute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Guid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttribute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAttribute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttribute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAttribute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttribute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttribute
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
func skipAttribute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttribute
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
					return 0, ErrIntOverflowAttribute
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAttribute
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
				return 0, ErrInvalidLengthAttribute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAttribute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAttribute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAttribute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttribute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAttribute = fmt.Errorf("proto: unexpected end of group")
)
