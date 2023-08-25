package main

import "fmt"

func main() {
	// make pure array
	// if we put []string{"a", "b", "c"} it will be slice
	var months = [...]string{
		"january",
		"february",
		"march",
		"april",
		"mei",
		"june",
		"july",
		"august",
		"september",
		"october",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1, len(slice1), cap(slice1))

	months[5] = "changed"
	slice1[0] = "CHANGED"

	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "new")
	fmt.Println(slice3)

	slice3[1] = "CHANGED"
	fmt.Println(slice3)

	fmt.Println(months)
}
