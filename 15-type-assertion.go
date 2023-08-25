package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)

	fmt.Println(resultString)

	// safe way
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value)
	case int:
		fmt.Println("Value", value)
	default:
		fmt.Println("Unknown type")
	}
}
