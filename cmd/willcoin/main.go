package main

import (
	"fmt"
	"os"
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
			fmt.Println("init")
		case "add-block":
			// add a block to the blockchain
			fmt.Println("add-block")
		case "print-chain":
			// print all blocks on the blockchain
			fmt.Println("print-chain")
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