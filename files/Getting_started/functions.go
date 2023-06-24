// making  of a function

/*

package main

import "fmt"

func Mymessage() {

	fmt.Println("This is the first message")
}

func Mysecondmessage() {
	fmt.Println("Hi Amit")
	Mymessage() // calling first function inside 2nd message.
}

func main() {
	Mysecondmessage() //calling 2nd function.

}


//Go Function Parameters and Arguments

// func  FunctionName (parma_1 type, parma_2 type, param_3 type){
// code to be executed
// }

package main

import "fmt"

func familyName(fname string, num int) {
	fmt.Println("Hello", fname, "shukla", "my age is", num)
	// fname and num is a parameter
}

func main() {
	familyName("sonali", 15) // sonali & 15 is an argument.
	//sonali and 15 is an argument.
}



// Go return function Return values
// fun FunctionName(param1 type, param type2) type{
// code to executed
// return output
// }

package main

import "fmt"

func myFunction(x int, y int) int {
	return x + y

}

func main() {
	fmt.Println(myFunction(1, 2))
}



// Named the Return  Values

package main

import "fmt"

func myFunction(x int, y int) (result int) {
	// here declaring the return type value
	result = x + y
	return // This is called naked return  or we can use
	// or //
	// return result
}

func main() {
	value := myFunction(1, 2)
	fmt.Println(value)
}


package main

import "fmt"

func myFunction(x int, y string) (result int, txt_1 string) {
	result = x + x
	txt_1 = y + " World!"
	return
}

func main() {
	a, b := myFunction(5, "Hello") // here the the retrun of  x + x is stored in a and Hello World is stored in b.
	fmt.Println(a, b)
}

*/

// skip the return value of the variable.
package main

import "fmt"

func myFunction(x int, y string) (result int, txt_1 string) {
	result = x + x
	txt_1 = y + " World !"
	return
}

func main() {
	_, b := myFunction(6, "World!")
	fmt.Println(b)

	a, _ := myFunction(6, "World!")
	fmt.Println(a)

}
