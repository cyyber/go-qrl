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

package snap

import (
	"bytes"
	"errors"
	"testing"

	"github.com/theQRL/go-qrl/p2p"
)

// Tests that a message exceeding the size limit is rejected, and that its
// payload is discarded nonetheless.
func TestOversizedMessageDiscarded(t *testing.T) {
	t.Parallel()

	app, net := p2p.MsgPipe()
	defer app.Close()

	peer := NewFakePeer(SNAP2, "oversized", net)

	payload := bytes.NewReader(make([]byte, maxMessageSize+1))
	go app.WriteMsg(p2p.Msg{Code: GetAccountRangeMsg, Size: uint32(payload.Len()), Payload: payload})

	if err := HandleMessage(&dummyBackend{}, peer); !errors.Is(err, errMsgTooLarge) {
		t.Fatalf("error mismatch: have %v, want %v", err, errMsgTooLarge)
	}
	if unread := payload.Len(); unread != 0 {
		t.Errorf("message payload not discarded: %d bytes unread", unread)
	}
}
