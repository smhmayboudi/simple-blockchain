package main

import (
	"time"
)

// Block is a structure to store each point of data.
type Block struct {
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
	Timestamp     int64
	// Nonce is an index.
	Nonce int
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Data:          []byte(data),
		Hash:          []byte{},
		PrevBlockHash: prevBlockHash,
		Timestamp:     time.Now().Unix(),
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
