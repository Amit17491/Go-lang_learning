// variable declartion with intial value

/*
package main

import (
	"fmt"
)

func main() {
	// var variable_name type = value
	var student1 string = "John" //type is string

	// in this case the type of variable is inferred from the value
	student2 := "Jane" //type is inferred
	x := 2             //type is inferred

	var num int = 15
	num1 := 22.5 // type is inferred   means compiler decides the variable type.

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(num)
	fmt.Println(num1)
}


// variable Declaration without inital value
package main

import (
	"fmt"
)

func main() {
	// Here we have declared the variable but we have not assigned the value to it.
	var a string // for this value it will print the blank value
	var b int    // for this it will print 0
	var c bool   // for this it will print False

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}


// declaring a variable outside the function and assigning the value inside the function.

package main

import (
	"fmt"
)

var a int
var b float32
var c string

func main() {
	a = 1
	b = 22.5
	c = "john"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// using inferred variable out side the function.

package main
import ("fmt")
a := 1     // op: # command-line-arguments .\hello.go:93:1: syntax error: non-declaration statement outside function body
func main() {
	fmt.println(a)

}


