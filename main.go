package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	edgar := person{firstName: "Edgar", lastName: "Serna"}
	fmt.Println(edgar)
}
