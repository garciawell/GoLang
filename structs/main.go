package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	name := person{
		firstName: "Wellington",
		lastName:  "Garcia",
		contactInfo: contactInfo{
			email: "teste@gmail.com",
			zip:   80520000,
		},
	}

	name.updateName("Jimmy")
	name.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
