package main

import "fmt"

type Filter func(string) bool

func filterKataKasar(name string) bool {
	return name != "anjing"
}

func sayHello(firstName string, lastName string, filter Filter) {
	if !filter(firstName) || !filter(lastName) {
		fmt.Println("word not allowed")
		return
	}

	// _ is skipper return value
	firstName, _, lastName = multiReturn(firstName, lastName)

	fmt.Println("Hello,", returnHello(firstName, lastName))
}

func returnHello(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func multiReturn(firstName string, lastName string) (string, string, string) {
	return firstName, " ", lastName
}

// variadic
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func factorial(number int) int {
	if number == 1 {
		return 1
	}

	return number * factorial(number-1)
}

func logging() {
	fmt.Println("logs done!")
}

func recovery_mode() {
	errorMsg := recover()

	if errorMsg != nil {
		fmt.Println("error was", errorMsg)
	}
}

func main() {
	// RUN FUNCTION AT END TIME MAIN EXECUTED, EVENT HAS ERROR OR SUCCESS
	defer logging()
	defer recovery_mode()

	sayHello("Abdul Aziz", "anjing", filterKataKasar)
	sayHello("TEST", "Al Basyir", filterKataKasar)

	total := sumAll(1, 2, 5)
	fmt.Println(total)

	slice := []int{1, 2, 5}
	total_slice := sumAll(slice...)
	fmt.Println(total_slice)

	myfunction := sayHello
	myfunction("My", "function", filterKataKasar)

	// callback
	sayHello("Budi", "Astaf", func(name string) bool {
		return true
	})

	fmt.Println(factorial(5))

	panic("meledak")

}
