package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	edgar := person{
		firstName: "Edgar",
		lastName:  "Serna",
		contactInfo: contactInfo{
			email:   "eserna03@gmail.com",
			zipCode: 72758,
		},
	}

	edgar.updateName("ed")
	edgar.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
