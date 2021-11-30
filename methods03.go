package main

import (
	"fmt"
)

type Virtmach struct{
    ip string
    hostname string
    diskgb int
    ram int
}

func (v *Virtmach) setip(ip string) {
	v.ip=ip
}

func (v *Virtmach) expanddisk(gb int){
	v.diskgb += gb
}

func (v Virtmach) display() {
	fmt.Println(v)

	fmt.Printf("%v\n", v)
	fmt.Printf("%+v\n", v)
	fmt.Printf("%#v\n", v)
}

func main() {
	vm := Virtmach{"1.0.0.1", "king", 10, 16}

	vm.setip("2.0.0.2")
	vm.expanddisk(5)
	vm.display()
}

