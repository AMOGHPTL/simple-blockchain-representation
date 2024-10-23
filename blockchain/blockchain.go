package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

// Blockchain is a struct which contains two fields last hash and database
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

// blockchain iterator is used to print the blockchain from the db
type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// initblockchain func is modified now this function intializes the blockchain in the database
// initially it opens the database and handles the error if the blockchain is not opened
// now we chack if the database alreday contains the lh(lasthash item) if not the genisis block is called and the lh is assigned to the hash pf genisis block
// if the lh is found then the lh value is take to be put in the new created block as the hash of the previous block
// initblockchain func returns a blockchain
func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.Value()
			return err
		}
	})

	Handle(err)

	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

//addblock function is modified to add blocks inside the database
//addblock function adds the block to the existing blockchain and takes data of the block as the input and it doesnt return anything

// func (chain *BlockChain) AddBlock(data string) {
// 	var lastHash []byte

// 	err := chain.Database.View(func(txn *badger.Txn) error {
// 		item, err := txn.Get([]byte("lh"))
// 		Handle(err)
// 		lastHash, err = item.Value()

// 		return err
// 	})
// 	Handle(err)

// 	newBlock := CreateBlock(data, lastHash)

// 	err = chain.Database.Update(func(txn *badger.Txn) error {
// 		err := txn.Set(newBlock.Hash, newBlock.Serialize())
// 		Handle(err)
// 		err = txn.Set([]byte("lh"), newBlock.Hash)

// 		chain.LastHash = newBlock.Hash

// 		return err
// 	})
// 	Handle(err)
// }

func (chain *BlockChain) AddBlock(data string) {
	var lasthash []byte

	err := chain.Database.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lasthash, err = item.Value()
		return err
	})
	newblock := CreateBlock(data, lasthash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newblock.Hash, newblock.Serialize())
		Handle(err)

		err = txn.Set([]byte("lh"), newblock.Hash)
		return err
	})
	Handle(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)
		encodedBlock, err := item.Value()
		block = Deserialize(encodedBlock)

		return err
	})
	Handle(err)

	iter.CurrentHash = block.PrevHash

	return block
}
