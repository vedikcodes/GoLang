package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("ITs the weekend")
	default:
		fmt.Println("its a weekday")
	}

}
