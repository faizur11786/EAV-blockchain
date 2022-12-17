// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: belshare/eav/merchant_new.proto

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

type MerchantNew struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Guid    string `protobuf:"bytes,2,opt,name=guid,proto3" json:"guid,omitempty"`
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MerchantNew) Reset()         { *m = MerchantNew{} }
func (m *MerchantNew) String() string { return proto.CompactTextString(m) }
func (*MerchantNew) ProtoMessage()    {}
func (*MerchantNew) Descriptor() ([]byte, []int) {
	return fileDescriptor_49aa576f77911fcd, []int{0}
}
func (m *MerchantNew) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MerchantNew) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MerchantNew.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MerchantNew) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantNew.Merge(m, src)
}
func (m *MerchantNew) XXX_Size() int {
	return m.Size()
}
func (m *MerchantNew) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantNew.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantNew proto.InternalMessageInfo

func (m *MerchantNew) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MerchantNew) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *MerchantNew) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*MerchantNew)(nil), "belshare.eav.MerchantNew")
}

func init() { proto.RegisterFile("belshare/eav/merchant_new.proto", fileDescriptor_49aa576f77911fcd) }

var fileDescriptor_49aa576f77911fcd = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x4a, 0xcd, 0x29,
	0xce, 0x48, 0x2c, 0x4a, 0xd5, 0x4f, 0x4d, 0x2c, 0xd3, 0xcf, 0x4d, 0x2d, 0x4a, 0xce, 0x48, 0xcc,
	0x2b, 0x89, 0xcf, 0x4b, 0x2d, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0x29, 0xd0,
	0x4b, 0x4d, 0x2c, 0x53, 0x0a, 0xe5, 0xe2, 0xf6, 0x85, 0xaa, 0xf1, 0x4b, 0x2d, 0x17, 0x92, 0xe0,
	0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
	0x71, 0x85, 0x84, 0xb8, 0x58, 0xd2, 0x4b, 0x33, 0x53, 0x24, 0x98, 0xc0, 0xc2, 0x60, 0x36, 0x48,
	0x75, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x7e, 0x91, 0x04, 0x33, 0x44, 0x35, 0x94, 0xeb, 0xa4, 0x77,
	0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7,
	0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x22, 0x49, 0xa9, 0x39, 0xc1, 0x60,
	0xf7, 0x55, 0x80, 0x5d, 0x58, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x76, 0x9b, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x0d, 0x32, 0x47, 0x7f, 0xbe, 0x00, 0x00, 0x00,
}

func (m *MerchantNew) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MerchantNew) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MerchantNew) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMerchantNew(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Guid) > 0 {
		i -= len(m.Guid)
		copy(dAtA[i:], m.Guid)
		i = encodeVarintMerchantNew(dAtA, i, uint64(len(m.Guid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMerchantNew(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMerchantNew(dAtA []byte, offset int, v uint64) int {
	offset -= sovMerchantNew(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MerchantNew) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMerchantNew(uint64(l))
	}
	l = len(m.Guid)
	if l > 0 {
		n += 1 + l + sovMerchantNew(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMerchantNew(uint64(l))
	}
	return n
}

func sovMerchantNew(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMerchantNew(x uint64) (n int) {
	return sovMerchantNew(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MerchantNew) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMerchantNew
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
			return fmt.Errorf("proto: MerchantNew: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MerchantNew: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerchantNew
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
				return ErrInvalidLengthMerchantNew
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMerchantNew
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Guid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerchantNew
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
				return ErrInvalidLengthMerchantNew
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMerchantNew
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Guid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerchantNew
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
				return ErrInvalidLengthMerchantNew
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMerchantNew
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMerchantNew(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMerchantNew
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
func skipMerchantNew(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMerchantNew
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
					return 0, ErrIntOverflowMerchantNew
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
					return 0, ErrIntOverflowMerchantNew
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
				return 0, ErrInvalidLengthMerchantNew
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMerchantNew
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMerchantNew
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMerchantNew        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMerchantNew          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMerchantNew = fmt.Errorf("proto: unexpected end of group")
)
