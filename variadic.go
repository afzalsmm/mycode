package main

import (
	"fmt"

)

func hello(names ...string){
	fmt.Println(names, " ")

	for _,name := range names {
		fmt.Println("hello", name)
	}
}

func main() {
	hello("one", "two")

	students := []string{"three", "four"}
	hello(students...)
}

