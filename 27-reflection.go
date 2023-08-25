package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func main() {
	sample := Sample{"sample"}
	reflect := reflect.TypeOf(sample)

	fmt.Println(reflect.NumField())
	fmt.Println(
		reflect.Field(0).Name,
		reflect.Field(0).Tag.Get("required"),
		reflect.Field(0).Tag.Get("max"),
	)
}
