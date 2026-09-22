package engine

import (
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	"github.com/theQRL/go-qrl/core/types"
)

func TestExecutableDataToBlockNilWithdrawals(t *testing.T) {
	block, err := ExecutableDataToBlockNoHash(ExecutableData{
		LogsBloom:     make([]byte, 256),
		BaseFeePerGas: big.NewInt(1),
	})
	if err != nil {
		t.Fatal(err)
	}
	if block.Header().WithdrawalsHash == nil || *block.Header().WithdrawalsHash != types.EmptyWithdrawalsHash {
		t.Fatalf("WithdrawalsHash: got %v, want %x", block.Header().WithdrawalsHash, types.EmptyWithdrawalsHash)
	}
	if block.Withdrawals() == nil || len(block.Withdrawals()) != 0 {
		t.Fatalf("withdrawals: got %v, want empty list", block.Withdrawals())
	}
}

func TestExecutableDataToBlockNilBaseFee(t *testing.T) {
	_, err := ExecutableDataToBlockNoHash(ExecutableData{
		LogsBloom:   make([]byte, 256),
		Withdrawals: []*types.Withdrawal{},
	})
	if err == nil || !strings.Contains(err.Error(), "baseFeePerGas") {
		t.Fatalf("expected missing baseFeePerGas error, got %v", err)
	}
}

func TestExecutableDataJSONRequiresWithdrawals(t *testing.T) {
	var data ExecutableData
	err := json.Unmarshal([]byte(`{
		"parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"feeRecipient": "Q00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"stateRoot": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"receiptsRoot": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"logsBloom": "0x",
		"prevRandao": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"blockNumber": "0x0",
		"gasLimit": "0x0",
		"gasUsed": "0x0",
		"timestamp": "0x0",
		"extraData": "0x",
		"baseFeePerGas": "0x1",
		"blockHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"transactions": []
	}`), &data)
	if err == nil || !strings.Contains(err.Error(), "withdrawals") {
		t.Fatalf("expected missing withdrawals error, got %v", err)
	}
}
