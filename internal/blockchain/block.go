package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hex"
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
	record := fmt.Sprintf("%d%d%s%s", index, timestamp, data, stringPrevHash)
	hash := sha256.Sum256([]byte(record))
	return hash
}

func NewGenesisBlock() {
	genesisData := "Genesis Block"
	timestamp := time.Now().Unix()
	index := 0
	seedPrevHash := [32]byte{}
	hash := calculateHash(index, timestamp, genesisData, seedPrevHash)

	genesisBlock := Block{index, timestamp, genesisData, seedPrevHash, hash}
}