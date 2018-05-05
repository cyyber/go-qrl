// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qrlmining.proto

package generated

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetBlockMiningCompatibleReq struct {
	// Used for getlastblockheader and getblockheaderbyheight
	Height uint64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *GetBlockMiningCompatibleReq) Reset()                    { *m = GetBlockMiningCompatibleReq{} }
func (m *GetBlockMiningCompatibleReq) String() string            { return proto.CompactTextString(m) }
func (*GetBlockMiningCompatibleReq) ProtoMessage()               {}
func (*GetBlockMiningCompatibleReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GetBlockMiningCompatibleReq) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetLastBlockHeaderReq struct {
	Height uint64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *GetLastBlockHeaderReq) Reset()                    { *m = GetLastBlockHeaderReq{} }
func (m *GetLastBlockHeaderReq) String() string            { return proto.CompactTextString(m) }
func (*GetLastBlockHeaderReq) ProtoMessage()               {}
func (*GetLastBlockHeaderReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GetLastBlockHeaderReq) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetBlockMiningCompatibleResp struct {
	Blockheader   *BlockHeader   `protobuf:"bytes,1,opt,name=blockheader" json:"blockheader,omitempty"`
	Blockmetadata *BlockMetaData `protobuf:"bytes,2,opt,name=blockmetadata" json:"blockmetadata,omitempty"`
}

func (m *GetBlockMiningCompatibleResp) Reset()                    { *m = GetBlockMiningCompatibleResp{} }
func (m *GetBlockMiningCompatibleResp) String() string            { return proto.CompactTextString(m) }
func (*GetBlockMiningCompatibleResp) ProtoMessage()               {}
func (*GetBlockMiningCompatibleResp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetBlockMiningCompatibleResp) GetBlockheader() *BlockHeader {
	if m != nil {
		return m.Blockheader
	}
	return nil
}

func (m *GetBlockMiningCompatibleResp) GetBlockmetadata() *BlockMetaData {
	if m != nil {
		return m.Blockmetadata
	}
	return nil
}

type GetLastBlockHeaderResp struct {
	Difficulty uint64 `protobuf:"varint,1,opt,name=difficulty" json:"difficulty,omitempty"`
	Height     uint64 `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Timestamp  uint64 `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Reward     uint64 `protobuf:"varint,4,opt,name=reward" json:"reward,omitempty"`
	Hash       string `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	Depth      uint64 `protobuf:"varint,6,opt,name=depth" json:"depth,omitempty"`
}

func (m *GetLastBlockHeaderResp) Reset()                    { *m = GetLastBlockHeaderResp{} }
func (m *GetLastBlockHeaderResp) String() string            { return proto.CompactTextString(m) }
func (*GetLastBlockHeaderResp) ProtoMessage()               {}
func (*GetLastBlockHeaderResp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetLastBlockHeaderResp) GetDifficulty() uint64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *GetLastBlockHeaderResp) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetLastBlockHeaderResp) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *GetLastBlockHeaderResp) GetReward() uint64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *GetLastBlockHeaderResp) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *GetLastBlockHeaderResp) GetDepth() uint64 {
	if m != nil {
		return m.Depth
	}
	return 0
}

type GetBlockToMineReq struct {
	WalletAddress string `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress" json:"wallet_address,omitempty"`
}

func (m *GetBlockToMineReq) Reset()                    { *m = GetBlockToMineReq{} }
func (m *GetBlockToMineReq) String() string            { return proto.CompactTextString(m) }
func (*GetBlockToMineReq) ProtoMessage()               {}
func (*GetBlockToMineReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *GetBlockToMineReq) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

type GetBlockToMineResp struct {
	BlocktemplateBlob string `protobuf:"bytes,1,opt,name=blocktemplate_blob,json=blocktemplateBlob" json:"blocktemplate_blob,omitempty"`
	Difficulty        uint64 `protobuf:"varint,2,opt,name=difficulty" json:"difficulty,omitempty"`
	Height            uint64 `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	ReservedOffset    uint32 `protobuf:"varint,4,opt,name=reserved_offset,json=reservedOffset" json:"reserved_offset,omitempty"`
}

func (m *GetBlockToMineResp) Reset()                    { *m = GetBlockToMineResp{} }
func (m *GetBlockToMineResp) String() string            { return proto.CompactTextString(m) }
func (*GetBlockToMineResp) ProtoMessage()               {}
func (*GetBlockToMineResp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GetBlockToMineResp) GetBlocktemplateBlob() string {
	if m != nil {
		return m.BlocktemplateBlob
	}
	return ""
}

func (m *GetBlockToMineResp) GetDifficulty() uint64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *GetBlockToMineResp) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetBlockToMineResp) GetReservedOffset() uint32 {
	if m != nil {
		return m.ReservedOffset
	}
	return 0
}

type SubmitMinedBlockReq struct {
	Blob []byte `protobuf:"bytes,1,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (m *SubmitMinedBlockReq) Reset()                    { *m = SubmitMinedBlockReq{} }
func (m *SubmitMinedBlockReq) String() string            { return proto.CompactTextString(m) }
func (*SubmitMinedBlockReq) ProtoMessage()               {}
func (*SubmitMinedBlockReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *SubmitMinedBlockReq) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

type SubmitMinedBlockResp struct {
	Error bool `protobuf:"varint,1,opt,name=error" json:"error,omitempty"`
}

func (m *SubmitMinedBlockResp) Reset()                    { *m = SubmitMinedBlockResp{} }
func (m *SubmitMinedBlockResp) String() string            { return proto.CompactTextString(m) }
func (*SubmitMinedBlockResp) ProtoMessage()               {}
func (*SubmitMinedBlockResp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *SubmitMinedBlockResp) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func init() {
	proto.RegisterType((*GetBlockMiningCompatibleReq)(nil), "qrl.GetBlockMiningCompatibleReq")
	proto.RegisterType((*GetLastBlockHeaderReq)(nil), "qrl.GetLastBlockHeaderReq")
	proto.RegisterType((*GetBlockMiningCompatibleResp)(nil), "qrl.GetBlockMiningCompatibleResp")
	proto.RegisterType((*GetLastBlockHeaderResp)(nil), "qrl.GetLastBlockHeaderResp")
	proto.RegisterType((*GetBlockToMineReq)(nil), "qrl.GetBlockToMineReq")
	proto.RegisterType((*GetBlockToMineResp)(nil), "qrl.GetBlockToMineResp")
	proto.RegisterType((*SubmitMinedBlockReq)(nil), "qrl.SubmitMinedBlockReq")
	proto.RegisterType((*SubmitMinedBlockResp)(nil), "qrl.SubmitMinedBlockResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MiningAPI service

type MiningAPIClient interface {
	GetBlockMiningCompatible(ctx context.Context, in *GetBlockMiningCompatibleReq, opts ...grpc.CallOption) (*GetBlockMiningCompatibleResp, error)
	GetLastBlockHeader(ctx context.Context, in *GetLastBlockHeaderReq, opts ...grpc.CallOption) (*GetLastBlockHeaderResp, error)
	GetBlockToMine(ctx context.Context, in *GetBlockToMineReq, opts ...grpc.CallOption) (*GetBlockToMineResp, error)
	SubmitMinedBlock(ctx context.Context, in *SubmitMinedBlockReq, opts ...grpc.CallOption) (*SubmitMinedBlockResp, error)
}

type miningAPIClient struct {
	cc *grpc.ClientConn
}

func NewMiningAPIClient(cc *grpc.ClientConn) MiningAPIClient {
	return &miningAPIClient{cc}
}

func (c *miningAPIClient) GetBlockMiningCompatible(ctx context.Context, in *GetBlockMiningCompatibleReq, opts ...grpc.CallOption) (*GetBlockMiningCompatibleResp, error) {
	out := new(GetBlockMiningCompatibleResp)
	err := grpc.Invoke(ctx, "/qrl.MiningAPI/GetBlockMiningCompatible", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miningAPIClient) GetLastBlockHeader(ctx context.Context, in *GetLastBlockHeaderReq, opts ...grpc.CallOption) (*GetLastBlockHeaderResp, error) {
	out := new(GetLastBlockHeaderResp)
	err := grpc.Invoke(ctx, "/qrl.MiningAPI/GetLastBlockHeader", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miningAPIClient) GetBlockToMine(ctx context.Context, in *GetBlockToMineReq, opts ...grpc.CallOption) (*GetBlockToMineResp, error) {
	out := new(GetBlockToMineResp)
	err := grpc.Invoke(ctx, "/qrl.MiningAPI/GetBlockToMine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miningAPIClient) SubmitMinedBlock(ctx context.Context, in *SubmitMinedBlockReq, opts ...grpc.CallOption) (*SubmitMinedBlockResp, error) {
	out := new(SubmitMinedBlockResp)
	err := grpc.Invoke(ctx, "/qrl.MiningAPI/SubmitMinedBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MiningAPI service

type MiningAPIServer interface {
	GetBlockMiningCompatible(context.Context, *GetBlockMiningCompatibleReq) (*GetBlockMiningCompatibleResp, error)
	GetLastBlockHeader(context.Context, *GetLastBlockHeaderReq) (*GetLastBlockHeaderResp, error)
	GetBlockToMine(context.Context, *GetBlockToMineReq) (*GetBlockToMineResp, error)
	SubmitMinedBlock(context.Context, *SubmitMinedBlockReq) (*SubmitMinedBlockResp, error)
}

func RegisterMiningAPIServer(s *grpc.Server, srv MiningAPIServer) {
	s.RegisterService(&_MiningAPI_serviceDesc, srv)
}

func _MiningAPI_GetBlockMiningCompatible_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockMiningCompatibleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiningAPIServer).GetBlockMiningCompatible(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qrl.MiningAPI/GetBlockMiningCompatible",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiningAPIServer).GetBlockMiningCompatible(ctx, req.(*GetBlockMiningCompatibleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiningAPI_GetLastBlockHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastBlockHeaderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiningAPIServer).GetLastBlockHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qrl.MiningAPI/GetLastBlockHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiningAPIServer).GetLastBlockHeader(ctx, req.(*GetLastBlockHeaderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiningAPI_GetBlockToMine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockToMineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiningAPIServer).GetBlockToMine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qrl.MiningAPI/GetBlockToMine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiningAPIServer).GetBlockToMine(ctx, req.(*GetBlockToMineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiningAPI_SubmitMinedBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitMinedBlockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiningAPIServer).SubmitMinedBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qrl.MiningAPI/SubmitMinedBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiningAPIServer).SubmitMinedBlock(ctx, req.(*SubmitMinedBlockReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MiningAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qrl.MiningAPI",
	HandlerType: (*MiningAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockMiningCompatible",
			Handler:    _MiningAPI_GetBlockMiningCompatible_Handler,
		},
		{
			MethodName: "GetLastBlockHeader",
			Handler:    _MiningAPI_GetLastBlockHeader_Handler,
		},
		{
			MethodName: "GetBlockToMine",
			Handler:    _MiningAPI_GetBlockToMine_Handler,
		},
		{
			MethodName: "SubmitMinedBlock",
			Handler:    _MiningAPI_SubmitMinedBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qrlmining.proto",
}

func init() { proto.RegisterFile("qrlmining.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xe1, 0x6e, 0xd3, 0x30,
	0x10, 0x56, 0xda, 0xae, 0xa2, 0x37, 0xda, 0x6d, 0xc7, 0x28, 0x21, 0x9b, 0x50, 0x89, 0x84, 0x18,
	0x12, 0x0c, 0xa9, 0x08, 0x09, 0xf1, 0xaf, 0x03, 0xa9, 0x20, 0x51, 0x81, 0x02, 0xff, 0x2b, 0x67,
	0xb9, 0x2e, 0x11, 0x4e, 0xe3, 0xda, 0x1e, 0x13, 0xef, 0xc0, 0x63, 0xf0, 0x00, 0xbc, 0x05, 0xaf,
	0x85, 0x72, 0x6e, 0xb7, 0xa6, 0x6b, 0xcb, 0x3f, 0xdf, 0x77, 0xdf, 0x67, 0xdf, 0x7d, 0x3e, 0x1b,
	0xf6, 0x66, 0x5a, 0xe6, 0xd9, 0x34, 0x9b, 0x5e, 0x9c, 0x2a, 0x5d, 0xd8, 0x02, 0xeb, 0x33, 0x2d,
	0x83, 0xd6, 0x4c, 0x4b, 0x17, 0x87, 0xaf, 0xe1, 0x68, 0x48, 0xf6, 0x4c, 0x16, 0xe7, 0xdf, 0x47,
	0xcc, 0x7b, 0x57, 0xe4, 0x4a, 0xd8, 0x2c, 0x96, 0x14, 0xd1, 0x0c, 0xbb, 0xd0, 0x4c, 0x29, 0xbb,
	0x48, 0xad, 0xef, 0xf5, 0xbc, 0x93, 0x46, 0x34, 0x8f, 0xc2, 0x97, 0x70, 0x7f, 0x48, 0xf6, 0x93,
	0x30, 0x4e, 0xfa, 0x81, 0x44, 0x42, 0x7a, 0x9b, 0xe0, 0x97, 0x07, 0xc7, 0x9b, 0x0f, 0x32, 0x0a,
	0xfb, 0xb0, 0x1b, 0x97, 0xc9, 0x94, 0xb7, 0x62, 0xf5, 0x6e, 0x7f, 0xff, 0xb4, 0xac, 0x74, 0xf9,
	0x88, 0x65, 0x12, 0xbe, 0x81, 0x36, 0x87, 0x39, 0x59, 0x91, 0x08, 0x2b, 0xfc, 0x1a, 0xab, 0xf0,
	0x46, 0x35, 0x22, 0x2b, 0xde, 0x0b, 0x2b, 0xa2, 0x2a, 0x31, 0xfc, 0xe3, 0x41, 0x77, 0x5d, 0x03,
	0x46, 0xe1, 0x23, 0x80, 0x24, 0x9b, 0x4c, 0xb2, 0xf3, 0x4b, 0x69, 0x7f, 0xce, 0xbb, 0x58, 0x42,
	0x96, 0x3a, 0xac, 0x2d, 0x77, 0x88, 0xc7, 0xd0, 0xb2, 0x59, 0x4e, 0xc6, 0x8a, 0x5c, 0xf9, 0x75,
	0x4e, 0xdd, 0x00, 0xa5, 0x4a, 0xd3, 0x95, 0xd0, 0x89, 0xdf, 0x70, 0x2a, 0x17, 0x21, 0x42, 0x23,
	0x15, 0x26, 0xf5, 0x77, 0x7a, 0xde, 0x49, 0x2b, 0xe2, 0x35, 0x1e, 0xc2, 0x4e, 0x42, 0xca, 0xa6,
	0x7e, 0x93, 0xa9, 0x2e, 0x08, 0xdf, 0xc2, 0xc1, 0xc2, 0xc0, 0x6f, 0xc5, 0x28, 0x9b, 0xf2, 0xfd,
	0x3c, 0x81, 0xce, 0x95, 0x90, 0x92, 0xec, 0x58, 0x24, 0x89, 0x26, 0x63, 0xb8, 0xe0, 0x56, 0xd4,
	0x76, 0xe8, 0xc0, 0x81, 0xe1, 0x6f, 0x0f, 0x70, 0x55, 0x6c, 0x14, 0xbe, 0x00, 0x64, 0x5b, 0x2c,
	0xe5, 0x4a, 0x0a, 0x4b, 0xe3, 0x58, 0x16, 0xf1, 0x7c, 0x87, 0x83, 0x4a, 0xe6, 0x4c, 0x16, 0xf1,
	0x8a, 0x33, 0xb5, 0x2d, 0xce, 0xd4, 0x2b, 0xce, 0x3c, 0x85, 0x3d, 0x4d, 0x86, 0xf4, 0x0f, 0x4a,
	0xc6, 0xc5, 0x64, 0x62, 0xc8, 0xb2, 0x09, 0xed, 0xa8, 0xb3, 0x80, 0x3f, 0x33, 0x1a, 0x3e, 0x83,
	0x7b, 0x5f, 0x2f, 0xe3, 0x3c, 0xb3, 0x65, 0x85, 0x09, 0x57, 0x5b, 0x36, 0x89, 0xd0, 0xb8, 0x2e,
	0xec, 0x6e, 0xc4, 0xeb, 0xf0, 0x39, 0x1c, 0xde, 0xa6, 0x1a, 0x55, 0x7a, 0x47, 0x5a, 0x17, 0x6e,
	0x80, 0xee, 0x44, 0x2e, 0xe8, 0xff, 0xad, 0x41, 0xcb, 0x4d, 0xdd, 0xe0, 0xcb, 0x47, 0x1c, 0x83,
	0xbf, 0x69, 0x14, 0xb1, 0xc7, 0xb3, 0xb3, 0xe5, 0x49, 0x04, 0x8f, 0xff, 0xc3, 0x30, 0x0a, 0x47,
	0xec, 0xf6, 0xca, 0x70, 0x61, 0xb0, 0x10, 0xde, 0x7e, 0x36, 0xc1, 0xd1, 0xc6, 0x9c, 0x51, 0x38,
	0x80, 0x4e, 0xf5, 0xf2, 0xb0, 0x5b, 0xa9, 0xe1, 0x7a, 0x1c, 0x82, 0x07, 0x6b, 0x71, 0xa3, 0x70,
	0x08, 0xfb, 0xab, 0x76, 0xa1, 0xcf, 0xe4, 0x35, 0x86, 0x07, 0x0f, 0x37, 0x64, 0x8c, 0x8a, 0x9b,
	0xfc, 0x6d, 0xbc, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0x75, 0xb5, 0x3c, 0x81, 0x59, 0x04, 0x00,
	0x00,
}
