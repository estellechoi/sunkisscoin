package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*Block // slice of pointers
}

// singletone !
var b *blockchain  // private as lowercased
var once sync.Once // see sync pkg doc

func getPrevHash() string {
	totalBlocksLen := len(GetBlockChain().blocks)
	if totalBlocksLen == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlocksLen-1].Hash
}

// use receiver func
func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	hexHash := fmt.Sprintf("%x", hash) // return in hex string
	b.Hash = hexHash
}

func createBlock(data string) *Block {
	newBlock := Block{Data: data, Hash: "", PrevHash: getPrevHash()}
	newBlock.setHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockChain() *blockchain {
	if b == nil {
		// this makes sure that this line is executed only once although multiple goroutines are triggered
		once.Do(func() {
			b = &blockchain{} // get address of the new blockchain initialized
			b.AddBlock("Genesis Block")
		})
	}
	return b
}

func (b *blockchain) GetAllBlocks() []*Block {
	return b.blocks
}
