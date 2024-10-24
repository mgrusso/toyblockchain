package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Cryptoblock struct {
	blocknum int
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (c *Cryptoblock) BuildHash() {
	details := bytes.Join([][]byte{c.Data, c.PrevHash}, []byte{})
	hash := sha256.Sum256(details)
	c.Hash = hash[:]
}

func BuildBlock(blocknum int, data string, prevHash []byte) *Cryptoblock {
	block := &Cryptoblock{blocknum, []byte{}, []byte(data), prevHash}
	block.BuildHash()
	return block
}

type BlockChain struct {
	blocks []*Cryptoblock
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := BuildBlock(len(chain.blocks), data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Cryptoblock {
	return BuildBlock(0, "GENESIS BLOCK - A blockchain is born", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Cryptoblock{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	fmt.Println("Fillling blockchain...")
	for i := 0; i < 1000; i++ {
		chain.AddBlock(fmt.Sprintf("I'm data of block %d", i+1))
	}

	//chain.AddBlock("First Block after genesis!")
	//chain.AddBlock("Second Block after genesis!")
	//chain.AddBlock("Third Block after genesis!")

	for _, block := range chain.blocks {
		fmt.Printf("Block number %d\n", block.blocknum)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("*****")
	}

}
