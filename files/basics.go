/*
package main // package name

import (
	"fmt"           // importing file from the lib it import the basic input/ouput propgram like println..etc
)

func main() {
	fmt.Println("Hello world ")   // using println funct from fmt

}

package main

import (
	"fmt"
)

func main() {
	var student1 string = "John" //type is string
	student2 := "Jane"           //type is inferred
	x := 2                       //type is inferred
	var num int = 15
	num1 := 22.5 // type is inferred   means compiler decides the variable type.

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(num)
	fmt.Println(num1)
}
// variable declaration without initalization

package main

import (
	"fmt"
)

func main() {
	var a string // values are " "
	var b int // 0
	var c bool  // false

	fmt.Println(a) // op blank since no value is assign.
	fmt.Println(b) // 0
	fmt.Println(c) // false since no value is assign.
}

// value assignment after declaration

package main

import (
	"fmt"
)

func main() {
	var student string   // declaring a variable
	student = "John"     // assign the value
	fmt.Println(student)
}


// declaring a variable outside the function and assigning the value inside the function.

package main

import (
	"fmt"
)
var a int
var b int = 2
var c = 3

func main() {
	a = 1
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


// multiple variable declaration.

package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}


// refer doc for this declaring multiple variable.
package main

import (
	"fmt"
)

func main() {
	var a, b = 6, "Hello"
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}


// or
package main

import (
	"fmt"
)

func main() {

	var (
		a int
		b int    = 1
		c string = "hello"
		d int    = 5
	)
	y := 15

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(y)
}

// GO constant rule and expression.

package main

import (
	"fmt"
)

func main() {

	const PI = 3.14
	fmt.Println(PI)
	// Two type of constant declartion
	const A = 15
	A = 22  // this is not possible as const are unchnagable and read-only.
	const b int = 20
	fmt.Println(A)
	fmt.Println(b)

}

// print types

// print
// printf
// println

//GO FORMATTING VERBS.

// example of printf  with formatting
// variable %V and  %T.
//some more formatting , refer w3school.
package main

import (
	"fmt"
)

func main() {
	var i string = "Hello world"
	var j int = 15

	//value will be printed in 2 lines.
	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)

	// value will be printed in single line.
	fmt.Printf("i has value: %v and type: %T", i, i)
	fmt.Printf("j has value: %v and type: %T\n", j, j)

	fmt.Printf("%v\n", i)   // it will print the Hello world
	fmt.Printf("%#v\n", i)  // it will print Hello world in quotes.
	fmt.Printf("%v%%\n", i) //
	fmt.Printf("%T\n", i)   // it will give type
	fmt.Printf("%s\n", i)   // prints the value as plain string.
	fmt.Printf("%q", i)     // prints the value in double-quoted string.
	fmt.Println("%x", i)    // prints the value in hex
	fmt.Printf("% x", i)
}
*/

package main

import (
	"fmt"
)

func main() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
}
