package main

import (
	"flag"
	"fmt"
	"myBlockchain/blockchain"
	"os"
	"runtime"
	"strconv"
)

type commandLine struct {
	blockchain *blockchain.BlockChain
}

func (cli *commandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("add -block BLOCK DATA - add a block to the chain")
	fmt.Println("print - Prints the block in the chain")
}

func (cli *commandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *commandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (cli *commandLine) printChain() {
	iter := cli.blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Prev. hash:%x\n", block.PrevHash)
		fmt.Printf("Data:%v\n", block.Data)
		fmt.Printf("hash:%x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW:%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *commandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}
	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	defer os.Exit(0)

	chain := blockchain.InitBlockChain()
	// time.Sleep(time.Second * 2)

	// blockHeight := 1
	// for {
	// 	fmt.Printf("---------------------BLOCK:%v-------------------------------\n", blockHeight)
	// 	msg := fmt.Sprintf("block %v added", blockHeight)
	// 	chain.AddBlock(msg)

	// 	time.Sleep(time.Second * 2)
	// 	blockHeight++
	// }

	defer chain.Database.Close()

	cli := commandLine{chain}

	cli.run()
}
