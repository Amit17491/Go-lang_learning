// slice data structure.

/* slice declaration.

package main

import (
	"fmt"
)

func main() {
	slice_1 := []int{1, 2, 3, 4, 5}
	myslice := []string{"vlovo", "Tata", "xylo", "car"}
	fmt.Println(slice_1)
	fmt.Println(myslice)
	fmt.Println(cap(slice_1))  // cap the number of elements the slice can grow or shrink to
	fmt.Println(cap(myslice))

}


// Create a slice From an Array.

package main

import (
	"fmt"
)

func main() {
	var array_1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := array_1[2:8]

	fmt.Println(myslice)
	fmt.Printf("Length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

}


//create a slice with the make function.
package main

import (
	"fmt"
)

func main() {
	myslice_1 := make([]int, 5, 10)
	fmt.Println(myslice_1)
	fmt.Println(len(myslice_1))
	fmt.Println((cap(myslice_1)))

	fmt.Println(" ")

	myslice_2 := make([]int, 5)
	fmt.Println(myslice_2)
	fmt.Println(len(myslice_2))
	fmt.Println((cap(myslice_2)))
}



// Go access, change, append and copy slices.

//Acess of element through the indexes  is same as array.
//change of element is similar to that of array.

//Append the element To a slice.

package main

import (
	"fmt"
)

func main() {
	myslice_1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	myslice_2 := append(myslice_1, 15, 14, 18, 19, 20)
	fmt.Println(cap(myslice_1))
	fmt.Println(myslice_2)
	fmt.Println(cap(myslice_2)) // old capacity * 2

}



// append one slice to another slice.

package main

import (
	"fmt"
)

func main() {
	myslice_1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	myslice_2 := []int{26, 27, 28}
	myslice := append(myslice_1, myslice_2...)
	fmt.Println(myslice)
	fmt.Println(len(myslice_1))
	fmt.Println(len(myslice_2))
	fmt.Println(cap(myslice)) // cap will always be *2 but not less than *2
}


// change the length of the slice.

package main

import (
	"fmt"
)

func main() {
	var array_1 = []int{1, 2, 3, 4, 5, 6}
	myslice := array_1[1:5]
	fmt.Println(myslice)

	myslice = array_1[0:3]
	myslice = append(myslice, 15, 16, 17, 18, 19, 20)
	fmt.Println(myslice)
}

*/

// for memoryt mgmt refer to copy function
