package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100 //We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
	fmt.Println("set :", a)
	fmt.Println("get", a[4])

	fmt.Println("len:", len(a)) //The builtin len returns the length of an array.

	b := [5]int{1, 2, 3, 4, 5} //Use this syntax to declare and initialize an array in one line.
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} //You can also have the compiler count the number of elements for you with ...
	fmt.Println("idx:", b)         //If you specify the index with :, the elements in between will be zeroed.

	var twoD [2][3]int //
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d :", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
