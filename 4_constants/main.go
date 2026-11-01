package main

import "fmt"

func main() {

	const name = "Raj"

	fmt.Println(name)

	// multi contants

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port)
}
