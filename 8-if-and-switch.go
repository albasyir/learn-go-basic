package main

import "fmt"

func main() {
	name := "budi"

	if name == "aziz" {
		fmt.Println("ok")
	} else if name == "budi" {
		fmt.Println("ok, but this budi")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println(name, "length word is", len(name))

	switch name {
	case "aziz":
		fmt.Println("aziz is team")
	case "budi":
		fmt.Println("he's supervisor")
	default:
		fmt.Println("have no idea")
	}

	switch length := len(name); length > 4 {
	case true:
		fmt.Println("seams this not right")

	case false:
		fmt.Println("seams this right name")
	}

	switch {
	case len(name) > 4:
		fmt.Println("our team never have like this lang")

	case len(name) == 4:
		fmt.Println("most likely we have team member has lenght like this")
	}
}
