package main

import (
	"flag"
	"fmt"
)

// go run 21-package-flag.go -host='mydomain.com'
func main() {
	host := flag.String("host", "localhost", "Put your hostname")

	flag.Parse()

	fmt.Println(*host)
}
