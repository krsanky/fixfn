package main

import (
	"fmt"
	"os"
)

func main() {
	for i, a := range os.Args[1:] {
		fmt.Printf("%d:%s ", i, a)
	}
	fmt.Println()

	fmt.Printf("Hey!\n")
}
