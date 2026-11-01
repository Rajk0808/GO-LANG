package main

import "fmt"

func main() {
	var nums [4]int

	fmt.Println(len(nums))

	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(nums)

	var vals [5]bool

	fmt.Println(vals)

	// declare in single line
	values_array := [3]int{1, 2, 3}

	fmt.Println(values_array)

	// 2d values_array
	array := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(array)
}
