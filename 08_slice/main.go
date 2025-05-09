package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("welcome to slices")

	var s []string //Unlike arrays, slices in Go are more flexible and are dynamically sized. They do not have a fixed length defined in their type, unlike arrays which require a specified number of elements.
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	//to create a slice with non zero length , use the builtin "make"
	s = make([]string, 3)
	fmt.Println("emp:", s, "len: ", len(s), "cap: ", cap(s))

	//set and get just like arrays

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) //len returns the length of the slice as expected.

	//n addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")

	//Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)

	//slices upto (but excluding s[5])
	l = s[:5]
	fmt.Println("sl2:", 1)

	//slices up from (and including s[2])
	l = s[2:]
	fmt.Println("sl3:", l)

	//We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//The slices package contains a number of useful utility functions for slices.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")

	}

	//Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}

	}

	fmt.Println("2d :", twoD)
}
