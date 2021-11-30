package main

import "fmt"

type Astro struct {
	name string
	age int
	lastmission string
	isneeded bool
}

type NasaMission struct {
	people []Astro
	number int
	message string
}



func main() {
	ast1 := Astro{"Megan McArthur", 35, "ISS", false}
    ast2 := Astro{"Barry Wilmore", 41, "Boeing Crew Flight Test (CFT)", true}
    ast3 := Astro{"Raja Chari", 39, "SpaceX Crew-3", true}

    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)
    
    p := [] Astro {ast1, ast2, ast3}

    fmt.Println(p)

    fmt.Println(p[2].lastmission)
    
    s :=  NasaMission{p, 3, "success"}

    fmt.Println(s)
    fmt.Println("%+v", s)
    
    }

