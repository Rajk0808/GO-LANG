package main

import "fmt"

func main() {
	// when multiple cases are there

	i := 5

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("other")
	}
}
