package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index int
	Timestamp int64
	Data string
	PrevHash [32]byte
	Hash [32]byte
}

func calculateHash(index int, timestamp int64, data string, prevHash [32]byte) [32]byte {
	stringPrevHash := hex.EncodeToString(prevHash[:])
	record := fmt.Sprintf("%d|%d|%s|%s", index, timestamp, data, stringPrevHash)
	hash := sha256.Sum256([]byte(record))
	return hash
}

func NewGenesisBlock() Block {
	return NewBlock(0, "Genesis Block", [32]byte{})
}

func NewBlock(index int, data string, prevHash [32]byte) Block {
	timestamp := time.Now().Unix()
	hash := calculateHash(index, timestamp, data, prevHash)

	block := Block{index, timestamp, data, prevHash, hash}
	return block
}