package main

import "fmt"

/**
 * Array must be fixed size
 *
 * Can't change after initialized
 */
func main() {
	var names [5]string

	names[0] = "budi"

	var values = [3]int{
		1, 2, 3,
	}

	fmt.Println(values)
}
