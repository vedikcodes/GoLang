package main

import "fmt"

func main() {

	if 8%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if 8%4 == 0 {
		fmt.Println("its divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either of it are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")

	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
