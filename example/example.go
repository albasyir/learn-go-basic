package example

// init function
func init() {
	println("example package init")
}

// accessable by other packages
func Hello(name string) string {
	return "Hello " + name
}

// not accessable by other packages
func hello(name string) string {
	return "Hello " + name
}

// accessable by other packages
var HelloVar = "Hello"

// not accessable by other packages
var helloVar = "Hello"
