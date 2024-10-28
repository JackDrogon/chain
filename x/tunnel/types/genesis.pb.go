// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tunnel/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState represents the initial state of the blockchain.
type GenesisState struct {
	// params is all parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// tunnel_count is the number of tunnels.
	TunnelCount uint64 `protobuf:"varint,2,opt,name=tunnel_count,json=tunnelCount,proto3" json:"tunnel_count,omitempty"`
	// tunnels is the list of tunnels.
	Tunnels []Tunnel `protobuf:"bytes,3,rep,name=tunnels,proto3" json:"tunnels"`
	// latest_signal_prices_list is the list of latest signal prices.
	LatestSignalPricesList []LatestSignalPrices `protobuf:"bytes,4,rep,name=latest_signal_prices_list,json=latestSignalPricesList,proto3" json:"latest_signal_prices_list"`
	// deposits is the list of deposits.
	Deposits []Deposit `protobuf:"bytes,5,rep,name=deposits,proto3" json:"deposits"`
	// total_fees is the type for the total fees collected by the tunnel
	TotalFees TotalFees `protobuf:"bytes,6,opt,name=total_fees,json=totalFees,proto3" json:"total_fees"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb87d9373a74da2e, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTunnelCount() uint64 {
	if m != nil {
		return m.TunnelCount
	}
	return 0
}

func (m *GenesisState) GetTunnels() []Tunnel {
	if m != nil {
		return m.Tunnels
	}
	return nil
}

func (m *GenesisState) GetLatestSignalPricesList() []LatestSignalPrices {
	if m != nil {
		return m.LatestSignalPricesList
	}
	return nil
}

func (m *GenesisState) GetDeposits() []Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetTotalFees() TotalFees {
	if m != nil {
		return m.TotalFees
	}
	return TotalFees{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tunnel.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("tunnel/v1beta1/genesis.proto", fileDescriptor_cb87d9373a74da2e) }

var fileDescriptor_cb87d9373a74da2e = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4f, 0xf2, 0x30,
	0x18, 0x80, 0x37, 0xe0, 0xe3, 0xd3, 0x42, 0x3c, 0x2c, 0x06, 0x07, 0x9a, 0x89, 0x9c, 0x38, 0xad,
	0x41, 0x8c, 0x89, 0x17, 0x0f, 0x6a, 0x34, 0x26, 0x1c, 0x08, 0x78, 0xf2, 0xb2, 0x74, 0xa3, 0x8e,
	0x26, 0xa5, 0x5d, 0xe8, 0x0b, 0xd1, 0x7f, 0xe1, 0xcf, 0xe2, 0xc8, 0xd1, 0x93, 0x31, 0xf0, 0x43,
	0x34, 0xb4, 0x95, 0x44, 0xd0, 0xdb, 0xfa, 0x3e, 0x4f, 0x9f, 0x75, 0x2b, 0x3a, 0x82, 0x89, 0x10,
	0x94, 0xe3, 0x69, 0x2b, 0xa6, 0x40, 0x5a, 0x38, 0xa5, 0x82, 0x2a, 0xa6, 0xc2, 0x6c, 0x2c, 0x41,
	0x7a, 0x7b, 0x86, 0x86, 0x96, 0xd6, 0xf6, 0x53, 0x99, 0x4a, 0x8d, 0xf0, 0xea, 0xc9, 0x58, 0xb5,
	0xc3, 0x8d, 0x46, 0x46, 0xc6, 0x64, 0xa4, 0xfe, 0x80, 0xb6, 0xa8, 0x61, 0xe3, 0x33, 0x87, 0xca,
	0x77, 0xe6, 0x8d, 0x7d, 0x20, 0x40, 0xbd, 0x33, 0x54, 0x34, 0xbb, 0x7d, 0xb7, 0xee, 0x36, 0x4b,
	0xa7, 0x95, 0xf0, 0xe7, 0x09, 0xc2, 0xae, 0xa6, 0x57, 0x85, 0xd9, 0xfb, 0xb1, 0xd3, 0xb3, 0xae,
	0x77, 0x82, 0xca, 0x46, 0x8b, 0x12, 0x39, 0x11, 0xe0, 0xe7, 0xea, 0x6e, 0xb3, 0xd0, 0x2b, 0x99,
	0xd9, 0xf5, 0x6a, 0xe4, 0x9d, 0xa3, 0xff, 0x66, 0xa9, 0xfc, 0x7c, 0x3d, 0xff, 0x5b, 0xf9, 0x41,
	0x2f, 0x6d, 0xf9, 0x5b, 0xf6, 0x12, 0x54, 0xe5, 0x04, 0xa8, 0x82, 0x48, 0xb1, 0x54, 0x10, 0x1e,
	0x65, 0x63, 0x96, 0x50, 0x15, 0x71, 0xa6, 0xc0, 0x2f, 0xe8, 0x52, 0x63, 0xb3, 0xd4, 0xd1, 0x1b,
	0xfa, 0xda, 0xef, 0x6a, 0xdd, 0x56, 0x2b, 0x7c, 0x8b, 0x74, 0x98, 0x02, 0xef, 0x02, 0xed, 0x0c,
	0x68, 0x26, 0x15, 0x03, 0xe5, 0xff, 0xd3, 0xcd, 0x83, 0xcd, 0xe6, 0x8d, 0xe1, 0x36, 0xb4, 0xd6,
	0xbd, 0x4b, 0x84, 0x40, 0x02, 0xe1, 0xd1, 0x13, 0xa5, 0xca, 0x2f, 0xea, 0x9f, 0x56, 0xdd, 0xfa,
	0xb4, 0x95, 0x71, 0x4b, 0xd7, 0xe7, 0xd8, 0x85, 0xf5, 0xe0, 0x7e, 0xb6, 0x08, 0xdc, 0xf9, 0x22,
	0x70, 0x3f, 0x16, 0x81, 0xfb, 0xba, 0x0c, 0x9c, 0xf9, 0x32, 0x70, 0xde, 0x96, 0x81, 0xf3, 0x88,
	0x53, 0x06, 0xc3, 0x49, 0x1c, 0x26, 0x72, 0x84, 0x63, 0x22, 0x06, 0xfa, 0xc6, 0x12, 0xc9, 0x71,
	0x32, 0x24, 0x4c, 0xe0, 0x69, 0x1b, 0x3f, 0xdb, 0xcb, 0xc4, 0xf0, 0x92, 0x51, 0x15, 0x17, 0xb5,
	0xd1, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x61, 0xda, 0x98, 0xcc, 0x53, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TotalFees.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LatestSignalPricesList) > 0 {
		for iNdEx := len(m.LatestSignalPricesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LatestSignalPricesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Tunnels) > 0 {
		for iNdEx := len(m.Tunnels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tunnels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.TunnelCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TunnelCount))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.TunnelCount != 0 {
		n += 1 + sovGenesis(uint64(m.TunnelCount))
	}
	if len(m.Tunnels) > 0 {
		for _, e := range m.Tunnels {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LatestSignalPricesList) > 0 {
		for _, e := range m.LatestSignalPricesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.TotalFees.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TunnelCount", wireType)
			}
			m.TunnelCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TunnelCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tunnels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tunnels = append(m.Tunnels, Tunnel{})
			if err := m.Tunnels[len(m.Tunnels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestSignalPricesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestSignalPricesList = append(m.LatestSignalPricesList, LatestSignalPrices{})
			if err := m.LatestSignalPricesList[len(m.LatestSignalPricesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
