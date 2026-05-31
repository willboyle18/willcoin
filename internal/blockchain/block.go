package blockchain

import (
	"time"
)

type Block struct {
	Index     int    `json:"index"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	Nonce     int64  `json:"nonce"`
	PrevHash  string `json:"prev_hash"`
	Hash      string `json:"hash"`
}

func NewGenesisBlock() Block {
	return NewBlock(0, "Genesis Block", 0, "")
}

func NewBlock(index int, data string, nonce int64, prevHash string) Block {
	timestamp := time.Now().Unix()
	hash := calculateHash(index, timestamp, data, nonce, prevHash)

	block := Block{index, timestamp, data, nonce, string(prevHash), string(hash)}
	return block
}
