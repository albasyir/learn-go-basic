package main

import "fmt"

func main() {
	a := len("aaa")

	fmt.Println(a)
	fmt.Println("aaaa"[0])

	var xx = "aaaa"
	var yy string = "aaaa"
	var aaa uint = 1

	var (
		c = "1"
		b = 5
	)

	const (
		first_name string = "aziz"
		last_name         = "al"
	)

	var y int16 = 32767
	var x int8 = int8(y)

	fmt.Println(x)

	var a_from_aziz byte = first_name[0]
	var a_string_from_aziz = string(a_from_aziz)
}
