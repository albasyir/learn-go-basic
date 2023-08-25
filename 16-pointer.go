package main

import "fmt"

type Address struct {
	City, Country, Province string
}

func main() {
	// create address1
	address1 := Address{"Sukabangun", "Palembang", "Sumatra Selatan"}

	// copy address1 to address2
	var address2 Address = address1

	// copy reference address1 to address3
	var address3 *Address = &address1

	// copy reference address2 to address4
	var address4 *Address = &address2

	// change copied address2
	address2.Country = "Jakarta"

	// change copied reference address3
	address3.Country = "Bandung"

	// create new address and change address4 pointer
	address4 = &Address{"Malang", "Jawa Timur", "Indonesia"}

	// force all pointer
	var address5 *Address = address4
	*address4 = Address{"A", "B", "C"}

	var address6 *Address = new(Address)

	fmt.Println("address1", address1)
	fmt.Println("address2", address2)
	fmt.Println("address3", address3)
	fmt.Println("address4", address4)
	fmt.Println("address5", address5)
	fmt.Println("address6", address6)
}
