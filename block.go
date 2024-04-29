package main

import (
	"bytes"
	"encoding/gob"
	"log"
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

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}
