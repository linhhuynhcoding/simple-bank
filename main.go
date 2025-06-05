package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	var x int = "string" // This should trigger an error: cannot use "string" (type string) as int
}
