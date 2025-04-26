/*A pointer is a variable that stores the memory address of another variable.

Instead of holding a value directly, it points to a memory location where the value is stored.*/

package main

import "fmt"

func main() {
	// var a int = 10
	// var ptr *int = &a

	// fmt.Println("values of a :", a)
	// fmt.Println("Adress of a :", &a)
	// fmt.Println("Value of Ptr stores :", ptr)
	// fmt.Println("Value at ptr :", *ptr)

	/*Why do we use Pointers?
	To modify a variable inside a function (because functions get copies of values by default).

	To save memory by not copying large structures.

	To share data between different parts of a program. */

	a := 10
	fmt.Println("value of a before function change: ", a)
	changeValue(&a)
	fmt.Println("value of a after function change :", a)
}
func changeValue(val *int) {
	*val = 100
}
