package main

import "fmt"

func main() {
	// while loop
	var i = 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Print("For loop")

	// for loop
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
		continue
		
	}

	// range
	for i := range 3 {
		fmt.Println(i)
	}
}
