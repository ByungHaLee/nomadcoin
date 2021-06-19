package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	blocks := GetBlockchain().blocks
	if len(blocks) == 0 {
		return ""
	}
	return blocks[len(blocks)-1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}

func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}
