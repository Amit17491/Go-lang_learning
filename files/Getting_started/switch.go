// single switch cases

/*
package main

import "fmt"

func main() {

	day := 9

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("Sorry no match")
	}
}

// case type must be same otherwise error

package main

import "fmt"

func main() {
	food := 5

	switch food {
	case 1:
		fmt.Println("Swaad")
	case b: //------> the switch case must of same type.  .\switch.go:46:7: undefined: b
		fmt.Println("Mango")
	}
}

*/

// Multi switch case

package main

import "fmt"

func main() {

	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd day in a week")
	case 2, 4:
		fmt.Println("Even weekdays")
	case 6, 7:
		fmt.Println("Weekends")
	default:
		fmt.Println(" No match")
	}
}
