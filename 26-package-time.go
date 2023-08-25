package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println(time.Now().Format("2006-01-02"))
}
