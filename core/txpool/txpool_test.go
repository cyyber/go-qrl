package txpool

import (
	"math/big"
	"testing"

	"github.com/theQRL/go-qrl/core"
	"github.com/theQRL/go-qrl/core/types"
	"github.com/theQRL/go-qrl/event"
)

type testChain struct {
	head *types.Header
	feed event.Feed
}

func (c *testChain) CurrentBlock() *types.Header { return c.head }

func (c *testChain) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return c.feed.Subscribe(ch)
}

// TestNewSubscribesBeforeReturning checks that a chain head change right after
// New returns reaches the pool, instead of being lost before the pool's loop
// has started.
func TestNewSubscribesBeforeReturning(t *testing.T) {
	chain := &testChain{head: &types.Header{Number: big.NewInt(0)}}
	pool, err := New(new(big.Int), chain, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	block := types.NewBlockWithHeader(&types.Header{Number: big.NewInt(1)})
	if n := chain.feed.Send(core.ChainHeadEvent{Block: block}); n != 1 {
		t.Fatalf("head event delivered to %d subscribers, want 1", n)
	}
}
