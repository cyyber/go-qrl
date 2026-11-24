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

package qrl

import (
	"errors"
	"testing"
	"time"

	"github.com/theQRL/go-qrl/p2p"
	"github.com/theQRL/go-qrl/p2p/qnode"
)

// Tests that the delivery of a response is aborted when the peer disconnects,
// instead of waiting forever for the requester to pick the response up.
func TestDispatchResponseDisconnect(t *testing.T) {
	t.Parallel()

	backend := newTestBackend(0)
	defer backend.close()

	app, net := p2p.MsgPipe()
	defer app.Close()

	peer := NewPeer(QRL2, p2p.NewPeer(qnode.ID{1}, "peer", nil), net, backend.TxPool())

	// Dispatch a request whose response is never picked up. The request goes
	// out on the wire, so it has to be read on the remote side.
	go func() {
		if msg, err := app.ReadMsg(); err == nil {
			msg.Discard()
		}
	}()
	req := &Request{
		id:   1,
		sink: make(chan *Response),
		code: GetBlockHeadersMsg,
		want: BlockHeadersMsg,
		data: &GetBlockHeadersPacket{
			RequestId:              1,
			GetBlockHeadersRequest: &GetBlockHeadersRequest{Amount: 1},
		},
	}
	if err := peer.dispatchRequest(req); err != nil {
		t.Fatalf("failed to dispatch request: %v", err)
	}

	// Deliver the response. The metadata callback runs once the dispatcher has
	// accepted the response, right before it is handed over to the requester.
	var (
		accepted = make(chan struct{})
		done     = make(chan error, 1)
	)
	go func() {
		res := &Response{id: 1, code: BlockHeadersMsg, Res: new(BlockHeadersRequest)}
		done <- peer.dispatchResponse(res, func() any {
			close(accepted)
			return nil
		})
	}()
	select {
	case <-accepted:
	case err := <-done:
		t.Fatalf("response not accepted: %v", err)
	case <-time.After(5 * time.Second):
		t.Fatal("response not accepted within 5 seconds")
	}

	// Nobody is reading the response, so the delivery can only end by the peer
	// going away.
	peer.Close()
	select {
	case err := <-done:
		if !errors.Is(err, errDisconnected) {
			t.Fatalf("error mismatch: have %v, want %v", err, errDisconnected)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("response delivery still blocked after the peer disconnected")
	}
}
