// Go Arrays

/*
package main

import "fmt"

func main() {
	// with var keywords
	//var arrya_name = [length]datatypes{values}
	var students = [5]string{"A", "B", "C", "D", "E"}
	fmt.Println(students)

	// without length of the array
	// var array_name = [...]datatype{values} // here length is inferred
	var students_1 = [...]string{"A", "B", "C", "D", "E"}
	fmt.Println(students_1)

	// Without the var keyword
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	numbers_1 := [...]int{1, 2, 3, 4, 5} //here length is inferred
	fmt.Println(numbers_1)
}

*/

//Access element of the array.

package main

import "fmt"

func main() {
	var students_2 = [5]int{}                                                     // array not initialize //will give all five blank --->[0 0 0 0 0]
	var students_3 = [5]int{1, 2, 3}                                              // array partially initialize // will give the other -->[1 2 3 0 0]
	students_4 := [5]string{"a", "b", "c", "d"}                                   // array partially initalize
	name_of_places := [10]string{1: "Delhi", 3: "Goa", 6: "Pune", 8: "Hyderabad"} // 1:Dehi means delhi will be at index 1.
	fmt.Println(students_2)
	fmt.Println(students_3)
	fmt.Println(students_4)
	fmt.Println(name_of_places)
	fmt.Println(len(name_of_places)) // length of the array.

	// Access the element of an array
	fmt.Println(students_4[3])

	// Change the element of the array
	students_4[3] = "Amit"
	fmt.Println(students_4)
}
