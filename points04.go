/* RZFeeser | Alta3 Research
   Without Pointers - Why we need pointers  */    

package main

import (
 "fmt"
)

type User struct {
 Name string
 Pets []string
}

func (u User) newPet() User {
 u.Pets = append(u.Pets, "Lucy")                    // Simple function should ensure "Fido" is added to User
 fmt.Println(u, &u)                                     // This user is **NOT** the same user as the one in main()!!
 return u
}

func main() {
 u := User{Name: "Anna", Pets: []string{"Bailey"}}
 
 var p *User

 p = &u
 fmt.Println(*p, p)

 n := u.newPet()                                         // {Anna [Bailey Fido]} -- This *should* add "Fido" to "u"
 fmt.Println(u, &u) 					// We **DO NOT** see Fido -- what happened?
 
 fmt.Println(n, &n)
}

