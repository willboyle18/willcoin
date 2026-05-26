package blockchain

import (
	"fmt"
	"log"
	"os"
)

type Blockchain struct {
	Blocks []Block `json:"blocks"`
}

func NewBlockchain() {
	genesisBlock := NewGenesisBlock()
	blockchain := []Block{genesisBlock}

	err := os.MkdirAll("data", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	_, err = os.Create("data/blockchain.json")
	if err != nil {
		log.Fatal(err)
	}

	writeBlockchain(blockchain)
}


func AddBlock(data string) {
	blockchain := getBlockchain()

	index := len(blockchain)
	lastBlock := blockchain[index - 1]
	prevHash := lastBlock.Hash

	newBlock := NewBlock(index, data, prevHash)
	blockchain = append(blockchain, newBlock)

	writeBlockchain(blockchain)
}

func PrintBlockchain() {
	blockchain := getBlockchain()

	for _, block := range blockchain {
		fmt.Println("Index:", block.Index)
		fmt.Println("Timestamp:", block.Timestamp)
		fmt.Println("Data:", block.Data)
		fmt.Println("PrevHash:", block.PrevHash)
		fmt.Println("Hash:", block.Hash)
		fmt.Println()
	}
}

