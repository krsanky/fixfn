package main

import (
	"fmt"
	"os"

	"github.com/krsanky/fixfn"
)

func main() {
	//print_args()
	if len(os.Args[1:]) < 1 {
		usage()
	} else {
		dir := os.Args[1]
		fixfn.Inspect(dir)
	}
}

func print_args() {
	//fmt.Printf("len args: %d\n", len(os.Args[1:]))
	for i, a := range os.Args[1:] {
		fmt.Printf("%d:%s \n", i, a)
	}
}

func usage() {
	fmt.Printf("%s <dir>\n", os.Args[0])
}
