package main

import (
	"fmt"
)

type Test map[int]string
/*{
	1: "one",
	2: "two",
}*/

func main() {

	var fileExt = map[string]string {
		"Python": ".py",
		"C++": ".cpp",
		"Java": ".java",
	}

	t := make(Test)
	t[1] = "one"
	fmt.Println(t)

	fmt.Println(fileExt)

	delete(fileExt, "C++")

	fileExt["Julia"] = ".jl"

	fmt.Println(fileExt)
}
