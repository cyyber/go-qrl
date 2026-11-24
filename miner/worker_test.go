// Copyright 2026 The go-qrl Authors
// This file is part of the go-qrl library.
//
// The go-qrl library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-qrl library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-qrl library. If not, see <http://www.gnu.org/licenses/>.

package miner

import (
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/theQRL/go-qrl/consensus/beacon"
	"github.com/theQRL/go-qrl/core"
	"github.com/theQRL/go-qrl/core/rawdb"
	"github.com/theQRL/go-qrl/core/txpool"
	"github.com/theQRL/go-qrl/core/txpool/legacypool"
	"github.com/theQRL/go-qrl/core/types"
	"github.com/theQRL/go-qrl/core/vm"
	"github.com/theQRL/go-qrl/crypto"
	"github.com/theQRL/go-qrl/params"
	"github.com/theQRL/go-qrl/qrldb"
)

// Tests that the block builder stops adding transactions before the block
// reaches the size it is allowed to build.
func TestTxFitsSize(t *testing.T) {
	tx := types.NewTx(&types.DynamicFeeTx{Data: make([]byte, 1024)})

	env := &environment{size: maxBuildBlockSize - tx.Size() - 1}
	if !env.txFitsSize(tx) {
		t.Fatal("transaction below the size limit does not fit")
	}
	env.size++
	if env.txFitsSize(tx) {
		t.Fatal("transaction reaching the size limit fits")
	}
}

// Tests that no block is produced if the state database reported an error
// while the block was being built.
func TestGenerateWorkStateReadFailure(t *testing.T) {
	var (
		code   = []byte{byte(vm.STOP)}
		engine = beacon.NewFaker()
	)
	// The pending transaction calls a contract, so building the block has to
	// load the contract code from the database.
	setup := func() (*Miner, qrldb.Database) {
		db := rawdb.NewMemoryDatabase()
		gspec := &core.Genesis{
			Config: params.TestChainConfig,
			Alloc: core.GenesisAlloc{
				testBankAddress: {Balance: testBankFunds},
				testUserAddress: {Code: code},
			},
		}
		chain, err := core.NewBlockChain(db, &core.CacheConfig{TrieDirtyDisabled: true}, gspec, engine, vm.Config{}, nil)
		if err != nil {
			t.Fatalf("core.NewBlockChain failed: %v", err)
		}
		t.Cleanup(chain.Stop)

		pool := legacypool.New(testTxPoolConfig, chain)
		txPool, err := txpool.New(new(big.Int).SetUint64(testTxPoolConfig.PriceLimit), chain, []txpool.SubPool{pool})
		if err != nil {
			t.Fatalf("txpool.New failed: %v", err)
		}
		t.Cleanup(func() { txPool.Close() })
		for i, err := range txPool.Add(pendingTxs, true, true) {
			if err != nil {
				t.Fatalf("failed to add transaction %d: %v", i, err)
			}
		}
		backend := &testWorkerBackend{db: db, chain: chain, txPool: txPool, genesis: gspec}
		return New(backend, testConfig, engine), db
	}
	build := func(w *Miner) *newPayloadResult {
		return w.generateWork(&generateParams{
			timestamp: uint64(time.Now().Unix()),
			forceTime: true,
			coinbase:  testBankAddress,
		})
	}

	// With the code available, the block is built and includes the call.
	w, _ := setup()
	res := build(w)
	if res.err != nil {
		t.Fatalf("failed to build block: %v", res.err)
	}
	if have, want := len(res.block.Transactions()), len(pendingTxs); have != want {
		t.Fatalf("transaction count mismatch: have %d, want %d", have, want)
	}

	// With the code missing, the state read fails and no block is built.
	w, db := setup()
	codeHash := crypto.Keccak512Hash(code)
	if !rawdb.HasCode(db, codeHash) {
		t.Fatal("contract code not found in database")
	}
	rawdb.DeleteCode(db, codeHash)

	res = build(w)
	if !errors.Is(res.err, errStateReadFailure) {
		t.Fatalf("error mismatch: have %v, want %v", res.err, errStateReadFailure)
	}
	if res.block != nil {
		t.Fatal("block built despite the state read failure")
	}
}
