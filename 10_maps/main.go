package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["name"] = "Raj"

	m["expertise"] = "CSE"

	fmt.Println(m) // map[expertise:CSE name:Raj]
	// none existing key
	fmt.Println(m["phone"]) // -> " "

	// delete
	delete(m, "expertise")

	fmt.Println(m)

	

}
