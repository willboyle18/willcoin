package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func writeBlockchain(blockchain Blockchain) {
	data, err := json.Marshal(blockchain)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")

	os.WriteFile("data/blockchain.json", out.Bytes(), 0666)
}

func getBlockchain() Blockchain {
	blockchainBytes, err := os.ReadFile("data/blockchain.json")
	if err != nil {
		log.Fatal(err)
	}

	var blockchain Blockchain

	err = json.Unmarshal(blockchainBytes, &blockchain)
	if err != nil {
		log.Fatal(err)
	}

	return blockchain
}

func getLastHash() string {
	blockchain := getBlockchain()
	return blockchain.Blocks[len(blockchain.Blocks)-1].Hash
}

func getNextIndex() int {
	blockchain := getBlockchain()
	return len(blockchain.Blocks)
}

func calculateHash(index int, timestamp int64, data string, nonce int64, prevHash string) string {
	record := fmt.Sprintf("%d|%d|%s|%d|%s", index, timestamp, data, nonce, prevHash)
	hash := sha256.Sum256([]byte(record))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
