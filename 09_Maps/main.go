package main

import (
	"fmt"
	"maps"
)

func main() {
	// to create an empty map : use Make(map[key-type]val-type)

	m := make(map[string]int)

	//set key/value pairs using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map :", m) //print key-value pairs

	//get the value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//if the key doesnt exist the zero value of the value type is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	//the builtin len return the number of key-value pairs when called on maps

	fmt.Println("len:", len(m))

	//the builtin delete remove key/value pairs from a map
	delete(m, "k2")
	fmt.Println("new map:", m)

	_, present := m["k2"]
	fmt.Println("present:", present)

	//we can also declare and initialize a new map in the same line

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// some utility function in Maps
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n==n2")
	}
}
