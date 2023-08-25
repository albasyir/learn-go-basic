package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])z")

	fmt.Println(regex.MatchString("aziz"))
	fmt.Println(regex.MatchString("aiz"))

	fmt.Println(regex.FindAllString("aziz aiz aoz", -1))
	fmt.Println(regex.FindAllString("aziz aiz aoz", 1))
}
