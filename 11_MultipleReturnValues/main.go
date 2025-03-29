package main

import "fmt"

func vals() (int, int) {
	return 3, 7
	//the(int , int)in this function signature shows that the function return 2 ints

}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
	//if you only want the subset of the returned values use the blank identifier
}
