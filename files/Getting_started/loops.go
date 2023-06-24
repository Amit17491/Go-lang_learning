// for loop

/*
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}



package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {

		if i == 3 {
			continue // it will skip if i == 3 and will continue to the rest of the condition.
		}
		fmt.Println(i)
	}
}



package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {

		if i == 3 {
			break // it will skip if i == 3 and will break for  the rest of the condition as well.
		}
		fmt.Println(i)
	}
}


package main

import "fmt"

func main() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "mango"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], "-->", fruits[j])
		}
	}
}



package main

import "fmt"

func main() {
	fruits := [5]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Println(idx, val)
	}
}



package main

import "fmt"

func main() {
	fruits := [5]string{"apple", "orange", "banana"}
	for _, val := range fruits {
		fmt.Println(val)
	}
}

*/

package main

import "fmt"

func main() {
	fruits := [5]string{"apple", "orange", "banana"}
	for idx, _ := range fruits {
		fmt.Println(idx)
	}
}
