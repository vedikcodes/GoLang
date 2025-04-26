package main

import "fmt"

func main() {
	// var name string = "vedik"
	// var ptr *string = &name

	// fmt.Println("value of name:", name)
	// fmt.Println("Adress of name is : ", &name)
	// fmt.Println("value the ptr stores:", ptr)
	// fmt.Println("value of the ptr is:", *ptr)

	a := 100
	fmt.Println("value before func change :", a)
	double(&a)
	fmt.Println("values after func change :", a)

	b := 200
	fmt.Println("value before:", b)
	resetZero(&b)
	fmt.Println("Value reset to zero :", b)

}
func double(ptr *int) {
	*ptr = *ptr * 2
}

func resetZero(ptr *int) {
	*ptr = *ptr * 0
}
