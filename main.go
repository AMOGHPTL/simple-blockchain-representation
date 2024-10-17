package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

type Blockchain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) deriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func createBlock(data string, prevhash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevhash}
	block.deriveHash()
	return block
}

func (chain *Blockchain) addBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)

	fmt.Printf("hash of previous block:%x\n", new.PrevHash)
	fmt.Printf("data of the block:%s\n", new.Data)
	fmt.Printf("hash of current block:%x\n", new.Hash)
}

func Genesis() *Block {
	new := createBlock("Genisis", []byte{})
	fmt.Printf("---------------------BLOCK:GENISIS-------------------------------\n")
	fmt.Printf("hash of previous block:%x\n", new.PrevHash)
	fmt.Printf("data of the block:%s\n", new.Data)
	fmt.Printf("hash of current block:%x\n", new.Hash)
	return new
}

func initial() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := initial()
	time.Sleep(time.Second * 2)

	blockHeight := 1
	for {
		fmt.Printf("---------------------BLOCK:%v-------------------------------\n", blockHeight)
		msg := fmt.Sprintf("block %v added", blockHeight)
		chain.addBlock(msg)
		time.Sleep(time.Second * 2)
		blockHeight++
	}
}
