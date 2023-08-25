package main

import "container/ring"

func main() {
	myring := ring.New(5)

	n := myring.Len()

	for i := 0; i < n; i++ {
		myring.Value = i
		myring = myring.Next()
	}

	myring.Do(func(p any) {
		println(p.(int))
	})

	println(myring.Link(myring))
}
