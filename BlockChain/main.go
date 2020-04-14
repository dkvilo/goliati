package blockchain

import (
	block "github.com/dkvilo/goliati/Block"
)

// Chain structure
type Chain struct {
	Ledger []block.Block
}

// New Adds genesis block to chain
func (ch Chain) New() Chain {
	b := block.Block{}
	ch.Ledger = append(ch.Ledger, b.CreateGenesis())
	return ch
}

// AddBlock adds block to blockchaing ledger
func (ch *Chain) AddBlock(block block.Block) []block.Block {
	block.PreviousHash = ch.GetLastBlock().Hash
	ch.Ledger = append(ch.Ledger, block)
	return ch.Ledger
}

// GetLastBlock returns last block from ledger
func (ch *Chain) GetLastBlock() block.Block {
	return ch.Ledger[len(ch.Ledger)-1]
}
