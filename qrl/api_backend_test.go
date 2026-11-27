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
	"github.com/theQRL/go-qrl/internal/qrlapi"
	"github.com/theQRL/go-qrl/params"
	"github.com/theQRL/go-qrl/rpc"
	"github.com/theQRL/go-qrl/trie/triedb/database"
)

// newTestAPIBackend returns a backend over a fresh chain holding only the
// genesis block, stored with the given state scheme.
func newTestAPIBackend(t *testing.T, scheme string) *QRLAPIBackend {
	t.Helper()
	db := rawdb.NewMemoryDatabase()
	genesis := &core.Genesis{Config: params.TestChainConfig}
	chain, err := core.NewBlockChain(db, core.DefaultCacheConfigWithScheme(scheme), genesis, beacon.NewFaker(), vm.Config{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(chain.Stop)
	return &QRLAPIBackend{qrl: &QRL{blockchain: chain, chainDb: db}}
}

func TestBlockTagNotFound(t *testing.T) {
	b := newTestAPIBackend(t, rawdb.HashScheme)
	for number, want := range map[rpc.BlockNumber]error{
		rpc.SafeBlockNumber:      qrlapi.ErrSafeBlockNotFound,
		rpc.FinalizedBlockNumber: qrlapi.ErrFinalizedBlockNotFound,
	} {
		if _, err := b.HeaderByNumber(context.Background(), number); !errors.Is(err, want) {
			t.Errorf("header %v: have %v, want %v", number, err, want)
		}
		if _, err := b.BlockByNumber(context.Background(), number); !errors.Is(err, want) {
			t.Errorf("block %v: have %v, want %v", number, err, want)
		}
	}
}

func TestStateAtBlockUnavailable(t *testing.T) {
	for _, scheme := range []string{rawdb.HashScheme, rawdb.PathScheme} {
		b := newTestAPIBackend(t, scheme)

		// Neither the state of this block nor its parent is known.
		block := types.NewBlockWithHeader(&types.Header{Number: big.NewInt(1), ParentHash: common.Hash{1}, Root: common.Hash512{2}})
		if _, _, err := b.StateAtBlock(context.Background(), block, 128, nil, true, false); !errors.Is(err, database.ErrStateUnavailable) {
			t.Errorf("%s: have %v, want %v", scheme, err, database.ErrStateUnavailable)
		}
	}
}
