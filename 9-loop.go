package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan", counter)

		if counter%2 == 0 {
			counter++
			continue
		}

		if counter == 7 {
			break
		}

		counter++
	}

	for i := 1; i < 10; i++ {
		fmt.Println("Perulangan biasa", i)
	}
}
