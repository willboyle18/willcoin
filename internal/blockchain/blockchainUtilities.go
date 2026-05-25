package blockchain

import (
	"encoding/json"
	"bytes"
	"os"
	"log"
)

func writeBlockchain(blockchain []Block) {
	data, err := json.Marshal(blockchain)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")

	os.WriteFile("data/blockchain.json", out.Bytes(), 0666)
}


func getBlockchain() []Block {
	blockchainBytes, err := os.ReadFile("data/blockchain.json")
	if err != nil {
		log.Fatal(err)
	}

	var blockchain []Block

	err = json.Unmarshal(blockchainBytes, &blockchain)
	if err != nil {
		log.Fatal(err)
	}

	return blockchain
}