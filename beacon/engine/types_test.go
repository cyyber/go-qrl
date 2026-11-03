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

package engine

import (
	"math/big"
	"testing"

	"github.com/theQRL/go-qrl/core/types"
)

// TestExecutableDataToBlockNoHashBaseFee checks that a missing, negative or
// oversized base fee is rejected with an error rather than a panic.
func TestExecutableDataToBlockNoHashBaseFee(t *testing.T) {
	for _, tc := range []struct {
		name    string
		baseFee *big.Int
		wantErr bool
	}{
		{"nil", nil, true},
		{"negative", big.NewInt(-1), true},
		{"too big", new(big.Int).Lsh(big.NewInt(1), 256), true},
		{"valid", big.NewInt(7), false},
	} {
		data := ExecutableData{
			LogsBloom:     make([]byte, 256),
			BaseFeePerGas: tc.baseFee,
			Withdrawals:   []*types.Withdrawal{},
		}
		_, err := ExecutableDataToBlockNoHash(data)
		if have := err != nil; have != tc.wantErr {
			t.Errorf("%s: have error %v, want error %v", tc.name, err, tc.wantErr)
		}
	}
}
