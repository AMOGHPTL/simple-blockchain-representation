package blockchain

import (
	"fmt"
	"strconv"
)

//creating a blockchain which is a struct which has element blocks which is a slice/array of type block

type Blockchain struct {
	Blocks []*Block
}

//creating a block which is of type struct and has 3 elements inside 1.hash(slice of byte) 2. data(slice of byte) 3.previoushash(slice of byte)

type Block struct {
	Hash     []byte
	Data     string
	PrevHash []byte
	Nonce    int
}

// derive function combines data and prevhash and joins them then use sha256 hashing algorithm to create the hash of the current block
// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{[]byte(b.Data), b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

//createBlock function creates a block using hash previous hash and data of the block and returns it
//first step: create a variable and address it to the block and give the block its entries
//second step: apply the derive hash function on this function which adds the current hash to the block

func CreateBlock(data string, prevhash []byte) *Block {
	block := &Block{[]byte{}, data, prevhash, 0}
	// block.DeriveHash()
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

//addBlock function calls the create block function and appends it to the the blockchain
//the previous hash is extracted from -1 block's current hash
//it also prints the information of the block that is being added

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)

	fmt.Printf("hash of previous block:%x\n", new.PrevHash)
	fmt.Printf("data of the block:%s\n", new.Data)
	fmt.Printf("hash of current block:%x\n", new.Hash)
	pow := NewProof(new)
	fmt.Printf("PoW:%s\n", strconv.FormatBool(pow.Validate()))
	fmt.Printf("nonce:%v\n", new.Nonce)
	fmt.Println()

}

//genisis functioncreates the genisis block which doesn't have a previous hash

func Genesis() *Block {
	new := CreateBlock("Genisis", []byte{})
	fmt.Printf("---------------------BLOCK:GENISIS-------------------------------\n")
	fmt.Printf("hash of previous block:%x\n", new.PrevHash)
	fmt.Printf("data of the block:%s\n", new.Data)
	fmt.Printf("hash of current block:%x\n", new.Hash)
	return new
}

//initial function initialises or starts the block chain with genisis block as the parent block

func Initial() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
