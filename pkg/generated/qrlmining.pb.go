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
func (*GetBlockMiningCompatibleReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

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
func (*GetLastBlockHeaderReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

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
func (*GetBlockMiningCompatibleResp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

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
func (*GetLastBlockHeaderResp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

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
	WalletAddress []byte `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
}

func (m *GetBlockToMineReq) Reset()                    { *m = GetBlockToMineReq{} }
func (m *GetBlockToMineReq) String() string            { return proto.CompactTextString(m) }
func (*GetBlockToMineReq) ProtoMessage()               {}
func (*GetBlockToMineReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *GetBlockToMineReq) GetWalletAddress() []byte {
	if m != nil {
		return m.WalletAddress
	}
	return nil
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
func (*GetBlockToMineResp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

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
func (*SubmitMinedBlockReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

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
func (*SubmitMinedBlockResp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

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

func init() { proto.RegisterFile("qrlmining.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xe1, 0x6e, 0xd3, 0x30,
	0x10, 0x56, 0xda, 0xae, 0xa2, 0xb7, 0xb5, 0xdb, 0x8e, 0x51, 0x42, 0x36, 0xa1, 0x12, 0x09, 0x31,
	0x24, 0x18, 0x52, 0x11, 0x12, 0xe2, 0x5f, 0x07, 0x52, 0x41, 0xa2, 0x02, 0x05, 0xfe, 0x57, 0xce,
	0x72, 0x5d, 0x22, 0x9c, 0xc6, 0xb5, 0x3d, 0x26, 0xde, 0x81, 0xc7, 0xe0, 0x01, 0x78, 0x0b, 0x5e,
	0x0b, 0xe5, 0xbc, 0x6e, 0x4d, 0xd7, 0x76, 0xff, 0x7c, 0xdf, 0x7d, 0x9f, 0x7d, 0xf7, 0xf9, 0x6c,
	0xd8, 0x9d, 0x69, 0x99, 0x67, 0xd3, 0x6c, 0x7a, 0x7e, 0xa2, 0x74, 0x61, 0x0b, 0xac, 0xcf, 0xb4,
	0x0c, 0x5a, 0x33, 0x2d, 0x5d, 0x1c, 0xbe, 0x81, 0xc3, 0x21, 0xd9, 0x53, 0x59, 0x9c, 0xfd, 0x18,
	0x31, 0xef, 0x7d, 0x91, 0x2b, 0x61, 0xb3, 0x58, 0x52, 0x44, 0x33, 0xec, 0x42, 0x33, 0xa5, 0xec,
	0x3c, 0xb5, 0xbe, 0xd7, 0xf3, 0x8e, 0x1b, 0xd1, 0x55, 0x14, 0xbe, 0x82, 0x07, 0x43, 0xb2, 0x9f,
	0x85, 0x71, 0xd2, 0x8f, 0x24, 0x12, 0xd2, 0x9b, 0x04, 0xbf, 0x3d, 0x38, 0x5a, 0x7f, 0x90, 0x51,
	0xd8, 0x87, 0xed, 0xb8, 0x4c, 0xa6, 0xbc, 0x15, 0xab, 0xb7, 0xfb, 0x7b, 0x27, 0x65, 0xa5, 0x8b,
	0x47, 0x2c, 0x92, 0xf0, 0x2d, 0xb4, 0x39, 0xcc, 0xc9, 0x8a, 0x44, 0x58, 0xe1, 0xd7, 0x58, 0x85,
	0x37, 0xaa, 0x11, 0x59, 0xf1, 0x41, 0x58, 0x11, 0x55, 0x89, 0xe1, 0x5f, 0x0f, 0xba, 0xab, 0x1a,
	0x30, 0x0a, 0x1f, 0x03, 0x24, 0xd9, 0x64, 0x92, 0x9d, 0x5d, 0x48, 0xfb, 0xeb, 0xaa, 0x8b, 0x05,
	0x64, 0xa1, 0xc3, 0xda, 0x62, 0x87, 0x78, 0x04, 0x2d, 0x9b, 0xe5, 0x64, 0xac, 0xc8, 0x95, 0x5f,
	0xe7, 0xd4, 0x0d, 0x50, 0xaa, 0x34, 0x5d, 0x0a, 0x9d, 0xf8, 0x0d, 0xa7, 0x72, 0x11, 0x22, 0x34,
	0x52, 0x61, 0x52, 0x7f, 0xab, 0xe7, 0x1d, 0xb7, 0x22, 0x5e, 0xe3, 0x01, 0x6c, 0x25, 0xa4, 0x6c,
	0xea, 0x37, 0x99, 0xea, 0x82, 0xf0, 0x1d, 0xec, 0xcf, 0x0d, 0xfc, 0x5e, 0x8c, 0xb2, 0x29, 0xdf,
	0xcf, 0x53, 0xe8, 0x5c, 0x0a, 0x29, 0xc9, 0x8e, 0x45, 0x92, 0x68, 0x32, 0x86, 0x0b, 0xde, 0x89,
	0xda, 0x0e, 0x1d, 0x38, 0x30, 0xfc, 0xe3, 0x01, 0x2e, 0x8b, 0x8d, 0xc2, 0x97, 0x80, 0x6c, 0x8b,
	0xa5, 0x5c, 0x49, 0x61, 0x69, 0x1c, 0xcb, 0x22, 0xe6, 0x1d, 0x5a, 0xd1, 0x7e, 0x25, 0x73, 0x2a,
	0x8b, 0x78, 0xc9, 0x99, 0xda, 0x06, 0x67, 0xea, 0x15, 0x67, 0x9e, 0xc1, 0xae, 0x26, 0x43, 0xfa,
	0x27, 0x25, 0xe3, 0x62, 0x32, 0x31, 0x64, 0xd9, 0x84, 0x76, 0xd4, 0x99, 0xc3, 0x5f, 0x18, 0x0d,
	0x9f, 0xc3, 0xfd, 0x6f, 0x17, 0x71, 0x9e, 0xd9, 0xb2, 0xc2, 0x84, 0xab, 0x2d, 0x9b, 0x44, 0x68,
	0x5c, 0x17, 0xb6, 0x13, 0xf1, 0x3a, 0x7c, 0x01, 0x07, 0xb7, 0xa9, 0x46, 0x95, 0xde, 0x91, 0xd6,
	0x85, 0x1b, 0xa0, 0x7b, 0x91, 0x0b, 0xfa, 0xff, 0x6a, 0xd0, 0x72, 0x53, 0x37, 0xf8, 0xfa, 0x09,
	0xc7, 0xe0, 0xaf, 0x1b, 0x45, 0xec, 0xf1, 0xec, 0x6c, 0x78, 0x12, 0xc1, 0x93, 0x3b, 0x18, 0x46,
	0xe1, 0x88, 0xdd, 0x5e, 0x1a, 0x2e, 0x0c, 0xe6, 0xc2, 0xdb, 0xcf, 0x26, 0x38, 0x5c, 0x9b, 0x33,
	0x0a, 0x07, 0xd0, 0xa9, 0x5e, 0x1e, 0x76, 0x2b, 0x35, 0x5c, 0x8f, 0x43, 0xf0, 0x70, 0x25, 0x6e,
	0x14, 0x0e, 0x61, 0x6f, 0xd9, 0x2e, 0xf4, 0x99, 0xbc, 0xc2, 0xf0, 0xe0, 0xd1, 0x9a, 0x8c, 0x51,
	0x71, 0x93, 0xbf, 0x8d, 0xd7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x62, 0x49, 0x8d, 0x59,
	0x04, 0x00, 0x00,
}