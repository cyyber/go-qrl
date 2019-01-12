package transactions

import (
	"bytes"
	"encoding/binary"
	"reflect"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	c "github.com/theQRL/go-qrl/pkg/config"
	"github.com/theQRL/go-qrl/pkg/core/addressstate"
	"github.com/theQRL/go-qrl/pkg/crypto"
	"github.com/theQRL/go-qrl/pkg/generated"
	"github.com/theQRL/go-qrl/pkg/log"
	"github.com/theQRL/go-qrl/pkg/misc"
	"github.com/theQRL/qrllib/goqrllib/goqrllib"
)

type PlainTransactionInterface interface {
	TransactionFromPBData(tx *generated.Transaction)
}


type TransactionInterface interface {
	Size() int

	PBData() *generated.Transaction

	SetPBData(*generated.Transaction)

	Type()

	Fee() uint64

	Nonce() uint64

	MasterAddr() []byte

	AddrFrom() []byte

	AddrFromPK() string

	OtsKey() uint64

	GetOtsFromSignature(signature []byte) uint64

	PK() []byte

	Signature() []byte

	FromPBdata(pbdata generated.Transaction) //Set return type

	GetSlave() []byte

	Txhash() []byte

	UpdateTxhash(hashableBytes []byte)

	GetHashableBytes() []byte

	Sign(xmss *crypto.XMSS, message goqrllib.UcharVector)

	ApplyStateChanges(addressesState map[string]*addressstate.AddressState)

	RevertStateChanges(addressesState map[string]*addressstate.AddressState)

	applyStateChangesForPK(addressesState map[string]*addressstate.AddressState)

	revertStateChangesForPK(addressesState map[string]*addressstate.AddressState)

	SetAffectedAddress(addressesState map[string]*addressstate.AddressState)

	validateCustom() bool

	Validate(verifySignature bool) bool

	ValidateSlave(addrFromState *addressstate.AddressState, addrFromPKState *addressstate.AddressState) bool

	ValidateExtended(addrFromState *addressstate.AddressState, addrFromPkState *addressstate.AddressState) bool

	ValidateExtendedCoinbase(blockNumber uint64) bool

	FromJSON(jsonData string) *Transaction

	JSON() (string, error)
}

type Transaction struct {
	log    log.LoggerInterface
	data   *generated.Transaction
	config *c.Config
}

func (tx *Transaction) Size() int {
	return proto.Size(tx.data)
}

func (tx *Transaction) PBData() *generated.Transaction {
	return tx.data
}

func (tx *Transaction) SetPBData(pbData *generated.Transaction) {
	tx.data = pbData
	tx.config = c.GetConfig()
	tx.log = log.GetLogger()
}

func (tx *Transaction) Type() {
	// TODO
	//tx.data.transactionType.(type)
}

func (tx *Transaction) Fee() uint64 {
	return tx.data.Fee
}

func (tx *Transaction) Nonce() uint64 {
	return tx.data.Nonce
}

func (tx *Transaction) SetNonce(n uint64) {
	tx.data.Nonce = n
}

func (tx *Transaction) MasterAddr() []byte {
	return tx.data.MasterAddr
}

func (tx *Transaction) AddrFrom() []byte {
	if tx.MasterAddr() != nil {
		return tx.MasterAddr()
	}

	return misc.UCharVectorToBytes(goqrllib.QRLHelperGetAddress(misc.BytesToUCharVector(tx.PK())))
}

func (tx *Transaction) AddrFromPK() string {
	return misc.PK2Qaddress(tx.PK())
}

func (tx *Transaction) OtsKey() uint64 {
	return uint64(binary.BigEndian.Uint32(tx.data.Signature[0:4]))
}

func (tx *Transaction) GetOtsFromSignature(signature []byte) uint64 {
	return binary.BigEndian.Uint64(signature[0:8])
}

func (tx *Transaction) PK() []byte {
	return tx.data.PublicKey
}

func (tx *Transaction) Signature() []byte {
	return tx.data.Signature
}

func (tx *Transaction) FromPBdata(pbdata generated.Transaction) {
	tx.data = &pbdata
}

func (tx *Transaction) GetSlave() []byte {
	pk := tx.PK()
	upk := misc.NewUCharVector()
	upk.AddBytes(pk)
	upk.New(goqrllib.QRLHelperGetAddress(upk.GetData()))

	if !reflect.DeepEqual(upk.GetBytes(), tx.AddrFrom()) {
		return upk.GetBytes()
	}

	return nil
}

func (tx *Transaction) Txhash() []byte {
	return tx.data.TransactionHash
}

func (tx *Transaction) UpdateTxhash(hashableBytes []byte) {
	tx.data.TransactionHash = tx.GenerateTxHash(hashableBytes)
}

func (tx *Transaction) GetHashableBytes() []byte {
	panic("Not Implemented")
}

func (tx *Transaction) GenerateTxHash(hashableBytes []byte) []byte {
	tmp := new(bytes.Buffer)
	tmp.Write(hashableBytes)
	tmp.Write(tx.Signature())
	tmp.Write(tx.PK())

	return misc.UCharVectorToBytes(goqrllib.Sha2_256(misc.BytesToUCharVector(tmp.Bytes())))
}

func (tx *Transaction) Sign(xmss *crypto.XMSS, message goqrllib.UcharVector) {
	tx.data.Signature = xmss.Sign(message)
	tx.data.TransactionHash = tx.GenerateTxHash(misc.UCharVectorToBytes(message))
}

func (tx *Transaction) applyStateChangesForPK(addressesState map[string]*addressstate.AddressState) {
	addrFromPK := misc.PK2Qaddress(tx.PK())
	if _, ok := addressesState[addrFromPK]; ok {
		//if misc.Bin2Qaddress(tx.AddrFrom()) != addrFromPK {
		//	addressesState[addrFromPK].AppendTransactionHash(tx.Txhash())
		//}
		if tx.OtsKey() >= tx.config.Dev.MaxOTSTracking {
			addressesState[addrFromPK].AppendTransactionHash(tx.Txhash())
		}
		addressesState[addrFromPK].IncreaseNonce()
		addressesState[addrFromPK].SetOTSKey(tx.OtsKey())
	}
}

func (tx *Transaction) revertStateChangesForPK(addressesState map[string]*addressstate.AddressState) {
	addrFromPK := misc.PK2Qaddress(tx.PK())
	if _, ok := addressesState[addrFromPK]; ok {
		//if misc.Bin2Qaddress(tx.AddrFrom()) != addrFromPK {
		//	addressesState[addrFromPK].RemoveTransactionHash(tx.Txhash())
		//}
		if tx.OtsKey() >= tx.config.Dev.MaxOTSTracking {
			addressesState[addrFromPK].RemoveTransactionHash(tx.Txhash())
		}
		addressesState[addrFromPK].DecreaseNonce()
		// Remember to Call UnsetOTSKey
	}
}

func (tx *Transaction) ApplyStateChanges(addressesState map[string]*addressstate.AddressState) {
	panic("Not Implemented")
}

func (tx *Transaction) RevertStateChanges(addressesState map[string]*addressstate.AddressState) {
	panic("Not Implemented")
}

func (tx *Transaction) SetAffectedAddress(addressesState map[string]*addressstate.AddressState) {
	addressesState[misc.Bin2Qaddress(tx.AddrFrom())] = nil
	addressesState[misc.PK2Qaddress(tx.PK())] = nil
}

func (tx *Transaction) validateCustom() bool {
	panic("Not Implemented")
}

func (tx *Transaction) Validate(verifySignature bool) bool {
	panic("Not Implemented")
}

func (tx *Transaction) ValidateSlave(addrFromState *addressstate.AddressState, addrFromPKState *addressstate.AddressState) bool {
	addrFromPK := misc.UCharVectorToBytes(goqrllib.QRLHelperGetAddress(misc.BytesToUCharVector(tx.PK())))

	if reflect.DeepEqual(tx.MasterAddr(), addrFromPK) {
		tx.log.Warn("Matching master_addr field and address from PK")
		return false
	}

	if reflect.DeepEqual(addrFromPK, tx.AddrFrom()) {
		return true
	}

	accessType, ok := addrFromState.GetSlavePermission(tx.PK())

	if !ok {
		tx.log.Warn("Public key and address don't match")
		return false
	}

	if accessType != 0 {
		tx.log.Warn("Access Type ", accessType)
		tx.log.Warn("Slave Address doesnt have sufficient permission")
		return false
	}

	return true
}

func (tx *Transaction) ValidateExtended(addrFromState *addressstate.AddressState, addrFromPkState *addressstate.AddressState) bool {
	panic("Not Implemented")
}

func (tx *Transaction) ValidateExtendedCoinbase(blockNumber uint64) bool {
	panic("Not Implemented")
}

func (tx *Transaction) FromJSON(jsonData string) *Transaction {
	tx.data = &generated.Transaction{}
	jsonpb.UnmarshalString(jsonData, tx.data)
	return tx
}

func (tx *Transaction) JSON() (string, error) {
	ma := jsonpb.Marshaler{}
	return ma.MarshalToString(tx.data)
}

func ProtoToTransaction(protoTX *generated.Transaction) TransactionInterface {
	var tx TransactionInterface
	switch protoTX.TransactionType.(type) {
	case *generated.Transaction_Transfer_:
		tx = &TransferTransaction{}
	case *generated.Transaction_Coinbase:
		tx = &CoinBase{}
	case *generated.Transaction_Token_:
		tx = &TokenTransaction{}
	case *generated.Transaction_TransferToken_:
		tx = &TransferTokenTransaction{}
	case *generated.Transaction_Message_:
		tx = &MessageTransaction{}
	case *generated.Transaction_Slave_:
		tx = &SlaveTransaction{}
	}

	if tx != nil {
		tx.SetPBData(protoTX)
	}

	return tx
}

func ProtoToPlainTransaction(protoTX *generated.Transaction) PlainTransactionInterface {
	var tx PlainTransactionInterface
	switch protoTX.TransactionType.(type) {
	case *generated.Transaction_Transfer_:
		tx = &PlainTransferTransaction{}
	case *generated.Transaction_Coinbase:
		tx = &PlainCoinBaseTransaction{}
	case *generated.Transaction_Token_:
		tx = &PlainTokenTransaction{}
	case *generated.Transaction_TransferToken_:
		tx = &PlainTransferTokenTransaction{}
	case *generated.Transaction_Message_:
		tx = &PlainMessageTransaction{}
	case *generated.Transaction_Slave_:
		tx = &PlainSlaveTransaction{}
	}

	if tx != nil {
		tx.TransactionFromPBData(protoTX)
	}

	return tx
}
