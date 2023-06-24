// multiple variable declaration.

/*
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int = 1, 3, 5, 7 // here we are using type keyword so we can declare only one type of value.

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}



// multiple variable declaration in single line.

package main

import (
	"fmt"
)

func main() {
	var a, b, c, d = 1, 3.25, "john", 7 // here we are using not using type keyword so we can declare multiple type of value.

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

*/

// Go variable declaration in Block.

package main

import (
	"fmt"
)

func main() {
	var (
		a float32 = 55.5
		b int     = 1
		c string  = "hello"
	)
	y := 15

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(y)
}
