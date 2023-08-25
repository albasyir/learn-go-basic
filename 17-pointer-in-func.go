package main

import "fmt"

type Address struct {
	City, Country, Province string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Sukabangun", "Palembang", "Sumatra Selatan"}

	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
