package main

import (
	"fmt"
	"os"
)

func main() {
	//len := len(os.Args[1:])

	for i, a := range os.Args[1:] {
		fmt.Printf("Arg %d is %s\n", i+1, a)
	}
}
