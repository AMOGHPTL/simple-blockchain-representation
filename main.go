package main

import (
	"fmt"
	"myBlockchain/blockchain"
	"time"
)

func main() {
	chain := blockchain.Initial()
	time.Sleep(time.Second * 2)

	blockHeight := 1
	for {
		fmt.Printf("---------------------BLOCK:%v-------------------------------\n", blockHeight)
		msg := fmt.Sprintf("block %v added", blockHeight)
		chain.AddBlock(msg)
		time.Sleep(time.Second * 2)
		blockHeight++
	}
}
