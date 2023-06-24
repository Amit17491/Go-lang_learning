// Go array playground:
package main

import (
	"fmt"
)

func main() {
	var array_1 = [5]int{1, 2, 3, 4, 5}
	var array_2 = [...]int{6, 7, 8, 9, 10} //complier picksup the size.

	fmt.Println(array_1)
	fmt.Println(array_2)

	//length of the array
	fmt.Println(len(array_1))
	fmt.Println(len(array_2))

	fmt.Println(" ")

	// seconde type of array declaration.

	arr_1 := [5]float64{1.1, 2.1, 3.1, 4.1, 5.1}
	arr_2 := [...]int{5, 6, 7, 8, 9}

	fmt.Println(arr_1)
	fmt.Println(arr_2)

	// now replacing the value of array. //

	var str1 = [5]string{"Vlovo", "Bmw", "Ford", "Mazda"}
	fmt.Println(str1)

	// Access Element of the Array.
	fmt.Print(str1[2], "\n") // by the index.
	fmt.Println(str1[0:3])   // using the slicing.

	//changing the value of the array.

	str1[2] = "TATA" // value changes form FORD to TATA.
	fmt.Println(str1)

	//array initiization.
	var ar_1 = [5]int{}              // not initlized
	var ar_2 = [5]int{1, 2}          // partially initialized
	var ar_3 = [5]int{1, 2, 3, 4, 5} // fully initiazlied.

	// declaring only the specific element

	var ar_4 = [...]int{1: 10, 4: 30} // upto 3 dots for empty leng.

	fmt.Println(ar_1)
	fmt.Println(ar_2)
	fmt.Println(ar_3)
	fmt.Println(ar_4) // it will only give the value at 1st and 4th place.
}
