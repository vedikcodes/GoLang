package main

import "fmt"

func main() {

	i := 1
	for i <= 3 { //The most basic type, with a single condition.
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ { //A classic initial/condition/after for loop.

		fmt.Println(j)

	}

	for i := range 3 {
		fmt.Println("range ", i)
	}

	for {
		fmt.Println("loop")
		break
	}
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
