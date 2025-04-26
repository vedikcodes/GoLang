package main

import "fmt"

//Go support anonymous functions
//useful when you want to define a afunction inline without having to name it

func intSeq() func() int {
	//So the intSeq() will only return the another func() closure function
	i := 0
	return func() int {
		//function working
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

}
