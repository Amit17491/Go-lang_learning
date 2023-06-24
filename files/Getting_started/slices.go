// Go Slices

/*

   package main

   import "fmt"

   func main() {
   	// slice can grow and shrink
   	// Declaring a slice,
   	// var_name = []datatype{}

   	my_slice_variable := []int{1, 2, 3}

   	fmt.Println(my_slice_variable)
   	fmt.Println(len(my_slice_variable))
   	fmt.Println((cap(my_slice_variable)))
   }



   // create a slice From an Array

   package main

   import "fmt"

   func main() {
   	var myarray = [5]string{"vlovo", "Tata", "xylo", "car", "Audi"}
   	my_slice_variable_2 := myarray[2:4] // slice creatred from array from index 2 to 4 so len is 2 and cap is 4.
   	fmt.Println(myarray)
   	fmt.Println(my_slice_variable_2)
   	fmt.Println(len(my_slice_variable_2))
   	fmt.Println(cap(my_slice_variable_2))
   }



// create a slice from Make function

package main

import "fmt"

func main() {

	//create a slice using make function
	// var_name := make([]datatype, length, capacity)
	// if capacity is not defined length will be equal to capacity
	my_slice := make([]int, 5, 10)
	fmt.Println((my_slice))
	fmt.Println(len(my_slice))
	fmt.Println(cap(my_slice))

	// creating slice without capacity
	my_slice_2 := make([]int, 5) // if capacity is not defined, it will be equal to length.
	fmt.Println((my_slice_2))
	fmt.Println(len(my_slice_2))
	fmt.Println(cap(my_slice_2))

}



// Access, change, Append and copy Slices

package main

import "fmt"

func main() {
	slice_1 := []string{"This", "is", "me", "my", "name", "Amit"}

	fmt.Println(slice_1)

	//Access the element
	fmt.Println(slice_1[2])

	//change the element in the slice
	slice_1[2] = "baba"
	fmt.Println(slice_1)

	// appned element and append slice from one slice to another.
	// myvariable_name = append(myvariable_name, element_1, element_2)
	slice_1 = append(slice_1, "Assam", "mumbai")
	fmt.Println(slice_1)

	slice_3 := []string{"This", "is", "You", "my", "name", "Abhishek"}
	//slice_3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // cannot append string to int
	fmt.Println(slice_3)

	slice_3 = append(slice_3, slice_1...)
	fmt.Println(slice_3)

}


//change the length of the slice
package main

import "fmt"

func main() {

	array_1 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // This is an array
	//slice_4 := []int{1, 2, 4, 5}                 // this is slice
	fmt.Println(array_1)
	//fmt.Print(slice_4)
	slice_5 := array_1[1:5] // slice create from array
	fmt.Println(slice_5)

	//re changing the length of array

	slice_5 = array_1[1:4] // rechanging the slice
	fmt.Println(slice_5)

	slice_5 = append(slice_5, 20, 25)
	fmt.Println(slice_5)

}

*/

// memory efficiency copy function
package main
