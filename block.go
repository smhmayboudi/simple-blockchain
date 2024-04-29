package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block is a structure to store each point of data.
type Block struct {
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
	Timestamp     int64
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Data:          []byte(data),
		Hash:          []byte{},
		PrevBlockHash: []byte{},
		Timestamp:     time.Now().Unix(),
	}
	block.SetHash()

	return block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
