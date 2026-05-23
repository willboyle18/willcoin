package blockchain

import (
	"encoding/hex"
	"fmt"
)

type Blockchain struct {
	Blocks []Block
}

func NewBlockchain() Blockchain {
	genesisBlock := NewGenesisBlock()

	return Blockchain{
		Blocks: []Block{genesisBlock}, // Currently in memory, will change when persistence is implemented
	}
}

func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1] // Currently in memory, will change when persistence is implemented

	newBlock := NewBlock(
		previousBlock.Index+1,
		data,
		previousBlock.Hash,
	)

	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc Blockchain) PrintBlockchain() {
	for _, block := range bc.Blocks {
		fmt.Println("Index:", block.Index)
		fmt.Println("Timestamp:", block.Timestamp)
		fmt.Println("Data:", block.Data)
		fmt.Println("PrevHash:", hex.EncodeToString(block.PrevHash[:]))
		fmt.Println("Hash:", hex.EncodeToString(block.Hash[:]))
		fmt.Println()
	}
}

