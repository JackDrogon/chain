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
	// admin is the address of the admin that is allowed to perform operations on modules.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// allowable_block_time_discrepancy is the allowed discrepancy (in seconds) between validator price timestamp and
	// block_time.
	AllowableBlockTimeDiscrepancy int64 `protobuf:"varint,2,opt,name=allowable_block_time_discrepancy,json=allowableBlockTimeDiscrepancy,proto3" json:"allowable_block_time_discrepancy,omitempty"`
	// grace_period is the time (in seconds) given for validators to adapt to changing in feed's interval.
	GracePeriod int64 `protobuf:"varint,3,opt,name=grace_period,json=gracePeriod,proto3" json:"grace_period,omitempty"`
	// min_interval is the minimum limit of every feeds' interval (in seconds).
	// If the calculated interval is lower than this, it will be capped at this value.
	MinInterval int64 `protobuf:"varint,4,opt,name=min_interval,json=minInterval,proto3" json:"min_interval,omitempty"`
	// max_interval is the maximum limit of every feeds' interval (in seconds).
	// If the calculated interval of a feed is higher than this, it will not be capped at this value.
	MaxInterval int64 `protobuf:"varint,5,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
	// power_step_threshold is the amount of minimum power required to put feed in the current feeds list.
	PowerStepThreshold int64 `protobuf:"varint,6,opt,name=power_step_threshold,json=powerStepThreshold,proto3" json:"power_step_threshold,omitempty"`
	// max_current_feeds is the maximum number of feeds supported at a time.
	MaxCurrentFeeds uint64 `protobuf:"varint,7,opt,name=max_current_feeds,json=maxCurrentFeeds,proto3" json:"max_current_feeds,omitempty"`
	// cooldown_time represents the duration (in seconds) during which validators are prohibited from sending new prices.
	CooldownTime int64 `protobuf:"varint,8,opt,name=cooldown_time,json=cooldownTime,proto3" json:"cooldown_time,omitempty"`
	// min_deviation_basis_point is the minimum limit of every feeds' deviation (in basis point).
	MinDeviationBasisPoint int64 `protobuf:"varint,9,opt,name=min_deviation_basis_point,json=minDeviationBasisPoint,proto3" json:"min_deviation_basis_point,omitempty"`
	// max_deviation_basis_point is the maximum limit of every feeds' deviation (in basis point).
	MaxDeviationBasisPoint int64 `protobuf:"varint,10,opt,name=max_deviation_basis_point,json=maxDeviationBasisPoint,proto3" json:"max_deviation_basis_point,omitempty"`
	// current_feeds_update_interval is the number of blocks after which the current feeds will be re-calculated.
	CurrentFeedsUpdateInterval int64 `protobuf:"varint,11,opt,name=current_feeds_update_interval,json=currentFeedsUpdateInterval,proto3" json:"current_feeds_update_interval,omitempty"`
	// price_quorum is the minimum percentage of power that needs to be reached for a price to be processed.
	PriceQuorum string `protobuf:"bytes,12,opt,name=price_quorum,json=priceQuorum,proto3" json:"price_quorum,omitempty"`
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

func (m *Params) GetAllowableBlockTimeDiscrepancy() int64 {
	if m != nil {
		return m.AllowableBlockTimeDiscrepancy
	}
	return 0
}

func (m *Params) GetGracePeriod() int64 {
	if m != nil {
		return m.GracePeriod
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

func (m *Params) GetPowerStepThreshold() int64 {
	if m != nil {
		return m.PowerStepThreshold
	}
	return 0
}

func (m *Params) GetMaxCurrentFeeds() uint64 {
	if m != nil {
		return m.MaxCurrentFeeds
	}
	return 0
}

func (m *Params) GetCooldownTime() int64 {
	if m != nil {
		return m.CooldownTime
	}
	return 0
}

func (m *Params) GetMinDeviationBasisPoint() int64 {
	if m != nil {
		return m.MinDeviationBasisPoint
	}
	return 0
}

func (m *Params) GetMaxDeviationBasisPoint() int64 {
	if m != nil {
		return m.MaxDeviationBasisPoint
	}
	return 0
}

func (m *Params) GetCurrentFeedsUpdateInterval() int64 {
	if m != nil {
		return m.CurrentFeedsUpdateInterval
	}
	return 0
}

func (m *Params) GetPriceQuorum() string {
	if m != nil {
		return m.PriceQuorum
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "feeds.v1beta1.Params")
}

func init() { proto.RegisterFile("feeds/v1beta1/params.proto", fileDescriptor_bbfae8ad171874f3) }

var fileDescriptor_bbfae8ad171874f3 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbb, 0x6e, 0x13, 0x41,
	0x14, 0x86, 0xbd, 0xc4, 0x31, 0xc9, 0xd8, 0x11, 0x62, 0x65, 0xa1, 0x8d, 0xa5, 0x6c, 0x0c, 0x34,
	0x16, 0x12, 0x5e, 0x02, 0x15, 0x74, 0x31, 0x11, 0x97, 0xce, 0x38, 0xa1, 0xa1, 0x19, 0xcd, 0xce,
	0x0c, 0xf6, 0x88, 0x9d, 0x0b, 0x33, 0x63, 0x7b, 0xf3, 0x16, 0x94, 0x94, 0x79, 0x08, 0x1e, 0x82,
	0x32, 0x82, 0x86, 0x12, 0xd9, 0x0d, 0x8f, 0x81, 0xe6, 0xac, 0x1d, 0x07, 0x89, 0x74, 0xbb, 0xff,
	0xf7, 0xfd, 0x23, 0x9d, 0xa3, 0x83, 0x3a, 0x1f, 0x39, 0x67, 0x2e, 0x9b, 0x1d, 0xe5, 0xdc, 0x93,
	0xa3, 0xcc, 0x10, 0x4b, 0xa4, 0xeb, 0x1b, 0xab, 0xbd, 0x8e, 0xf7, 0x80, 0xf5, 0x57, 0xac, 0xd3,
	0x1e, 0xeb, 0xb1, 0x06, 0x92, 0x85, 0xaf, 0x4a, 0xea, 0xec, 0x53, 0xed, 0xa4, 0x76, 0xb8, 0x02,
	0xd5, 0x4f, 0x85, 0x1e, 0xfc, 0xac, 0xa3, 0xc6, 0x10, 0x1e, 0x8c, 0xfb, 0x68, 0x9b, 0x30, 0x29,
	0x54, 0x12, 0x75, 0xa3, 0xde, 0xee, 0x20, 0xf9, 0xf1, 0xed, 0x71, 0x7b, 0xe5, 0x1e, 0x33, 0x66,
	0xb9, 0x73, 0xa7, 0xde, 0x0a, 0x35, 0x1e, 0x55, 0x5a, 0xfc, 0x1a, 0x75, 0x49, 0x51, 0xe8, 0x39,
	0xc9, 0x0b, 0x8e, 0xf3, 0x42, 0xd3, 0x4f, 0xd8, 0x0b, 0xc9, 0x31, 0x13, 0x8e, 0x5a, 0x6e, 0x88,
	0xa2, 0xe7, 0xc9, 0xad, 0x6e, 0xd4, 0xdb, 0x1a, 0x1d, 0x5c, 0x79, 0x83, 0xa0, 0x9d, 0x09, 0xc9,
	0x4f, 0x36, 0x52, 0x7c, 0x1f, 0xb5, 0xc6, 0x96, 0x50, 0x8e, 0x0d, 0xb7, 0x42, 0xb3, 0x64, 0x0b,
	0x4a, 0x4d, 0xc8, 0x86, 0x10, 0x05, 0x45, 0x0a, 0x85, 0x85, 0xf2, 0xdc, 0xce, 0x48, 0x91, 0xd4,
	0x2b, 0x45, 0x0a, 0xf5, 0x76, 0x15, 0x81, 0x42, 0xca, 0x8d, 0xb2, 0xbd, 0x52, 0x48, 0x79, 0xa5,
	0x3c, 0x41, 0x6d, 0xa3, 0xe7, 0xdc, 0x62, 0xe7, 0xb9, 0xc1, 0x7e, 0x62, 0xb9, 0x9b, 0xe8, 0x82,
	0x25, 0x0d, 0x50, 0x63, 0x60, 0xa7, 0x9e, 0x9b, 0xb3, 0x35, 0x89, 0x1f, 0xa1, 0xbb, 0xe1, 0x51,
	0x3a, 0xb5, 0x96, 0x2b, 0x8f, 0x61, 0xd9, 0xc9, 0xed, 0x6e, 0xd4, 0xab, 0x8f, 0xee, 0x48, 0x52,
	0xbe, 0xac, 0xf2, 0x57, 0x21, 0x8e, 0x1f, 0xa2, 0x3d, 0xaa, 0x75, 0xc1, 0xf4, 0x5c, 0xc1, 0x22,
	0x92, 0x1d, 0x78, 0xb6, 0xb5, 0x0e, 0xc3, 0xd8, 0xf1, 0x73, 0xb4, 0x1f, 0x06, 0x61, 0x7c, 0x26,
	0x88, 0x17, 0x5a, 0xe1, 0x9c, 0x38, 0xe1, 0xb0, 0xd1, 0x42, 0xf9, 0x64, 0x17, 0x0a, 0xf7, 0xa4,
	0x50, 0x27, 0x6b, 0x3e, 0x08, 0x78, 0x18, 0x28, 0x54, 0x49, 0x79, 0x43, 0x15, 0xad, 0xaa, 0xa4,
	0xfc, 0x5f, 0xf5, 0x18, 0x1d, 0xfc, 0x33, 0x02, 0x9e, 0x1a, 0x46, 0x3c, 0xdf, 0x2c, 0xab, 0x09,
	0xf5, 0x0e, 0xbd, 0x36, 0xcf, 0x7b, 0x50, 0xae, 0xaf, 0xd7, 0x58, 0x41, 0x39, 0xfe, 0x3c, 0xd5,
	0x76, 0x2a, 0x93, 0x56, 0x38, 0x92, 0x51, 0x13, 0xb2, 0x77, 0x10, 0xbd, 0xd8, 0xf9, 0x7a, 0x71,
	0x58, 0xfb, 0x73, 0x71, 0x18, 0x0d, 0xde, 0x7c, 0x5f, 0xa4, 0xd1, 0xe5, 0x22, 0x8d, 0x7e, 0x2f,
	0xd2, 0xe8, 0xcb, 0x32, 0xad, 0x5d, 0x2e, 0xd3, 0xda, 0xaf, 0x65, 0x5a, 0xfb, 0xd0, 0x1f, 0x0b,
	0x3f, 0x99, 0xe6, 0x7d, 0xaa, 0x65, 0x96, 0x13, 0xc5, 0xe0, 0x0a, 0xa9, 0x2e, 0x32, 0x3a, 0x21,
	0x42, 0x65, 0xb3, 0xa7, 0x59, 0x99, 0x55, 0xe7, 0xee, 0xcf, 0x0d, 0x77, 0x79, 0x03, 0x84, 0x67,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x68, 0x85, 0x9a, 0x04, 0x03, 0x00, 0x00,
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
	if this.AllowableBlockTimeDiscrepancy != that1.AllowableBlockTimeDiscrepancy {
		return false
	}
	if this.GracePeriod != that1.GracePeriod {
		return false
	}
	if this.MinInterval != that1.MinInterval {
		return false
	}
	if this.MaxInterval != that1.MaxInterval {
		return false
	}
	if this.PowerStepThreshold != that1.PowerStepThreshold {
		return false
	}
	if this.MaxCurrentFeeds != that1.MaxCurrentFeeds {
		return false
	}
	if this.CooldownTime != that1.CooldownTime {
		return false
	}
	if this.MinDeviationBasisPoint != that1.MinDeviationBasisPoint {
		return false
	}
	if this.MaxDeviationBasisPoint != that1.MaxDeviationBasisPoint {
		return false
	}
	if this.CurrentFeedsUpdateInterval != that1.CurrentFeedsUpdateInterval {
		return false
	}
	if this.PriceQuorum != that1.PriceQuorum {
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
	if len(m.PriceQuorum) > 0 {
		i -= len(m.PriceQuorum)
		copy(dAtA[i:], m.PriceQuorum)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PriceQuorum)))
		i--
		dAtA[i] = 0x62
	}
	if m.CurrentFeedsUpdateInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CurrentFeedsUpdateInterval))
		i--
		dAtA[i] = 0x58
	}
	if m.MaxDeviationBasisPoint != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxDeviationBasisPoint))
		i--
		dAtA[i] = 0x50
	}
	if m.MinDeviationBasisPoint != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinDeviationBasisPoint))
		i--
		dAtA[i] = 0x48
	}
	if m.CooldownTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CooldownTime))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxCurrentFeeds != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxCurrentFeeds))
		i--
		dAtA[i] = 0x38
	}
	if m.PowerStepThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PowerStepThreshold))
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
	if m.GracePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.GracePeriod))
		i--
		dAtA[i] = 0x18
	}
	if m.AllowableBlockTimeDiscrepancy != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AllowableBlockTimeDiscrepancy))
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
	if m.AllowableBlockTimeDiscrepancy != 0 {
		n += 1 + sovParams(uint64(m.AllowableBlockTimeDiscrepancy))
	}
	if m.GracePeriod != 0 {
		n += 1 + sovParams(uint64(m.GracePeriod))
	}
	if m.MinInterval != 0 {
		n += 1 + sovParams(uint64(m.MinInterval))
	}
	if m.MaxInterval != 0 {
		n += 1 + sovParams(uint64(m.MaxInterval))
	}
	if m.PowerStepThreshold != 0 {
		n += 1 + sovParams(uint64(m.PowerStepThreshold))
	}
	if m.MaxCurrentFeeds != 0 {
		n += 1 + sovParams(uint64(m.MaxCurrentFeeds))
	}
	if m.CooldownTime != 0 {
		n += 1 + sovParams(uint64(m.CooldownTime))
	}
	if m.MinDeviationBasisPoint != 0 {
		n += 1 + sovParams(uint64(m.MinDeviationBasisPoint))
	}
	if m.MaxDeviationBasisPoint != 0 {
		n += 1 + sovParams(uint64(m.MaxDeviationBasisPoint))
	}
	if m.CurrentFeedsUpdateInterval != 0 {
		n += 1 + sovParams(uint64(m.CurrentFeedsUpdateInterval))
	}
	l = len(m.PriceQuorum)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field AllowableBlockTimeDiscrepancy", wireType)
			}
			m.AllowableBlockTimeDiscrepancy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllowableBlockTimeDiscrepancy |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GracePeriod", wireType)
			}
			m.GracePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GracePeriod |= int64(b&0x7F) << shift
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
				return fmt.Errorf("proto: wrong wireType = %d for field PowerStepThreshold", wireType)
			}
			m.PowerStepThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PowerStepThreshold |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCurrentFeeds", wireType)
			}
			m.MaxCurrentFeeds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCurrentFeeds |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CooldownTime", wireType)
			}
			m.CooldownTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CooldownTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeviationBasisPoint", wireType)
			}
			m.MinDeviationBasisPoint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinDeviationBasisPoint |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDeviationBasisPoint", wireType)
			}
			m.MaxDeviationBasisPoint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxDeviationBasisPoint |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentFeedsUpdateInterval", wireType)
			}
			m.CurrentFeedsUpdateInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentFeedsUpdateInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceQuorum", wireType)
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
			m.PriceQuorum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
