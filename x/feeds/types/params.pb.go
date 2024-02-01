// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feeds/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params is the data structure that keeps the parameters of the feeds module.
type Params struct {
	// The address of the admin that is allowed to perform operations on modules.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// allow_diff_time is the allowed difference (in seconds) between timestamp and block_time when validator submits the
	// prices.
	AllowDiffTime int64 `protobuf:"varint,2,opt,name=allow_diff_time,json=allowDiffTime,proto3" json:"allow_diff_time,omitempty"`
	// transition_time is the time (in seconds) given for validators to adapt to changing in symbol's interval.
	TransitionTime int64 `protobuf:"varint,3,opt,name=transition_time,json=transitionTime,proto3" json:"transition_time,omitempty"`
	// min_interval is the minimum limit of every symbols' interval (in seconds).
	// If the calculated interval is lower than this, it will be capped at this value.
	MinInterval int64 `protobuf:"varint,4,opt,name=min_interval,json=minInterval,proto3" json:"min_interval,omitempty"`
	// max_interval is the maximum limit of every symbols' interval (in seconds).
	// If the calculated interval of a symbol is higher than this, it will not be recognized as a supported symbol.
	MaxInterval int64 `protobuf:"varint,5,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
	// time_dividend is the dividend for interval calculation.
	TimeDividend int64 `protobuf:"varint,6,opt,name=time_dividend,json=timeDividend,proto3" json:"time_dividend,omitempty"`
	// max_support_symbol is the maximum number of symbols supported at a time.
	MaxSupportedSymbol uint64 `protobuf:"varint,7,opt,name=max_supported_symbol,json=maxSupportedSymbol,proto3" json:"max_supported_symbol,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbfae8ad171874f3, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *Params) GetAllowDiffTime() int64 {
	if m != nil {
		return m.AllowDiffTime
	}
	return 0
}

func (m *Params) GetTransitionTime() int64 {
	if m != nil {
		return m.TransitionTime
	}
	return 0
}

func (m *Params) GetMinInterval() int64 {
	if m != nil {
		return m.MinInterval
	}
	return 0
}

func (m *Params) GetMaxInterval() int64 {
	if m != nil {
		return m.MaxInterval
	}
	return 0
}

func (m *Params) GetTimeDividend() int64 {
	if m != nil {
		return m.TimeDividend
	}
	return 0
}

func (m *Params) GetMaxSupportedSymbol() uint64 {
	if m != nil {
		return m.MaxSupportedSymbol
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "feeds.v1beta1.Params")
}

func init() { proto.RegisterFile("feeds/v1beta1/params.proto", fileDescriptor_bbfae8ad171874f3) }

var fileDescriptor_bbfae8ad171874f3 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd1, 0xbf, 0x6e, 0xe2, 0x30,
	0x00, 0xc7, 0xf1, 0x98, 0x7f, 0x77, 0xe7, 0x83, 0x43, 0x8a, 0x18, 0x72, 0x0c, 0x81, 0xbb, 0x93,
	0xee, 0x58, 0x2e, 0x3e, 0xae, 0x5b, 0xb7, 0x22, 0x86, 0x76, 0xab, 0xa0, 0x53, 0x97, 0xc8, 0x89,
	0x9d, 0x60, 0x29, 0xb6, 0xa3, 0xd8, 0xa4, 0xe1, 0x2d, 0x3a, 0x76, 0xe4, 0x11, 0x3a, 0xf4, 0x21,
	0x3a, 0xa2, 0x4e, 0x1d, 0x2b, 0x58, 0xfa, 0x18, 0x55, 0x9c, 0xb4, 0x6c, 0xc9, 0xef, 0xfb, 0x91,
	0x23, 0xc5, 0x70, 0x18, 0x51, 0x4a, 0x14, 0xca, 0xa7, 0x01, 0xd5, 0x78, 0x8a, 0x52, 0x9c, 0x61,
	0xae, 0xbc, 0x34, 0x93, 0x5a, 0xda, 0x3d, 0xd3, 0xbc, 0xba, 0x0d, 0x07, 0xb1, 0x8c, 0xa5, 0x29,
	0xa8, 0x7c, 0xaa, 0xd0, 0xf0, 0x7b, 0x28, 0x15, 0x97, 0xca, 0xaf, 0x42, 0xf5, 0x52, 0xa5, 0x9f,
	0xf7, 0x0d, 0xd8, 0xb9, 0x34, 0x07, 0xda, 0x1e, 0x6c, 0x63, 0xc2, 0x99, 0x70, 0xc0, 0x18, 0x4c,
	0xbe, 0xcc, 0x9c, 0xa7, 0x87, 0xbf, 0x83, 0xda, 0x9e, 0x11, 0x92, 0x51, 0xa5, 0x96, 0x3a, 0x63,
	0x22, 0x5e, 0x54, 0xcc, 0xfe, 0x0d, 0xfb, 0x38, 0x49, 0xe4, 0x8d, 0x4f, 0x58, 0x14, 0xf9, 0x9a,
	0x71, 0xea, 0x34, 0xc6, 0x60, 0xd2, 0x5c, 0xf4, 0xcc, 0x3c, 0x67, 0x51, 0x74, 0xc5, 0x38, 0xb5,
	0xff, 0xc0, 0xbe, 0xce, 0xb0, 0x50, 0x4c, 0x33, 0x29, 0x2a, 0xd7, 0x34, 0xee, 0xdb, 0x71, 0x36,
	0xf0, 0x07, 0xec, 0x72, 0x26, 0x7c, 0x26, 0x34, 0xcd, 0x72, 0x9c, 0x38, 0x2d, 0xa3, 0xbe, 0x72,
	0x26, 0x2e, 0xea, 0xc9, 0x10, 0x5c, 0x1c, 0x49, 0xbb, 0x26, 0xb8, 0xf8, 0x20, 0xbf, 0x60, 0xaf,
	0xfc, 0x86, 0x4f, 0x58, 0xce, 0x08, 0x15, 0xc4, 0xe9, 0x18, 0xd3, 0x2d, 0xc7, 0x79, 0xbd, 0xd9,
	0xff, 0xe0, 0xa0, 0x3c, 0x47, 0xad, 0xd3, 0x54, 0x66, 0x9a, 0x12, 0x5f, 0x6d, 0x78, 0x20, 0x13,
	0xe7, 0xd3, 0x18, 0x4c, 0x5a, 0x0b, 0x9b, 0xe3, 0x62, 0xf9, 0x9e, 0x96, 0xa6, 0x9c, 0x7e, 0xbe,
	0xdb, 0x8e, 0xac, 0xd7, 0xed, 0x08, 0xcc, 0xce, 0x1f, 0xf7, 0x2e, 0xd8, 0xed, 0x5d, 0xf0, 0xb2,
	0x77, 0xc1, 0xed, 0xc1, 0xb5, 0x76, 0x07, 0xd7, 0x7a, 0x3e, 0xb8, 0xd6, 0xb5, 0x17, 0x33, 0xbd,
	0x5a, 0x07, 0x5e, 0x28, 0x39, 0x0a, 0xb0, 0x20, 0xe6, 0x17, 0x87, 0x32, 0x41, 0xe1, 0x0a, 0x33,
	0x81, 0xf2, 0xff, 0xa8, 0x40, 0xd5, 0x5d, 0xea, 0x4d, 0x4a, 0x55, 0xd0, 0x31, 0xe0, 0xe4, 0x2d,
	0x00, 0x00, 0xff, 0xff, 0x73, 0x1e, 0x63, 0xd7, 0xe1, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.Admin != that1.Admin {
		return false
	}
	if this.AllowDiffTime != that1.AllowDiffTime {
		return false
	}
	if this.TransitionTime != that1.TransitionTime {
		return false
	}
	if this.MinInterval != that1.MinInterval {
		return false
	}
	if this.MaxInterval != that1.MaxInterval {
		return false
	}
	if this.TimeDividend != that1.TimeDividend {
		return false
	}
	if this.MaxSupportedSymbol != that1.MaxSupportedSymbol {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxSupportedSymbol != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSupportedSymbol))
		i--
		dAtA[i] = 0x38
	}
	if m.TimeDividend != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TimeDividend))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxInterval))
		i--
		dAtA[i] = 0x28
	}
	if m.MinInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinInterval))
		i--
		dAtA[i] = 0x20
	}
	if m.TransitionTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TransitionTime))
		i--
		dAtA[i] = 0x18
	}
	if m.AllowDiffTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AllowDiffTime))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.AllowDiffTime != 0 {
		n += 1 + sovParams(uint64(m.AllowDiffTime))
	}
	if m.TransitionTime != 0 {
		n += 1 + sovParams(uint64(m.TransitionTime))
	}
	if m.MinInterval != 0 {
		n += 1 + sovParams(uint64(m.MinInterval))
	}
	if m.MaxInterval != 0 {
		n += 1 + sovParams(uint64(m.MaxInterval))
	}
	if m.TimeDividend != 0 {
		n += 1 + sovParams(uint64(m.TimeDividend))
	}
	if m.MaxSupportedSymbol != 0 {
		n += 1 + sovParams(uint64(m.MaxSupportedSymbol))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowDiffTime", wireType)
			}
			m.AllowDiffTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllowDiffTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransitionTime", wireType)
			}
			m.TransitionTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransitionTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInterval", wireType)
			}
			m.MinInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxInterval", wireType)
			}
			m.MaxInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeDividend", wireType)
			}
			m.TimeDividend = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeDividend |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupportedSymbol", wireType)
			}
			m.MaxSupportedSymbol = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupportedSymbol |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
