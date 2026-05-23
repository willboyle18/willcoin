package blockchain

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Blockchain struct {
	Blocks []Block
}

func NewBlockchain() {
	genesisBlock := NewGenesisBlock()

	err := os.Mkdir("data", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	_, err = os.Create("data/blockchain.json")
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(genesisBlock)

	os.WriteFile("data/blockchain.json", data, 0666)
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

