// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qrllegacy.proto

/*
Package qrl is a generated protocol buffer package.

It is generated from these files:
	qrllegacy.proto
	qrlmining.proto
	qrl.proto

It has these top-level messages:
	LegacyMessage
	NoData
	VEData
	PLData
	PONGData
	MRData
	BKData
	FBData
	PBData
	SYNCData
	GetBlockMiningCompatibleReq
	GetLastBlockHeaderReq
	GetBlockMiningCompatibleResp
	GetLastBlockHeaderResp
	GetBlockToMineReq
	GetBlockToMineResp
	SubmitMinedBlockReq
	SubmitMinedBlockResp
	Empty
	GetNodeStateReq
	GetNodeStateResp
	GetKnownPeersReq
	GetKnownPeersResp
	GetPeersStatReq
	GetPeersStatResp
	GetBlockReq
	GetBlockResp
	GetStatsReq
	GetStatsResp
	GetAddressFromPKReq
	GetAddressFromPKResp
	BlockDataPoint
	GetAddressStateReq
	GetAddressStateResp
	GetObjectReq
	GetObjectResp
	GetLatestDataReq
	GetLatestDataResp
	TransferCoinsReq
	TransferCoinsResp
	PushTransactionReq
	PushTransactionResp
	MessageTxnReq
	TokenTxnReq
	TransferTokenTxnReq
	SlaveTxnReq
	LatticePublicKeyTxnReq
	GetLocalAddressesReq
	GetLocalAddressesResp
	NodeInfo
	StoredPeers
	Peer
	AddressState
	LatticePK
	AddressAmount
	BlockHeader
	BlockHeaderExtended
	TransactionCount
	TransactionExtended
	BlockExtended
	Block
	GenesisBalance
	BlockMetaDataList
	Transaction
	TokenList
	TokenMetadata
	CollectEphemeralMessageReq
	CollectEphemeralMessageResp
	PushEphemeralMessageReq
	EncryptedEphemeralMessage
	EphemeralChannelPayload
	EphemeralMessagePayload
	LatticePublicKeys
	EphemeralMetadata
	AddressList
	BlockHeightData
	BlockMetaData
	BlockNumberMapping
	StateLoader
	StateObjects
	LRUStateCache
	PeerStat
	NodeChainState
	NodeHeaderHash
	P2PAcknowledgement
	PeerInfo
	Peers
*/
package generated

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LegacyMessage_FuncName int32

const (
	LegacyMessage_VE           LegacyMessage_FuncName = 0
	LegacyMessage_PL           LegacyMessage_FuncName = 1
	LegacyMessage_PONG         LegacyMessage_FuncName = 2
	LegacyMessage_MR           LegacyMessage_FuncName = 3
	LegacyMessage_SFM          LegacyMessage_FuncName = 4
	LegacyMessage_BK           LegacyMessage_FuncName = 5
	LegacyMessage_FB           LegacyMessage_FuncName = 6
	LegacyMessage_PB           LegacyMessage_FuncName = 7
	LegacyMessage_BH           LegacyMessage_FuncName = 8
	LegacyMessage_TX           LegacyMessage_FuncName = 9
	LegacyMessage_LT           LegacyMessage_FuncName = 10
	LegacyMessage_EPH          LegacyMessage_FuncName = 11
	LegacyMessage_MT           LegacyMessage_FuncName = 12
	LegacyMessage_TK           LegacyMessage_FuncName = 13
	LegacyMessage_TT           LegacyMessage_FuncName = 14
	LegacyMessage_SL           LegacyMessage_FuncName = 15
	LegacyMessage_SYNC         LegacyMessage_FuncName = 16
	LegacyMessage_CHAINSTATE   LegacyMessage_FuncName = 17
	LegacyMessage_HEADERHASHES LegacyMessage_FuncName = 18
	LegacyMessage_P2P_ACK      LegacyMessage_FuncName = 19
)

var LegacyMessage_FuncName_name = map[int32]string{
	0:  "VE",
	1:  "PL",
	2:  "PONG",
	3:  "MR",
	4:  "SFM",
	5:  "BK",
	6:  "FB",
	7:  "PB",
	8:  "BH",
	9:  "TX",
	10: "LT",
	11: "EPH",
	12: "MT",
	13: "TK",
	14: "TT",
	15: "SL",
	16: "SYNC",
	17: "CHAINSTATE",
	18: "HEADERHASHES",
	19: "P2P_ACK",
}
var LegacyMessage_FuncName_value = map[string]int32{
	"VE":           0,
	"PL":           1,
	"PONG":         2,
	"MR":           3,
	"SFM":          4,
	"BK":           5,
	"FB":           6,
	"PB":           7,
	"BH":           8,
	"TX":           9,
	"LT":           10,
	"EPH":          11,
	"MT":           12,
	"TK":           13,
	"TT":           14,
	"SL":           15,
	"SYNC":         16,
	"CHAINSTATE":   17,
	"HEADERHASHES": 18,
	"P2P_ACK":      19,
}

func (x LegacyMessage_FuncName) String() string {
	return proto.EnumName(LegacyMessage_FuncName_name, int32(x))
}
func (LegacyMessage_FuncName) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Adding old code to refactor while keeping things working
type LegacyMessage struct {
	FuncName LegacyMessage_FuncName `protobuf:"varint,1,opt,name=func_name,json=funcName,enum=qrl.LegacyMessage_FuncName" json:"func_name,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*LegacyMessage_NoData
	//	*LegacyMessage_VeData
	//	*LegacyMessage_PlData
	//	*LegacyMessage_PongData
	//	*LegacyMessage_MrData
	//	*LegacyMessage_Block
	//	*LegacyMessage_FbData
	//	*LegacyMessage_PbData
	//	*LegacyMessage_BhData
	//	*LegacyMessage_TxData
	//	*LegacyMessage_MtData
	//	*LegacyMessage_TkData
	//	*LegacyMessage_TtData
	//	*LegacyMessage_LtData
	//	*LegacyMessage_SlData
	//	*LegacyMessage_EphData
	//	*LegacyMessage_SyncData
	//	*LegacyMessage_ChainStateData
	//	*LegacyMessage_NodeHeaderHash
	//	*LegacyMessage_P2PAckData
	Data isLegacyMessage_Data `protobuf_oneof:"data"`
}

func (m *LegacyMessage) Reset()                    { *m = LegacyMessage{} }
func (m *LegacyMessage) String() string            { return proto.CompactTextString(m) }
func (*LegacyMessage) ProtoMessage()               {}
func (*LegacyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isLegacyMessage_Data interface {
	isLegacyMessage_Data()
}

type LegacyMessage_NoData struct {
	NoData *NoData `protobuf:"bytes,2,opt,name=noData,oneof"`
}
type LegacyMessage_VeData struct {
	VeData *VEData `protobuf:"bytes,3,opt,name=veData,oneof"`
}
type LegacyMessage_PlData struct {
	PlData *PLData `protobuf:"bytes,4,opt,name=plData,oneof"`
}
type LegacyMessage_PongData struct {
	PongData *PONGData `protobuf:"bytes,5,opt,name=pongData,oneof"`
}
type LegacyMessage_MrData struct {
	MrData *MRData `protobuf:"bytes,6,opt,name=mrData,oneof"`
}
type LegacyMessage_Block struct {
	Block *Block `protobuf:"bytes,7,opt,name=block,oneof"`
}
type LegacyMessage_FbData struct {
	FbData *FBData `protobuf:"bytes,8,opt,name=fbData,oneof"`
}
type LegacyMessage_PbData struct {
	PbData *PBData `protobuf:"bytes,9,opt,name=pbData,oneof"`
}
type LegacyMessage_BhData struct {
	BhData *BlockHeightData `protobuf:"bytes,10,opt,name=bhData,oneof"`
}
type LegacyMessage_TxData struct {
	TxData *Transaction `protobuf:"bytes,11,opt,name=txData,oneof"`
}
type LegacyMessage_MtData struct {
	MtData *Transaction `protobuf:"bytes,12,opt,name=mtData,oneof"`
}
type LegacyMessage_TkData struct {
	TkData *Transaction `protobuf:"bytes,13,opt,name=tkData,oneof"`
}
type LegacyMessage_TtData struct {
	TtData *Transaction `protobuf:"bytes,14,opt,name=ttData,oneof"`
}
type LegacyMessage_LtData struct {
	LtData *Transaction `protobuf:"bytes,15,opt,name=ltData,oneof"`
}
type LegacyMessage_SlData struct {
	SlData *Transaction `protobuf:"bytes,16,opt,name=slData,oneof"`
}
type LegacyMessage_EphData struct {
	EphData *EncryptedEphemeralMessage `protobuf:"bytes,17,opt,name=ephData,oneof"`
}
type LegacyMessage_SyncData struct {
	SyncData *SYNCData `protobuf:"bytes,18,opt,name=syncData,oneof"`
}
type LegacyMessage_ChainStateData struct {
	ChainStateData *NodeChainState `protobuf:"bytes,19,opt,name=chainStateData,oneof"`
}
type LegacyMessage_NodeHeaderHash struct {
	NodeHeaderHash *NodeHeaderHash `protobuf:"bytes,20,opt,name=nodeHeaderHash,oneof"`
}
type LegacyMessage_P2PAckData struct {
	P2PAckData *P2PAcknowledgement `protobuf:"bytes,21,opt,name=p2pAckData,oneof"`
}

func (*LegacyMessage_NoData) isLegacyMessage_Data()         {}
func (*LegacyMessage_VeData) isLegacyMessage_Data()         {}
func (*LegacyMessage_PlData) isLegacyMessage_Data()         {}
func (*LegacyMessage_PongData) isLegacyMessage_Data()       {}
func (*LegacyMessage_MrData) isLegacyMessage_Data()         {}
func (*LegacyMessage_Block) isLegacyMessage_Data()          {}
func (*LegacyMessage_FbData) isLegacyMessage_Data()         {}
func (*LegacyMessage_PbData) isLegacyMessage_Data()         {}
func (*LegacyMessage_BhData) isLegacyMessage_Data()         {}
func (*LegacyMessage_TxData) isLegacyMessage_Data()         {}
func (*LegacyMessage_MtData) isLegacyMessage_Data()         {}
func (*LegacyMessage_TkData) isLegacyMessage_Data()         {}
func (*LegacyMessage_TtData) isLegacyMessage_Data()         {}
func (*LegacyMessage_LtData) isLegacyMessage_Data()         {}
func (*LegacyMessage_SlData) isLegacyMessage_Data()         {}
func (*LegacyMessage_EphData) isLegacyMessage_Data()        {}
func (*LegacyMessage_SyncData) isLegacyMessage_Data()       {}
func (*LegacyMessage_ChainStateData) isLegacyMessage_Data() {}
func (*LegacyMessage_NodeHeaderHash) isLegacyMessage_Data() {}
func (*LegacyMessage_P2PAckData) isLegacyMessage_Data()     {}

func (m *LegacyMessage) GetData() isLegacyMessage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *LegacyMessage) GetFuncName() LegacyMessage_FuncName {
	if m != nil {
		return m.FuncName
	}
	return LegacyMessage_VE
}

func (m *LegacyMessage) GetNoData() *NoData {
	if x, ok := m.GetData().(*LegacyMessage_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *LegacyMessage) GetVeData() *VEData {
	if x, ok := m.GetData().(*LegacyMessage_VeData); ok {
		return x.VeData
	}
	return nil
}

func (m *LegacyMessage) GetPlData() *PLData {
	if x, ok := m.GetData().(*LegacyMessage_PlData); ok {
		return x.PlData
	}
	return nil
}

func (m *LegacyMessage) GetPongData() *PONGData {
	if x, ok := m.GetData().(*LegacyMessage_PongData); ok {
		return x.PongData
	}
	return nil
}

func (m *LegacyMessage) GetMrData() *MRData {
	if x, ok := m.GetData().(*LegacyMessage_MrData); ok {
		return x.MrData
	}
	return nil
}

func (m *LegacyMessage) GetBlock() *Block {
	if x, ok := m.GetData().(*LegacyMessage_Block); ok {
		return x.Block
	}
	return nil
}

func (m *LegacyMessage) GetFbData() *FBData {
	if x, ok := m.GetData().(*LegacyMessage_FbData); ok {
		return x.FbData
	}
	return nil
}

func (m *LegacyMessage) GetPbData() *PBData {
	if x, ok := m.GetData().(*LegacyMessage_PbData); ok {
		return x.PbData
	}
	return nil
}

func (m *LegacyMessage) GetBhData() *BlockHeightData {
	if x, ok := m.GetData().(*LegacyMessage_BhData); ok {
		return x.BhData
	}
	return nil
}

func (m *LegacyMessage) GetTxData() *Transaction {
	if x, ok := m.GetData().(*LegacyMessage_TxData); ok {
		return x.TxData
	}
	return nil
}

func (m *LegacyMessage) GetMtData() *Transaction {
	if x, ok := m.GetData().(*LegacyMessage_MtData); ok {
		return x.MtData
	}
	return nil
}

func (m *LegacyMessage) GetTkData() *Transaction {
	if x, ok := m.GetData().(*LegacyMessage_TkData); ok {
		return x.TkData
	}
	return nil
}

func (m *LegacyMessage) GetTtData() *Transaction {
	if x, ok := m.GetData().(*LegacyMessage_TtData); ok {
		return x.TtData
	}
	return nil
}

func (m *LegacyMessage) GetLtData() *Transaction {
	if x, ok := m.GetData().(*LegacyMessage_LtData); ok {
		return x.LtData
	}
	return nil
}

func (m *LegacyMessage) GetSlData() *Transaction {
	if x, ok := m.GetData().(*LegacyMessage_SlData); ok {
		return x.SlData
	}
	return nil
}

func (m *LegacyMessage) GetEphData() *EncryptedEphemeralMessage {
	if x, ok := m.GetData().(*LegacyMessage_EphData); ok {
		return x.EphData
	}
	return nil
}

func (m *LegacyMessage) GetSyncData() *SYNCData {
	if x, ok := m.GetData().(*LegacyMessage_SyncData); ok {
		return x.SyncData
	}
	return nil
}

func (m *LegacyMessage) GetChainStateData() *NodeChainState {
	if x, ok := m.GetData().(*LegacyMessage_ChainStateData); ok {
		return x.ChainStateData
	}
	return nil
}

func (m *LegacyMessage) GetNodeHeaderHash() *NodeHeaderHash {
	if x, ok := m.GetData().(*LegacyMessage_NodeHeaderHash); ok {
		return x.NodeHeaderHash
	}
	return nil
}

func (m *LegacyMessage) GetP2PAckData() *P2PAcknowledgement {
	if x, ok := m.GetData().(*LegacyMessage_P2PAckData); ok {
		return x.P2PAckData
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LegacyMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LegacyMessage_OneofMarshaler, _LegacyMessage_OneofUnmarshaler, _LegacyMessage_OneofSizer, []interface{}{
		(*LegacyMessage_NoData)(nil),
		(*LegacyMessage_VeData)(nil),
		(*LegacyMessage_PlData)(nil),
		(*LegacyMessage_PongData)(nil),
		(*LegacyMessage_MrData)(nil),
		(*LegacyMessage_Block)(nil),
		(*LegacyMessage_FbData)(nil),
		(*LegacyMessage_PbData)(nil),
		(*LegacyMessage_BhData)(nil),
		(*LegacyMessage_TxData)(nil),
		(*LegacyMessage_MtData)(nil),
		(*LegacyMessage_TkData)(nil),
		(*LegacyMessage_TtData)(nil),
		(*LegacyMessage_LtData)(nil),
		(*LegacyMessage_SlData)(nil),
		(*LegacyMessage_EphData)(nil),
		(*LegacyMessage_SyncData)(nil),
		(*LegacyMessage_ChainStateData)(nil),
		(*LegacyMessage_NodeHeaderHash)(nil),
		(*LegacyMessage_P2PAckData)(nil),
	}
}

func _LegacyMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LegacyMessage)
	// data
	switch x := m.Data.(type) {
	case *LegacyMessage_NoData:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NoData); err != nil {
			return err
		}
	case *LegacyMessage_VeData:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VeData); err != nil {
			return err
		}
	case *LegacyMessage_PlData:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PlData); err != nil {
			return err
		}
	case *LegacyMessage_PongData:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PongData); err != nil {
			return err
		}
	case *LegacyMessage_MrData:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MrData); err != nil {
			return err
		}
	case *LegacyMessage_Block:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case *LegacyMessage_FbData:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FbData); err != nil {
			return err
		}
	case *LegacyMessage_PbData:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PbData); err != nil {
			return err
		}
	case *LegacyMessage_BhData:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BhData); err != nil {
			return err
		}
	case *LegacyMessage_TxData:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TxData); err != nil {
			return err
		}
	case *LegacyMessage_MtData:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MtData); err != nil {
			return err
		}
	case *LegacyMessage_TkData:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TkData); err != nil {
			return err
		}
	case *LegacyMessage_TtData:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TtData); err != nil {
			return err
		}
	case *LegacyMessage_LtData:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LtData); err != nil {
			return err
		}
	case *LegacyMessage_SlData:
		b.EncodeVarint(16<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SlData); err != nil {
			return err
		}
	case *LegacyMessage_EphData:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EphData); err != nil {
			return err
		}
	case *LegacyMessage_SyncData:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SyncData); err != nil {
			return err
		}
	case *LegacyMessage_ChainStateData:
		b.EncodeVarint(19<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ChainStateData); err != nil {
			return err
		}
	case *LegacyMessage_NodeHeaderHash:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodeHeaderHash); err != nil {
			return err
		}
	case *LegacyMessage_P2PAckData:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.P2PAckData); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LegacyMessage.Data has unexpected type %T", x)
	}
	return nil
}

func _LegacyMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LegacyMessage)
	switch tag {
	case 2: // data.noData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NoData)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_NoData{msg}
		return true, err
	case 3: // data.veData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VEData)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_VeData{msg}
		return true, err
	case 4: // data.plData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PLData)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_PlData{msg}
		return true, err
	case 5: // data.pongData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PONGData)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_PongData{msg}
		return true, err
	case 6: // data.mrData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MRData)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_MrData{msg}
		return true, err
	case 7: // data.block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Block)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_Block{msg}
		return true, err
	case 8: // data.fbData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FBData)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_FbData{msg}
		return true, err
	case 9: // data.pbData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PBData)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_PbData{msg}
		return true, err
	case 10: // data.bhData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BlockHeightData)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_BhData{msg}
		return true, err
	case 11: // data.txData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transaction)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_TxData{msg}
		return true, err
	case 12: // data.mtData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transaction)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_MtData{msg}
		return true, err
	case 13: // data.tkData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transaction)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_TkData{msg}
		return true, err
	case 14: // data.ttData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transaction)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_TtData{msg}
		return true, err
	case 15: // data.ltData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transaction)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_LtData{msg}
		return true, err
	case 16: // data.slData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transaction)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_SlData{msg}
		return true, err
	case 17: // data.ephData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EncryptedEphemeralMessage)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_EphData{msg}
		return true, err
	case 18: // data.syncData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SYNCData)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_SyncData{msg}
		return true, err
	case 19: // data.chainStateData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NodeChainState)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_ChainStateData{msg}
		return true, err
	case 20: // data.nodeHeaderHash
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NodeHeaderHash)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_NodeHeaderHash{msg}
		return true, err
	case 21: // data.p2pAckData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(P2PAcknowledgement)
		err := b.DecodeMessage(msg)
		m.Data = &LegacyMessage_P2PAckData{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LegacyMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LegacyMessage)
	// data
	switch x := m.Data.(type) {
	case *LegacyMessage_NoData:
		s := proto.Size(x.NoData)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_VeData:
		s := proto.Size(x.VeData)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_PlData:
		s := proto.Size(x.PlData)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_PongData:
		s := proto.Size(x.PongData)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_MrData:
		s := proto.Size(x.MrData)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_Block:
		s := proto.Size(x.Block)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_FbData:
		s := proto.Size(x.FbData)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_PbData:
		s := proto.Size(x.PbData)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_BhData:
		s := proto.Size(x.BhData)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_TxData:
		s := proto.Size(x.TxData)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_MtData:
		s := proto.Size(x.MtData)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_TkData:
		s := proto.Size(x.TkData)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_TtData:
		s := proto.Size(x.TtData)
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_LtData:
		s := proto.Size(x.LtData)
		n += proto.SizeVarint(15<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_SlData:
		s := proto.Size(x.SlData)
		n += proto.SizeVarint(16<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_EphData:
		s := proto.Size(x.EphData)
		n += proto.SizeVarint(17<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_SyncData:
		s := proto.Size(x.SyncData)
		n += proto.SizeVarint(18<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_ChainStateData:
		s := proto.Size(x.ChainStateData)
		n += proto.SizeVarint(19<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_NodeHeaderHash:
		s := proto.Size(x.NodeHeaderHash)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LegacyMessage_P2PAckData:
		s := proto.Size(x.P2PAckData)
		n += proto.SizeVarint(21<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NoData struct {
}

func (m *NoData) Reset()                    { *m = NoData{} }
func (m *NoData) String() string            { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()               {}
func (*NoData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type VEData struct {
	Version         string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	GenesisPrevHash []byte `protobuf:"bytes,2,opt,name=genesis_prev_hash,json=genesisPrevHash,proto3" json:"genesis_prev_hash,omitempty"`
	RateLimit       uint64 `protobuf:"varint,3,opt,name=rate_limit,json=rateLimit" json:"rate_limit,omitempty"`
}

func (m *VEData) Reset()                    { *m = VEData{} }
func (m *VEData) String() string            { return proto.CompactTextString(m) }
func (*VEData) ProtoMessage()               {}
func (*VEData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VEData) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VEData) GetGenesisPrevHash() []byte {
	if m != nil {
		return m.GenesisPrevHash
	}
	return nil
}

func (m *VEData) GetRateLimit() uint64 {
	if m != nil {
		return m.RateLimit
	}
	return 0
}

type PLData struct {
	PeerIps    []string `protobuf:"bytes,1,rep,name=peer_ips,json=peerIps" json:"peer_ips,omitempty"`
	PublicPort uint32   `protobuf:"varint,2,opt,name=public_port,json=publicPort" json:"public_port,omitempty"`
}

func (m *PLData) Reset()                    { *m = PLData{} }
func (m *PLData) String() string            { return proto.CompactTextString(m) }
func (*PLData) ProtoMessage()               {}
func (*PLData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PLData) GetPeerIps() []string {
	if m != nil {
		return m.PeerIps
	}
	return nil
}

func (m *PLData) GetPublicPort() uint32 {
	if m != nil {
		return m.PublicPort
	}
	return 0
}

type PONGData struct {
}

func (m *PONGData) Reset()                    { *m = PONGData{} }
func (m *PONGData) String() string            { return proto.CompactTextString(m) }
func (*PONGData) ProtoMessage()               {}
func (*PONGData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type MRData struct {
	Hash           []byte                 `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Type           LegacyMessage_FuncName `protobuf:"varint,2,opt,name=type,enum=qrl.LegacyMessage_FuncName" json:"type,omitempty"`
	StakeSelector  []byte                 `protobuf:"bytes,3,opt,name=stake_selector,json=stakeSelector,proto3" json:"stake_selector,omitempty"`
	BlockNumber    uint64                 `protobuf:"varint,4,opt,name=block_number,json=blockNumber" json:"block_number,omitempty"`
	PrevHeaderhash []byte                 `protobuf:"bytes,5,opt,name=prev_headerhash,json=prevHeaderhash,proto3" json:"prev_headerhash,omitempty"`
	RevealHash     []byte                 `protobuf:"bytes,6,opt,name=reveal_hash,json=revealHash,proto3" json:"reveal_hash,omitempty"`
}

func (m *MRData) Reset()                    { *m = MRData{} }
func (m *MRData) String() string            { return proto.CompactTextString(m) }
func (*MRData) ProtoMessage()               {}
func (*MRData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MRData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *MRData) GetType() LegacyMessage_FuncName {
	if m != nil {
		return m.Type
	}
	return LegacyMessage_VE
}

func (m *MRData) GetStakeSelector() []byte {
	if m != nil {
		return m.StakeSelector
	}
	return nil
}

func (m *MRData) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *MRData) GetPrevHeaderhash() []byte {
	if m != nil {
		return m.PrevHeaderhash
	}
	return nil
}

func (m *MRData) GetRevealHash() []byte {
	if m != nil {
		return m.RevealHash
	}
	return nil
}

type BKData struct {
	MrData *MRData `protobuf:"bytes,1,opt,name=mrData" json:"mrData,omitempty"`
	Block  *Block  `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
}

func (m *BKData) Reset()                    { *m = BKData{} }
func (m *BKData) String() string            { return proto.CompactTextString(m) }
func (*BKData) ProtoMessage()               {}
func (*BKData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BKData) GetMrData() *MRData {
	if m != nil {
		return m.MrData
	}
	return nil
}

func (m *BKData) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type FBData struct {
	Index uint64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
}

func (m *FBData) Reset()                    { *m = FBData{} }
func (m *FBData) String() string            { return proto.CompactTextString(m) }
func (*FBData) ProtoMessage()               {}
func (*FBData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FBData) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type PBData struct {
	Block *Block `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
}

func (m *PBData) Reset()                    { *m = PBData{} }
func (m *PBData) String() string            { return proto.CompactTextString(m) }
func (*PBData) ProtoMessage()               {}
func (*PBData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PBData) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type SYNCData struct {
	State string `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
}

func (m *SYNCData) Reset()                    { *m = SYNCData{} }
func (m *SYNCData) String() string            { return proto.CompactTextString(m) }
func (*SYNCData) ProtoMessage()               {}
func (*SYNCData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SYNCData) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*LegacyMessage)(nil), "qrl.LegacyMessage")
	proto.RegisterType((*NoData)(nil), "qrl.NoData")
	proto.RegisterType((*VEData)(nil), "qrl.VEData")
	proto.RegisterType((*PLData)(nil), "qrl.PLData")
	proto.RegisterType((*PONGData)(nil), "qrl.PONGData")
	proto.RegisterType((*MRData)(nil), "qrl.MRData")
	proto.RegisterType((*BKData)(nil), "qrl.BKData")
	proto.RegisterType((*FBData)(nil), "qrl.FBData")
	proto.RegisterType((*PBData)(nil), "qrl.PBData")
	proto.RegisterType((*SYNCData)(nil), "qrl.SYNCData")
	proto.RegisterEnum("qrl.LegacyMessage_FuncName", LegacyMessage_FuncName_name, LegacyMessage_FuncName_value)
}

func init() { proto.RegisterFile("qrllegacy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 904 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0xdf, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xe3, 0xc6, 0x71, 0x9c, 0x93, 0x7f, 0xd3, 0x69, 0x11, 0x01, 0xc4, 0x12, 0x8c, 0x56,
	0xac, 0x8a, 0x54, 0xa4, 0x70, 0x03, 0x48, 0x5c, 0x24, 0xad, 0x8b, 0x57, 0x4d, 0xb3, 0x91, 0x13,
	0xad, 0xe0, 0xca, 0x72, 0x9c, 0x69, 0x62, 0xc5, 0x1e, 0xbb, 0x63, 0x37, 0x6c, 0xdf, 0x8a, 0xb7,
	0xe0, 0x39, 0x78, 0x13, 0x34, 0x67, 0x6c, 0x6f, 0xb6, 0x28, 0xbb, 0x57, 0x93, 0x39, 0xe7, 0x77,
	0x3e, 0x7b, 0x8e, 0xcf, 0x7c, 0x81, 0xfe, 0x83, 0x88, 0x22, 0xb6, 0xf1, 0x83, 0xa7, 0xcb, 0x54,
	0x24, 0x79, 0x42, 0xeb, 0x0f, 0x22, 0xfa, 0xb2, 0xf5, 0x20, 0x22, 0xb5, 0xb7, 0xfe, 0x6e, 0x41,
	0x77, 0x8a, 0xc0, 0x1d, 0xcb, 0x32, 0x7f, 0xc3, 0xe8, 0xcf, 0xd0, 0xba, 0x7f, 0xe4, 0x81, 0xc7,
	0xfd, 0x98, 0x0d, 0xb4, 0xa1, 0xf6, 0xaa, 0x37, 0xfa, 0xea, 0x52, 0x16, 0x7c, 0x80, 0x5d, 0xde,
	0x3c, 0xf2, 0x60, 0xe6, 0xc7, 0xcc, 0x35, 0xef, 0x8b, 0x5f, 0xf4, 0x25, 0x18, 0x3c, 0xb9, 0xf6,
	0x73, 0x7f, 0x70, 0x32, 0xd4, 0x5e, 0xb5, 0x47, 0x6d, 0x2c, 0x9b, 0x61, 0xc8, 0xa9, 0xb9, 0x45,
	0x52, 0x62, 0x7b, 0x86, 0x58, 0xfd, 0x00, 0x7b, 0x6b, 0x97, 0x98, 0x4a, 0x4a, 0x2c, 0x8d, 0x10,
	0xd3, 0x0f, 0xb0, 0xf9, 0xb4, 0xc4, 0x54, 0x92, 0xfe, 0x00, 0x66, 0x9a, 0xf0, 0x0d, 0x82, 0x0d,
	0x04, 0xbb, 0x0a, 0x7c, 0x33, 0xfb, 0xbd, 0x40, 0x2b, 0x40, 0x6a, 0xc6, 0x02, 0x51, 0xe3, 0x40,
	0xf3, 0xce, 0x2d, 0x35, 0x55, 0x92, 0x5a, 0xd0, 0x58, 0x45, 0x49, 0xb0, 0x1b, 0x34, 0x91, 0x02,
	0xa4, 0x26, 0x32, 0xe2, 0xd4, 0x5c, 0x95, 0x92, 0x52, 0xf7, 0x2b, 0x94, 0x32, 0x0f, 0xa4, 0x6e,
	0x26, 0xa5, 0x94, 0x4a, 0xe2, 0x29, 0x14, 0xd6, 0x3a, 0x3c, 0x45, 0x85, 0xa9, 0x24, 0xbd, 0x04,
	0x63, 0xb5, 0x45, 0x0c, 0x10, 0x3b, 0x3f, 0x78, 0x24, 0x0b, 0x37, 0xdb, 0xbc, 0xe4, 0x15, 0x45,
	0x2f, 0xc0, 0xc8, 0xdf, 0x21, 0xdf, 0x46, 0x9e, 0x20, 0xbf, 0x14, 0x3e, 0xcf, 0xfc, 0x20, 0x0f,
	0x13, 0x2e, 0x59, 0x45, 0x48, 0x36, 0xc6, 0xfa, 0x41, 0xe7, 0x38, 0xab, 0x08, 0xd4, 0xdd, 0x21,
	0xdb, 0xfd, 0x88, 0xee, 0xae, 0x62, 0x95, 0x6e, 0xef, 0x23, 0x6c, 0xa5, 0x1b, 0x29, 0xb6, 0x7f,
	0x9c, 0x8d, 0x2a, 0x36, 0x53, 0x1f, 0x9e, 0x1c, 0x67, 0x15, 0x41, 0x7f, 0x85, 0x26, 0x4b, 0x55,
	0xe3, 0x4e, 0x11, 0x7e, 0x81, 0xb0, 0xcd, 0x03, 0xf1, 0x94, 0xe6, 0x6c, 0x6d, 0xa7, 0x5b, 0x16,
	0x33, 0xe1, 0x47, 0xc5, 0xd8, 0x3a, 0x35, 0xb7, 0x2c, 0x90, 0x93, 0x93, 0x3d, 0xf1, 0x00, 0x8b,
	0xe9, 0xc1, 0xe4, 0x2c, 0xfe, 0x9c, 0x5d, 0x95, 0x93, 0x53, 0x02, 0xf4, 0x37, 0xe8, 0x05, 0x5b,
	0x3f, 0xe4, 0x8b, 0xdc, 0xcf, 0xd5, 0xf0, 0x9e, 0x61, 0xc9, 0x59, 0x31, 0xe3, 0x6b, 0x76, 0x55,
	0xa5, 0x9d, 0x9a, 0xfb, 0x0c, 0x96, 0xe5, 0x3c, 0x59, 0x33, 0x87, 0xf9, 0x6b, 0x26, 0x1c, 0x3f,
	0xdb, 0x0e, 0xce, 0x9f, 0x95, 0xbf, 0x4f, 0xc9, 0xf2, 0x0f, 0x61, 0xfa, 0x0b, 0x40, 0x3a, 0x4a,
	0xc7, 0x81, 0xfa, 0x34, 0x9f, 0x61, 0xe9, 0xe7, 0x6a, 0x92, 0x46, 0xf3, 0x71, 0xb0, 0xe3, 0xc9,
	0x5f, 0x11, 0x5b, 0x6f, 0x58, 0xcc, 0x78, 0xee, 0xd4, 0xdc, 0x03, 0xd8, 0xfa, 0x47, 0x03, 0xb3,
	0xbc, 0xab, 0xd4, 0x80, 0x93, 0xb7, 0x36, 0xa9, 0xc9, 0x75, 0x3e, 0x25, 0x1a, 0x35, 0x41, 0x97,
	0xf7, 0x84, 0x9c, 0xc8, 0xc8, 0x9d, 0x4b, 0xea, 0xb4, 0x09, 0xf5, 0xc5, 0xcd, 0x1d, 0xd1, 0x65,
	0x60, 0x72, 0x4b, 0x1a, 0x72, 0xbd, 0x99, 0x10, 0x03, 0x4b, 0x26, 0xa4, 0x89, 0x71, 0x87, 0x98,
	0x72, 0x5d, 0xfe, 0x41, 0x5a, 0x72, 0x9d, 0x2e, 0x09, 0xc8, 0x42, 0x7b, 0xee, 0x90, 0x36, 0x2a,
	0x2d, 0x49, 0x07, 0x81, 0x5b, 0xd2, 0xc5, 0x75, 0x49, 0x7a, 0x72, 0x5d, 0x4c, 0x49, 0x5f, 0x3e,
	0x53, 0x76, 0x98, 0x10, 0xda, 0x03, 0xb8, 0x72, 0xc6, 0xaf, 0x67, 0x8b, 0xe5, 0x78, 0x69, 0x93,
	0x53, 0x4a, 0xa0, 0xe3, 0xd8, 0xe3, 0x6b, 0xdb, 0x75, 0xc6, 0x0b, 0xc7, 0x5e, 0x10, 0x4a, 0xdb,
	0xd0, 0x9c, 0x8f, 0xe6, 0xde, 0xf8, 0xea, 0x96, 0x9c, 0x4d, 0x0c, 0xd0, 0xd7, 0xf2, 0x44, 0x26,
	0x18, 0xca, 0x53, 0xac, 0x18, 0x0c, 0x65, 0x1b, 0x74, 0x00, 0xcd, 0x3d, 0x13, 0x59, 0x98, 0x70,
	0xb4, 0xac, 0x96, 0x5b, 0x6e, 0xe9, 0x05, 0x9c, 0x6e, 0x18, 0x67, 0x59, 0x98, 0x79, 0xa9, 0x60,
	0x7b, 0x6f, 0x2b, 0x9b, 0x2f, 0xfd, 0xa9, 0xe3, 0xf6, 0x8b, 0xc4, 0x5c, 0xb0, 0x3d, 0xb6, 0xf9,
	0x6b, 0x00, 0xe1, 0xe7, 0xcc, 0x8b, 0xc2, 0x38, 0xcc, 0xd1, 0x9d, 0x74, 0xb7, 0x25, 0x23, 0x53,
	0x19, 0xb0, 0xae, 0xc1, 0x50, 0xf6, 0x43, 0xbf, 0x00, 0x33, 0x65, 0x4c, 0x78, 0x61, 0x9a, 0x0d,
	0xb4, 0x61, 0x5d, 0x3e, 0x4f, 0xee, 0x5f, 0xa7, 0x19, 0xfd, 0x06, 0xda, 0xe9, 0xe3, 0x2a, 0x0a,
	0x03, 0x2f, 0x4d, 0x44, 0x8e, 0x4f, 0xea, 0xba, 0xa0, 0x42, 0xf3, 0x44, 0xe4, 0x16, 0x80, 0x59,
	0x7a, 0x93, 0xf5, 0xaf, 0x06, 0x86, 0x72, 0x1f, 0x4a, 0x41, 0xc7, 0x57, 0xd3, 0xf0, 0xd5, 0xf0,
	0x37, 0xfd, 0x11, 0xf4, 0xfc, 0x29, 0x65, 0x28, 0xf2, 0x09, 0x17, 0x46, 0x90, 0xbe, 0x84, 0x5e,
	0x96, 0xfb, 0x3b, 0xe6, 0x65, 0x2c, 0x62, 0x41, 0x9e, 0x08, 0x3c, 0x44, 0xc7, 0xed, 0x62, 0x74,
	0x51, 0x04, 0xe9, 0xb7, 0xd0, 0x41, 0x13, 0xf3, 0xf8, 0x63, 0xbc, 0x62, 0x02, 0x0d, 0x56, 0x77,
	0xdb, 0x18, 0x9b, 0x61, 0x88, 0x7e, 0x0f, 0x7d, 0xd5, 0x2e, 0x1c, 0x42, 0x7c, 0xb3, 0x06, 0x4a,
	0xf5, 0x64, 0xd8, 0xa9, 0xa2, 0xf2, 0xbc, 0x82, 0xed, 0x99, 0x1f, 0xa9, 0xce, 0x1a, 0x08, 0x81,
	0x0a, 0xc9, 0xa6, 0x5a, 0x6f, 0xc0, 0x98, 0xdc, 0xe2, 0x11, 0xbf, 0xab, 0xdc, 0x57, 0xfb, 0x9f,
	0xfb, 0x56, 0xde, 0x3b, 0x2c, 0xbd, 0xf7, 0xe4, 0xb9, 0xf7, 0x16, 0xce, 0x6b, 0xbd, 0x00, 0x43,
	0xd9, 0x2c, 0x3d, 0x87, 0x46, 0xc8, 0xd7, 0xec, 0x1d, 0xea, 0xe9, 0xae, 0xda, 0x58, 0x17, 0x60,
	0x28, 0x7f, 0x7d, 0xaf, 0xa5, 0x1d, 0xd3, 0x1a, 0x82, 0x59, 0x5e, 0x77, 0xa9, 0x96, 0xc9, 0x0b,
	0x5b, 0x4c, 0x90, 0xda, 0xac, 0x0c, 0xfc, 0x9f, 0xfc, 0xe9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x53, 0x31, 0x0e, 0x22, 0x4a, 0x07, 0x00, 0x00,
}
