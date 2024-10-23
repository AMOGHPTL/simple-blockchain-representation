package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

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

//genisis functioncreates the genisis block which doesn't have a previous hash

func Genesis() *Block {
	new := CreateBlock("Genisis", []byte{})
	// fmt.Printf("---------------------BLOCK:GENISIS-------------------------------\n")
	// fmt.Printf("hash of previous block:%x\n", new.PrevHash)
	// fmt.Printf("data of the block:%s\n", new.Data)
	// fmt.Printf("hash of current block:%x\n", new.Hash)
	return new
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
