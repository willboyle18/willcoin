package blockchain

import (
	"bytes"
	"encoding/json"
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
