// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: belshare/eav/entity_type.proto

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

type EntityType struct {
	Guid    string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *EntityType) Reset()         { *m = EntityType{} }
func (m *EntityType) String() string { return proto.CompactTextString(m) }
func (*EntityType) ProtoMessage()    {}
func (*EntityType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0d5db68a0557c75, []int{0}
}
func (m *EntityType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityType.Merge(m, src)
}
func (m *EntityType) XXX_Size() int {
	return m.Size()
}
func (m *EntityType) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityType.DiscardUnknown(m)
}

var xxx_messageInfo_EntityType proto.InternalMessageInfo

func (m *EntityType) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *EntityType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EntityType) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*EntityType)(nil), "belshare.eav.EntityType")
}

func init() { proto.RegisterFile("belshare/eav/entity_type.proto", fileDescriptor_f0d5db68a0557c75) }

var fileDescriptor_f0d5db68a0557c75 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x4a, 0xcd, 0x29,
	0xce, 0x48, 0x2c, 0x4a, 0xd5, 0x4f, 0x4d, 0x2c, 0xd3, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x8c,
	0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0xc9, 0xeb, 0xa5,
	0x26, 0x96, 0x29, 0xf9, 0x71, 0x71, 0xb9, 0x82, 0x95, 0x84, 0x54, 0x16, 0xa4, 0x0a, 0x09, 0x71,
	0xb1, 0xa4, 0x97, 0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20, 0xb1,
	0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x88, 0x18, 0x88, 0x2d, 0x24, 0xc1, 0xc5, 0x9e, 0x5c, 0x94,
	0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1, 0x0c, 0x16, 0x86, 0x71, 0x9d, 0xf4, 0x4e, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x24, 0x29, 0x35, 0x27, 0x18, 0xec, 0xae, 0x0a, 0xb0,
	0xcb, 0x40, 0x4e, 0x2a, 0x4e, 0x62, 0x03, 0x3b, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x5b,
	0xfa, 0x01, 0xbc, 0xb6, 0x00, 0x00, 0x00,
}

func (m *EntityType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEntityType(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEntityType(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Guid) > 0 {
		i -= len(m.Guid)
		copy(dAtA[i:], m.Guid)
		i = encodeVarintEntityType(dAtA, i, uint64(len(m.Guid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEntityType(dAtA []byte, offset int, v uint64) int {
	offset -= sovEntityType(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EntityType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Guid)
	if l > 0 {
		n += 1 + l + sovEntityType(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEntityType(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEntityType(uint64(l))
	}
	return n
}

func sovEntityType(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEntityType(x uint64) (n int) {
	return sovEntityType(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EntityType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntityType
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
			return fmt.Errorf("proto: EntityType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Guid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityType
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
				return ErrInvalidLengthEntityType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntityType
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
					return ErrIntOverflowEntityType
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
				return ErrInvalidLengthEntityType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntityType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityType
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
				return ErrInvalidLengthEntityType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntityType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntityType(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEntityType
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
func skipEntityType(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntityType
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
					return 0, ErrIntOverflowEntityType
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
					return 0, ErrIntOverflowEntityType
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
				return 0, ErrInvalidLengthEntityType
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEntityType
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEntityType
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEntityType        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntityType          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEntityType = fmt.Errorf("proto: unexpected end of group")
)