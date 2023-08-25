package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.New()

	list.PushBack(3)
	list.PushBack(2)
	list.PushBack(1)

	for element := list.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	for element := list.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
