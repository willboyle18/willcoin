package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index int `json:"index"`
	Timestamp int64 `json:"timestamp"`
	Data string `json:"data"`
	PrevHash string `json:"prev_hash"`
	Hash string `json:"hash"`
}

func calculateHash(index int, timestamp int64, data string, prevHash string) string {
	record := fmt.Sprintf("%d|%d|%s|%s", index, timestamp, data, prevHash)
	hash := sha256.Sum256([]byte(record))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func NewGenesisBlock() Block {
	return NewBlock(0, "Genesis Block", "")
}

func NewBlock(index int, data string, prevHash string) Block {
	timestamp := time.Now().Unix()
	hash := calculateHash(index, timestamp, data, prevHash)

	block := Block{index, timestamp, data, string(prevHash), string(hash)}
	return block
}