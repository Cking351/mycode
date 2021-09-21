package main

import "fmt"

type User struct {
	Name string
	Pet []string
}

// A method type user
func (p2 *User) newPet() {
	 fmt.Println(*p2, "underlying value of p2 before")
         p2.Pet = append(p2.Pet, "Fido")
         fmt.Println(*p2, "underlying value of p2 after")
}

func main() {
	u := User{Name: "Anna", Pet: []string{"Bailey"}}
	fmt.Println(u, "u before")

	p := &u // Let's make a pointer to u

	p.newPet()
	fmt.Println(u, "u after")
}
