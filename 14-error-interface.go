package main

import (
	"errors"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh kosong")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagi(100, 0)

	if err == nil {
		println("Hasil", hasil)
	} else {
		println("Error", err.Error())
	}
}
