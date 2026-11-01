package main

import (
	"fmt"
)

func main() {
	age := 18

	if age > 18 {
		fmt.Println("person is adult")
	} else if age == 18 {
		fmt.Println("Person is a teen")
	} else {
		fmt.Println("Person is a child")
	}

	// logical operators

	var name = "Raj"
	var has_permission = false

	if name == "Raj" && has_permission {
		println("Yes")
	} else {
		fmt.Println("No")
	}

	// declare inside if construct

	if age := 12; age < 18 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is an adult")
	}
}
