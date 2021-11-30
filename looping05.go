package main

import (
	"fmt"
)

func main() {

	//var m = make(map[int]string)
	m := map[int]string {
		11: "one",
		2: "two",
	}
	
	for pos, val := range m {
		fmt.Println("Position %d has value of %s/n", pos, val)
	}
	fmt.Println(m)

}
		
