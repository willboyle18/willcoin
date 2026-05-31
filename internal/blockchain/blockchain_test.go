package blockchain

import (
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	NewBlockchain()
	blockchain := getBlockchain()
	if len(blockchain.Blocks) != 1 {
		t.Fatalf("expected 1 block, got %d", len(blockchain.Blocks))
	}
}
