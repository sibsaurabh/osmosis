// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/incentives/autostaking.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AutoStaking struct {
	AutostakingValidator []string                               `protobuf:"bytes,1,rep,name=autostaking_validator,json=autostakingValidator,proto3" json:"autostaking_validator,omitempty"`
	AutostakingRate      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=autostaking_rate,json=autostakingRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"autostaking_rate" yaml:"autostake_rate"`
}

func (m *AutoStaking) Reset()      { *m = AutoStaking{} }
func (*AutoStaking) ProtoMessage() {}
func (*AutoStaking) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a53c96e1fd87658, []int{0}
}
func (m *AutoStaking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AutoStaking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AutoStaking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AutoStaking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoStaking.Merge(m, src)
}
func (m *AutoStaking) XXX_Size() int {
	return m.Size()
}
func (m *AutoStaking) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoStaking.DiscardUnknown(m)
}

var xxx_messageInfo_AutoStaking proto.InternalMessageInfo

func (m *AutoStaking) GetAutostakingValidator() []string {
	if m != nil {
		return m.AutostakingValidator
	}
	return nil
}

func init() {
	proto.RegisterType((*AutoStaking)(nil), "osmosis.incentives.AutoStaking")
}

func init() {
	proto.RegisterFile("osmosis/incentives/autostaking.proto", fileDescriptor_4a53c96e1fd87658)
}

var fileDescriptor_4a53c96e1fd87658 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0xcf, 0xcc, 0x4b, 0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0x2d, 0xd6, 0x4f,
	0x2c, 0x2d, 0xc9, 0x2f, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x82, 0xaa, 0xd2, 0x43, 0xa8, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xeb,
	0x83, 0x58, 0x10, 0x95, 0x4a, 0x07, 0x18, 0xb9, 0xb8, 0x1d, 0x4b, 0x4b, 0xf2, 0x83, 0x21, 0xfa,
	0x85, 0x8c, 0xb9, 0x44, 0x91, 0x8c, 0x8b, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f,
	0x92, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x12, 0x41, 0x92, 0x0c, 0x83, 0xc9, 0x09, 0x15, 0x71,
	0x09, 0x20, 0x6b, 0x2a, 0x4a, 0x2c, 0x49, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x74, 0x72, 0x3f,
	0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79, 0xb5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4,
	0xfc, 0x5c, 0xfd, 0x64, 0xb0, 0xe3, 0xa0, 0x94, 0x6e, 0x71, 0x4a, 0xb6, 0x7e, 0x49, 0x65, 0x41,
	0x6a, 0xb1, 0x9e, 0x4b, 0x6a, 0xf2, 0xa7, 0x7b, 0xf2, 0xa2, 0x95, 0x89, 0xb9, 0x39, 0x56, 0x4a,
	0x30, 0xf3, 0x52, 0xc1, 0xa6, 0x29, 0x05, 0xf1, 0x23, 0x59, 0x10, 0x94, 0x58, 0x92, 0x6a, 0xc5,
	0x31, 0x63, 0x81, 0x3c, 0x43, 0xc3, 0x1d, 0x05, 0x46, 0x27, 0x9f, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x42, 0xb2, 0x15, 0x1a, 0x22, 0xba, 0x39, 0x89, 0x49, 0xc5,
	0x30, 0x8e, 0x7e, 0x05, 0x72, 0x30, 0x82, 0x5d, 0x91, 0xc4, 0x06, 0x0e, 0x17, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xdf, 0x7d, 0xb9, 0x6b, 0x69, 0x01, 0x00, 0x00,
}

func (m *AutoStaking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AutoStaking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AutoStaking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AutostakingRate.Size()
		i -= size
		if _, err := m.AutostakingRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAutostaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.AutostakingValidator) > 0 {
		for iNdEx := len(m.AutostakingValidator) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AutostakingValidator[iNdEx])
			copy(dAtA[i:], m.AutostakingValidator[iNdEx])
			i = encodeVarintAutostaking(dAtA, i, uint64(len(m.AutostakingValidator[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAutostaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovAutostaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AutoStaking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AutostakingValidator) > 0 {
		for _, s := range m.AutostakingValidator {
			l = len(s)
			n += 1 + l + sovAutostaking(uint64(l))
		}
	}
	l = m.AutostakingRate.Size()
	n += 1 + l + sovAutostaking(uint64(l))
	return n
}

func sovAutostaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAutostaking(x uint64) (n int) {
	return sovAutostaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AutoStaking) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AutoStaking{`,
		`AutostakingValidator:` + fmt.Sprintf("%v", this.AutostakingValidator) + `,`,
		`AutostakingRate:` + fmt.Sprintf("%v", this.AutostakingRate) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAutostaking(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AutoStaking) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAutostaking
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
			return fmt.Errorf("proto: AutoStaking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AutoStaking: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutostakingValidator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAutostaking
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
				return ErrInvalidLengthAutostaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAutostaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AutostakingValidator = append(m.AutostakingValidator, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutostakingRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAutostaking
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
				return ErrInvalidLengthAutostaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAutostaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AutostakingRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAutostaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAutostaking
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAutostaking
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
func skipAutostaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAutostaking
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
					return 0, ErrIntOverflowAutostaking
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
					return 0, ErrIntOverflowAutostaking
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
				return 0, ErrInvalidLengthAutostaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAutostaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAutostaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAutostaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAutostaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAutostaking = fmt.Errorf("proto: unexpected end of group")
)
