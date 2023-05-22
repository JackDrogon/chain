// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tss/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_bandprotocol_chain_v2_pkg_tss "github.com/bandprotocol/chain/v2/pkg/tss"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgCreateGroup is the Msg/CreateGroup request type.
type MsgCreateGroup struct {
	// members is a list of members in this group
	Members []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	// threshold is a minimum number of members required to produce a
	// signature.
	Threshold uint64 `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	// sender is the signer of this message.
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgCreateGroup) Reset()         { *m = MsgCreateGroup{} }
func (m *MsgCreateGroup) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGroup) ProtoMessage()    {}
func (*MsgCreateGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{0}
}
func (m *MsgCreateGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGroup.Merge(m, src)
}
func (m *MsgCreateGroup) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGroup proto.InternalMessageInfo

func (m *MsgCreateGroup) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *MsgCreateGroup) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *MsgCreateGroup) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgCreateGroupResponse is the Msg/CreateGroup response type.
type MsgCreateGroupResponse struct {
}

func (m *MsgCreateGroupResponse) Reset()         { *m = MsgCreateGroupResponse{} }
func (m *MsgCreateGroupResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGroupResponse) ProtoMessage()    {}
func (*MsgCreateGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{1}
}
func (m *MsgCreateGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGroupResponse.Merge(m, src)
}
func (m *MsgCreateGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGroupResponse proto.InternalMessageInfo

// MsgSubmitDKGRound1 is the Msg/SubmitDKGRound1 request type.
type MsgSubmitDKGRound1 struct {
	// group_id is ID of the group.
	GroupID github_com_bandprotocol_chain_v2_pkg_tss.GroupID `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3,casttype=github.com/bandprotocol/chain/v2/pkg/tss.GroupID" json:"group_id,omitempty"`
	// coefficients_commit is commitment of all coefficients.
	CoefficientsCommit github_com_bandprotocol_chain_v2_pkg_tss.Points `protobuf:"bytes,3,rep,name=coefficients_commit,json=coefficientsCommit,proto3,castrepeated=github.com/bandprotocol/chain/v2/pkg/tss.Points" json:"coefficients_commit,omitempty"`
	// one_time_pub_key is a one-time used public key for sending secret shares.
	OneTimePubKey github_com_bandprotocol_chain_v2_pkg_tss.PublicKey `protobuf:"bytes,4,opt,name=one_time_pub_key,json=oneTimePubKey,proto3,casttype=github.com/bandprotocol/chain/v2/pkg/tss.PublicKey" json:"one_time_pub_key,omitempty"`
	// a0_sig is a proof of knowledge on a0.
	A0Sig github_com_bandprotocol_chain_v2_pkg_tss.Signature `protobuf:"bytes,5,opt,name=a0_sig,json=a0Sig,proto3,casttype=github.com/bandprotocol/chain/v2/pkg/tss.Signature" json:"a0_sig,omitempty"`
	// one_time_sig is a proof of knowledge on the OneTimePublicKey.
	OneTimeSig github_com_bandprotocol_chain_v2_pkg_tss.Signature `protobuf:"bytes,6,opt,name=one_time_sig,json=oneTimeSig,proto3,casttype=github.com/bandprotocol/chain/v2/pkg/tss.Signature" json:"one_time_sig,omitempty"`
	// member is the signer of this message. Must be the member of this group.
	Member string `protobuf:"bytes,7,opt,name=member,proto3" json:"member,omitempty"`
}

func (m *MsgSubmitDKGRound1) Reset()         { *m = MsgSubmitDKGRound1{} }
func (m *MsgSubmitDKGRound1) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitDKGRound1) ProtoMessage()    {}
func (*MsgSubmitDKGRound1) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{2}
}
func (m *MsgSubmitDKGRound1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitDKGRound1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitDKGRound1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitDKGRound1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitDKGRound1.Merge(m, src)
}
func (m *MsgSubmitDKGRound1) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitDKGRound1) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitDKGRound1.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitDKGRound1 proto.InternalMessageInfo

func (m *MsgSubmitDKGRound1) GetGroupID() github_com_bandprotocol_chain_v2_pkg_tss.GroupID {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *MsgSubmitDKGRound1) GetCoefficientsCommit() github_com_bandprotocol_chain_v2_pkg_tss.Points {
	if m != nil {
		return m.CoefficientsCommit
	}
	return nil
}

func (m *MsgSubmitDKGRound1) GetOneTimePubKey() github_com_bandprotocol_chain_v2_pkg_tss.PublicKey {
	if m != nil {
		return m.OneTimePubKey
	}
	return nil
}

func (m *MsgSubmitDKGRound1) GetA0Sig() github_com_bandprotocol_chain_v2_pkg_tss.Signature {
	if m != nil {
		return m.A0Sig
	}
	return nil
}

func (m *MsgSubmitDKGRound1) GetOneTimeSig() github_com_bandprotocol_chain_v2_pkg_tss.Signature {
	if m != nil {
		return m.OneTimeSig
	}
	return nil
}

func (m *MsgSubmitDKGRound1) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

// MsgSubmitDKGRound1Response is the Msg/SubmitDKGRound1 response type.
type MsgSubmitDKGRound1Response struct {
}

func (m *MsgSubmitDKGRound1Response) Reset()         { *m = MsgSubmitDKGRound1Response{} }
func (m *MsgSubmitDKGRound1Response) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitDKGRound1Response) ProtoMessage()    {}
func (*MsgSubmitDKGRound1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{3}
}
func (m *MsgSubmitDKGRound1Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitDKGRound1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitDKGRound1Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitDKGRound1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitDKGRound1Response.Merge(m, src)
}
func (m *MsgSubmitDKGRound1Response) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitDKGRound1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitDKGRound1Response.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitDKGRound1Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateGroup)(nil), "tss.v1beta1.MsgCreateGroup")
	proto.RegisterType((*MsgCreateGroupResponse)(nil), "tss.v1beta1.MsgCreateGroupResponse")
	proto.RegisterType((*MsgSubmitDKGRound1)(nil), "tss.v1beta1.MsgSubmitDKGRound1")
	proto.RegisterType((*MsgSubmitDKGRound1Response)(nil), "tss.v1beta1.MsgSubmitDKGRound1Response")
}

func init() { proto.RegisterFile("tss/v1beta1/tx.proto", fileDescriptor_58d13e1023e3ffaf) }

var fileDescriptor_58d13e1023e3ffaf = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x6b, 0xdc, 0x26, 0x64, 0x1b, 0xfe, 0x68, 0x29, 0x91, 0x09, 0x95, 0x13, 0x85, 0x03,
	0x39, 0xa0, 0x6c, 0x92, 0x4a, 0x5c, 0x91, 0xd2, 0x8a, 0xaa, 0x8a, 0x22, 0x2a, 0x87, 0x03, 0x2a,
	0x07, 0xe3, 0xb5, 0x27, 0x9b, 0x55, 0xe3, 0x5d, 0xcb, 0xbb, 0xae, 0x9a, 0xb7, 0xe0, 0x39, 0xb8,
	0xf2, 0x12, 0x1c, 0x7b, 0xe4, 0x14, 0x50, 0xf2, 0x12, 0xa8, 0x27, 0x64, 0x27, 0x81, 0x24, 0x48,
	0xb4, 0x82, 0x9b, 0x67, 0xe6, 0x9b, 0xdf, 0xec, 0x8e, 0xbf, 0x45, 0x7b, 0x5a, 0x29, 0x72, 0xd1,
	0xa2, 0xa0, 0xbd, 0x16, 0xd1, 0x97, 0x8d, 0x28, 0x96, 0x5a, 0xe2, 0x5d, 0xad, 0x54, 0x63, 0x91,
	0x2d, 0xef, 0x31, 0xc9, 0x64, 0x96, 0x27, 0xe9, 0xd7, 0x5c, 0x52, 0x7e, 0xc2, 0xa4, 0x64, 0x23,
	0x20, 0x59, 0x44, 0x93, 0x01, 0xf1, 0xc4, 0x78, 0x51, 0x7a, 0xbc, 0xc6, 0x54, 0x6a, 0x9e, 0xae,
	0x7d, 0x40, 0xf7, 0x7b, 0x8a, 0x1d, 0xc6, 0xe0, 0x69, 0x38, 0x8e, 0x65, 0x12, 0x61, 0x0b, 0xe5,
	0x43, 0x08, 0x29, 0xc4, 0xca, 0x32, 0xaa, 0x66, 0xbd, 0xe0, 0x2c, 0x43, 0xbc, 0x8f, 0x0a, 0x7a,
	0x18, 0x83, 0x1a, 0xca, 0x51, 0x60, 0xdd, 0xa9, 0x1a, 0xf5, 0x6d, 0xe7, 0x77, 0x02, 0x97, 0x50,
	0x4e, 0x81, 0x08, 0x20, 0xb6, 0xcc, 0xaa, 0x51, 0x2f, 0x38, 0x8b, 0xa8, 0x66, 0xa1, 0xd2, 0xfa,
	0x04, 0x07, 0x54, 0x24, 0x85, 0x82, 0xda, 0x0f, 0x13, 0xe1, 0x9e, 0x62, 0xfd, 0x84, 0x86, 0x5c,
	0x1f, 0x75, 0x8f, 0x1d, 0x99, 0x88, 0xa0, 0x85, 0xcf, 0xd0, 0x5d, 0x96, 0xea, 0x5c, 0x1e, 0x58,
	0x46, 0x3a, 0xa5, 0xf3, 0x6a, 0x3a, 0xa9, 0xe4, 0xb3, 0xde, 0x93, 0xa3, 0xeb, 0x49, 0xa5, 0xc9,
	0xb8, 0x1e, 0x26, 0xb4, 0xe1, 0xcb, 0x90, 0x50, 0x4f, 0x04, 0xd9, 0x4d, 0x7c, 0x39, 0x22, 0xfe,
	0xd0, 0xe3, 0x82, 0x5c, 0xb4, 0x49, 0x74, 0xce, 0xb2, 0x3b, 0x2e, 0x7a, 0x9c, 0x7c, 0x06, 0x3c,
	0x09, 0x70, 0x80, 0x1e, 0xf9, 0x12, 0x06, 0x03, 0xee, 0x73, 0x10, 0x5a, 0xb9, 0xbe, 0x0c, 0x43,
	0xae, 0x2d, 0xb3, 0x6a, 0xd6, 0x8b, 0x9d, 0x83, 0x4f, 0xdf, 0x2a, 0xe4, 0xd6, 0xec, 0x53, 0xc9,
	0x85, 0x56, 0x0e, 0x5e, 0xe5, 0x1d, 0x66, 0x38, 0xec, 0xa2, 0x87, 0x52, 0x80, 0xab, 0x79, 0x08,
	0x6e, 0x94, 0x50, 0xf7, 0x1c, 0xc6, 0xd6, 0x76, 0xd5, 0xa8, 0x17, 0x3b, 0x2f, 0xaf, 0x27, 0x95,
	0xf6, 0xed, 0x47, 0x24, 0x74, 0xc4, 0xfd, 0x2e, 0x8c, 0x9d, 0x7b, 0x52, 0xc0, 0x5b, 0x1e, 0xc2,
	0x69, 0x42, 0xbb, 0x30, 0xc6, 0x3d, 0x94, 0xf3, 0x9a, 0xae, 0xe2, 0xcc, 0xda, 0xf9, 0x07, 0x6c,
	0x9f, 0x33, 0xe1, 0xe9, 0x24, 0x06, 0x67, 0xc7, 0x6b, 0xf6, 0x39, 0xc3, 0xef, 0x50, 0xf1, 0xd7,
	0x79, 0x53, 0x68, 0xee, 0xbf, 0xa0, 0x68, 0x71, 0xd6, 0x94, 0x5c, 0x42, 0xb9, 0xb9, 0x7b, 0xac,
	0xfc, 0xdc, 0x14, 0xf3, 0xa8, 0xb6, 0x8f, 0xca, 0x7f, 0xfe, 0xf9, 0xa5, 0x31, 0xda, 0x9f, 0x0d,
	0x64, 0xf6, 0x14, 0xc3, 0x6f, 0xd0, 0xee, 0xaa, 0x33, 0x9f, 0x36, 0x56, 0x5e, 0x40, 0x63, 0xdd,
	0x54, 0xe5, 0x67, 0x7f, 0x29, 0x2e, 0xc1, 0xf8, 0x3d, 0x7a, 0xb0, 0xe9, 0xb6, 0xca, 0x66, 0xdf,
	0x86, 0xa0, 0xfc, 0xfc, 0x06, 0xc1, 0x12, 0xde, 0x79, 0xfd, 0x65, 0x6a, 0x1b, 0x57, 0x53, 0xdb,
	0xf8, 0x3e, 0xb5, 0x8d, 0x8f, 0x33, 0x7b, 0xeb, 0x6a, 0x66, 0x6f, 0x7d, 0x9d, 0xd9, 0x5b, 0x67,
	0x2f, 0x6e, 0xdc, 0xe2, 0x65, 0xba, 0x43, 0xa2, 0xc7, 0x11, 0x28, 0x9a, 0xcb, 0xca, 0x07, 0x3f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x07, 0x7b, 0x59, 0xe7, 0x06, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateGroup creates a new group with a list of members.
	CreateGroup(ctx context.Context, in *MsgCreateGroup, opts ...grpc.CallOption) (*MsgCreateGroupResponse, error)
	// SubmmitDKGRound1 submit dkg for compute round 1.
	SubmitDKGRound1(ctx context.Context, in *MsgSubmitDKGRound1, opts ...grpc.CallOption) (*MsgSubmitDKGRound1Response, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateGroup(ctx context.Context, in *MsgCreateGroup, opts ...grpc.CallOption) (*MsgCreateGroupResponse, error) {
	out := new(MsgCreateGroupResponse)
	err := c.cc.Invoke(ctx, "/tss.v1beta1.Msg/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitDKGRound1(ctx context.Context, in *MsgSubmitDKGRound1, opts ...grpc.CallOption) (*MsgSubmitDKGRound1Response, error) {
	out := new(MsgSubmitDKGRound1Response)
	err := c.cc.Invoke(ctx, "/tss.v1beta1.Msg/SubmitDKGRound1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateGroup creates a new group with a list of members.
	CreateGroup(context.Context, *MsgCreateGroup) (*MsgCreateGroupResponse, error)
	// SubmmitDKGRound1 submit dkg for compute round 1.
	SubmitDKGRound1(context.Context, *MsgSubmitDKGRound1) (*MsgSubmitDKGRound1Response, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateGroup(ctx context.Context, req *MsgCreateGroup) (*MsgCreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (*UnimplementedMsgServer) SubmitDKGRound1(ctx context.Context, req *MsgSubmitDKGRound1) (*MsgSubmitDKGRound1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDKGRound1 not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tss.v1beta1.Msg/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateGroup(ctx, req.(*MsgCreateGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitDKGRound1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitDKGRound1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitDKGRound1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tss.v1beta1.Msg/SubmitDKGRound1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitDKGRound1(ctx, req.(*MsgSubmitDKGRound1))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tss.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _Msg_CreateGroup_Handler,
		},
		{
			MethodName: "SubmitDKGRound1",
			Handler:    _Msg_SubmitDKGRound1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tss/v1beta1/tx.proto",
}

func (m *MsgCreateGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Threshold != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Members[iNdEx])
			copy(dAtA[i:], m.Members[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Members[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitDKGRound1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitDKGRound1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitDKGRound1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Member) > 0 {
		i -= len(m.Member)
		copy(dAtA[i:], m.Member)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Member)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.OneTimeSig) > 0 {
		i -= len(m.OneTimeSig)
		copy(dAtA[i:], m.OneTimeSig)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OneTimeSig)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.A0Sig) > 0 {
		i -= len(m.A0Sig)
		copy(dAtA[i:], m.A0Sig)
		i = encodeVarintTx(dAtA, i, uint64(len(m.A0Sig)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OneTimePubKey) > 0 {
		i -= len(m.OneTimePubKey)
		copy(dAtA[i:], m.OneTimePubKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OneTimePubKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CoefficientsCommit) > 0 {
		for iNdEx := len(m.CoefficientsCommit) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CoefficientsCommit[iNdEx])
			copy(dAtA[i:], m.CoefficientsCommit[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.CoefficientsCommit[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.GroupID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GroupID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitDKGRound1Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitDKGRound1Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitDKGRound1Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Members) > 0 {
		for _, s := range m.Members {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Threshold != 0 {
		n += 1 + sovTx(uint64(m.Threshold))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitDKGRound1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupID != 0 {
		n += 1 + sovTx(uint64(m.GroupID))
	}
	if len(m.CoefficientsCommit) > 0 {
		for _, b := range m.CoefficientsCommit {
			l = len(b)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.OneTimePubKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.A0Sig)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.OneTimeSig)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Member)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitDKGRound1Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateGroupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitDKGRound1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitDKGRound1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitDKGRound1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupID |= github_com_bandprotocol_chain_v2_pkg_tss.GroupID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoefficientsCommit", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoefficientsCommit = append(m.CoefficientsCommit, make([]byte, postIndex-iNdEx))
			copy(m.CoefficientsCommit[len(m.CoefficientsCommit)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneTimePubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OneTimePubKey = append(m.OneTimePubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.OneTimePubKey == nil {
				m.OneTimePubKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A0Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A0Sig = append(m.A0Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.A0Sig == nil {
				m.A0Sig = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneTimeSig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OneTimeSig = append(m.OneTimeSig[:0], dAtA[iNdEx:postIndex]...)
			if m.OneTimeSig == nil {
				m.OneTimeSig = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Member = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitDKGRound1Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitDKGRound1Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitDKGRound1Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
