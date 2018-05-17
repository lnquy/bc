package main

import (
	"github.com/lnquy/bc/block"
	"fmt"
)

func main() {
	b := block.GenesisBlock()
	fmt.Printf("%v", b)
}
