package blockchain

import (
	"fmt"
	"log"
	"os"
	"time"
	"math/big"
)

type Blockchain struct {
	Blocks []Block `json:"blocks"`
}

func NewBlockchain() {
	genesisBlock := NewGenesisBlock()
	blocks := []Block{genesisBlock}
	blockchain := Blockchain{blocks}

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
	blocks := blockchain.Blocks

	index := len(blocks)
	lastBlock := blocks[index-1]
	prevHash := lastBlock.Hash

	newBlock := NewBlock(index, data, 0, prevHash)
	blocks = append(blocks, newBlock)
	blockchain.Blocks = blocks

	writeBlockchain(blockchain)
}

func PrintBlockchain() {
	blockchain := getBlockchain()

	for _, block := range blockchain.Blocks {
		fmt.Println("Index:", block.Index)
		fmt.Println("Timestamp:", block.Timestamp)
		fmt.Println("Data:", block.Data)
		fmt.Println("PrevHash:", block.PrevHash)
		fmt.Println("Hash:", block.Hash)
		fmt.Println()
	}
}

func MineBlock(data string) {
	index := getNextIndex()
	prevHash := getLastHash()
	var nonce int64 = 0

	targetHashStr := "0000000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	targetHash := new(big.Int)
	targetHash.SetString(targetHashStr, 16)

	for true {
		timestamp := time.Now().Unix()
		hash := calculateHash(index, timestamp, data, nonce, prevHash)

		hashStr := hash
		convertedHash := new(big.Int)
		convertedHash.SetString(hashStr, 16)

		if convertedHash.Cmp(targetHash) < 0 {
			fmt.Println("Mined with nonce value of", nonce)
			minedBlock := Block{index, timestamp, data, nonce, hash, prevHash}
			blockchain := getBlockchain()
			blockchain.Blocks = append(blockchain.Blocks, minedBlock)
			writeBlockchain(blockchain)
			break
		}
		nonce++
	}
}
