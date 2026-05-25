package main

import (
	"fmt"
	"os"
	"github.com/willboyle18/willcoin/internal/blockchain"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: willcoin <command>")
	}
	command := args[1]
	switch command {
		case "init":
			// create blockchain
			blockchain.NewBlockchain()
		case "add-block":
			// add a block to the blockchain
			blockchain.AddBlock("This is another block")
		case "print-chain":
			// print all blocks on the blockchain
			blockchain.PrintBlockchain()
		case "validate":
			// validate blockchain
			fmt.Println("validate")
		case "help":
			// show user manual
			fmt.Println("help")
		default:
			fmt.Println("unknown command:", command)
	}
}