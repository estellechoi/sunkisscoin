package main

import (
	"fmt"

	"github.com/estellechoi/sunkisscoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	for _, block := range chain.GetAllBlocks() {
		fmt.Printf("%s\n", block)
	}
}
