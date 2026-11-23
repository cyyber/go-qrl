// Copyright 2026 The go-ethereum Authors
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

package qrl

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/theQRL/go-qrl/common"
	"github.com/theQRL/go-qrl/consensus/beacon"
	"github.com/theQRL/go-qrl/core"
	"github.com/theQRL/go-qrl/core/rawdb"
	"github.com/theQRL/go-qrl/core/types"
	"github.com/theQRL/go-qrl/core/vm"
	"github.com/theQRL/go-qrl/params"
)

// newAccessorChain generates blocks with a transfer per transaction and
// imports them into a fresh chain.
func newAccessorChain(t *testing.T, genesis *core.Genesis, txs int) (*QRL, []*types.Block) {
	t.Helper()
	signer := types.NewZondSigner(genesis.Config.ChainID)
	_, blocks, _ := core.GenerateChainWithGenesis(genesis, beacon.NewFaker(), 1, func(i int, b *core.BlockGen) {
		for range txs {
			tx, err := types.SignNewTx(testWallet, signer, &types.DynamicFeeTx{
				ChainID:   genesis.Config.ChainID,
				Nonce:     b.TxNonce(testWallet.GetAddress()),
				To:        &common.Address{1},
				Value:     big.NewInt(1),
				Gas:       params.TxGas,
				GasFeeCap: big.NewInt(params.InitialBaseFee),
			})
			if err != nil {
				t.Fatal(err)
			}
			b.AddTx(tx)
		}
	})
	db := rawdb.NewMemoryDatabase()
	chain, err := core.NewBlockChain(db, nil, genesis, beacon.NewFaker(), vm.Config{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(chain.Stop)
	if _, err := chain.InsertChain(blocks); err != nil {
		t.Fatal(err)
	}
	return &QRL{blockchain: chain, chainDb: db}, blocks
}

// A cancelled request does not replay the transactions before the target.
func TestStateAtTransactionCancelled(t *testing.T) {
	genesis := &core.Genesis{
		Config: params.TestChainConfig,
		Alloc:  core.GenesisAlloc{testWallet.GetAddress(): {Balance: new(big.Int).Mul(big.NewInt(params.Quanta), big.NewInt(1000))}},
	}
	qrl, blocks := newAccessorChain(t, genesis, 2)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, _, _, _, err := qrl.stateAtTransaction(ctx, blocks[0], 1, 0); !errors.Is(err, context.Canceled) {
		t.Fatalf("have %v, want %v", err, context.Canceled)
	}
}
