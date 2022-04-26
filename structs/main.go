package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email: "teste@gmail.com",
			zip:   80520000,
		},
	}

	fmt.Printf("%+v", alex)
}
