package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var person Customer
	person.Name = "John"
	person.Address = "USA"
	person.Age = 30

	otherPerson := Customer{
		Name:    "Jane",
		Address: "Canada",
		Age:     25,
	}

	fmt.Println(person, otherPerson)
}
