// if else conditions.

package main

import (
	"fmt"
)

func main() {

	var v int = 200

	if v < 200 {
		fmt.Printf("v is less than 1000\n")
	} else if v < 500 {
		fmt.Printf("%v is less than 500\n", v)
	} else {
		fmt.Printf("value of v is %d\n", v)
	}

}
