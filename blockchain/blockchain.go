package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []*block // slice of pointers
}

// singletone !
var b *blockchain  // private as lowercased
var once sync.Once // see sync pkg doc

func getPrevHash() string {
	totalBlocksLen := len(GetBlockChain().blocks)
	if totalBlocksLen == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlocksLen-1].hash
}

// use receiver func
func (b *block) setHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	hexHash := fmt.Sprintf("%x", hash) // return in hex string
	b.hash = hexHash
}

func createBlock(data string) *block {
	newBlock := block{data: data, hash: "", prevHash: getPrevHash()}
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

func (b *blockchain) GetAllBlocks() []*block {
	return b.blocks
}
