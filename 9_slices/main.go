package main

import "fmt"

func main() {
	var nums []int

	fmt.Println(nums)

	nums = append(nums, 1)

	fmt.Println(nums)

	values := make([]int, 3, 5)

	fmt.Println(cap(values))
	fmt.Println(values)
}
