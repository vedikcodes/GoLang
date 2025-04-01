package main

import "fmt"

//a function that will take an arbitary number of ints as argument
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}
func main() {

	//variadic function can be called in the usual way with individual arguments
	sum(1, 2)
	sum(1, 2, 3)

	//if you already have multiple args in slice , apply them to a variadic function using func(slice...)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
