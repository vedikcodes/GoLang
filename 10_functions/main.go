package main

import "fmt"

//here is function that takes two ints and return their sum as an int
func plus(a int, b int) int {

	return a + b

}

//multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters
//upto the final parameter that declare the type

func plusPlus(a, b, c int) int {
	return a + b + c
}

//call a function

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}
