package main

import (
	"GoBlockchain/block"
	"GoBlockchain/wallet"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	/*block := &block2.Block{}
	fmt.Printf("%x\n", block.Hash())
	*/

	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())

	myBlockchainAddress := "192.168.1.2"
	blockChain := block.NewBlockchain(myBlockchainAddress)
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("Sandra", "John", 2.0)
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("John", "Sandra", 2.5)
	blockChain.AddTransaction("Sandra", "Peter", 0.5)
	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("my 		%.1f\n", blockChain.CalculateTotalAmount(myBlockchainAddress))
	fmt.Printf("A 		%.1f\n", blockChain.CalculateTotalAmount("A"))
	fmt.Printf("B 		%.1f\n", blockChain.CalculateTotalAmount("B"))
	fmt.Printf("Sandra	%.1f\n", blockChain.CalculateTotalAmount("Sandra"))
	fmt.Printf("John 	%.1f\n", blockChain.CalculateTotalAmount("John"))
	fmt.Printf("Peter	%.1f\n", blockChain.CalculateTotalAmount("Peter"))
}
