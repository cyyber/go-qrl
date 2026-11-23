// Copyright 2025 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package txpool

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/theQRL/go-qrl/common"
	"github.com/theQRL/go-qrl/core"
	"github.com/theQRL/go-qrl/core/types"
	"github.com/theQRL/go-qrl/crypto/pqcrypto/wallet"
	"github.com/theQRL/go-qrl/params"
)

func TestValidateTransactionEIP2681(t *testing.T) {
	w, err := wallet.Generate(wallet.ML_DSA_87)
	if err != nil {
		t.Fatal(err)
	}

	head := &types.Header{
		Number:   big.NewInt(1),
		GasLimit: 5000000,
		Time:     1,
	}

	signer := types.NewZondSigner(params.TestChainConfig.ChainID)

	// Create validation options
	opts := &ValidationOptions{
		Config:  params.TestChainConfig,
		Accept:  0xFF, // Accept all transaction types
		MaxSize: 32 * 1024,
		MinTip:  big.NewInt(0),
	}

	tests := []struct {
		name    string
		nonce   uint64
		wantErr error
	}{
		{
			name:    "normal nonce",
			nonce:   42,
			wantErr: nil,
		},
		{
			name:    "max allowed nonce (2^64-2)",
			nonce:   math.MaxUint64 - 1,
			wantErr: nil,
		},
		{
			name:    "EIP-2681 nonce overflow (2^64-1)",
			nonce:   math.MaxUint64,
			wantErr: core.ErrNonceMax,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := createTestTransaction(w, signer, tt.nonce)
			err := ValidateTransaction(tx, head, signer, opts)

			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("ValidateTransaction() error = %v, wantErr nil", err)
				}
			} else {
				if err == nil {
					t.Errorf("ValidateTransaction() error = nil, wantErr %v", tt.wantErr)
				} else if !errors.Is(err, tt.wantErr) {
					t.Errorf("ValidateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

// createTestTransaction creates a basic transaction for testing
func createTestTransaction(w wallet.Wallet, signer types.Signer, nonce uint64) *types.Transaction {
	to := common.BytesToAddress([]byte{0x01})

	txdata := &types.DynamicFeeTx{
		ChainID:   params.TestChainConfig.ChainID,
		Nonce:     nonce,
		To:        &to,
		Value:     big.NewInt(1000),
		Gas:       21000,
		GasTipCap: big.NewInt(1),
		GasFeeCap: big.NewInt(1),
		Data:      nil,
	}

	signedTx, _ := types.SignNewTx(w, signer, txdata)
	return signedTx
}
