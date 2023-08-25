package main

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func emptyInterface() interface{} {
	return "String"
}

func main() {
	person := Person{Name: "John"}

	var test interface{} = emptyInterface()

	sayHello(person)
}
