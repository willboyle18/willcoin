package blockchain

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
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
	blockchainBytes, err := os.ReadFile("data/blockchain.json")
	if err != nil {
		log.Fatal(err)
	}

	var blockchain []Block

	err = json.Unmarshal(blockchainBytes, &blockchain)
	if err != nil {
		log.Fatal(err)
	}

	index := len(blockchain)
	lastBlock := blockchain[index - 1]
	prevHash := lastBlock.Hash

	newBlock := NewBlock(index, data, prevHash)
	blockchain = append(blockchain, newBlock)

	writeBlockchain(blockchain)
}

func writeBlockchain(blockchain []Block) {
	data, err := json.Marshal(blockchain)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")

	os.WriteFile("data/blockchain.json", out.Bytes(), 0666)
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

