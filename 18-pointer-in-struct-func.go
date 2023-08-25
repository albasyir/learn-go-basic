package main

import "fmt"

type Man struct {
	name string
}

func (man *Man) married() {
	man.name = "Mr. " + man.name
}

func main() {
	aziz := Man{"Aziz"}

	aziz.married()

	fmt.Println(aziz)
}
