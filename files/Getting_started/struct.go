// struct ---> stands for structure and it is used to store multiple value of diff data types.

package main

import "fmt"

type person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 person
	var pers2 person

	pers1.name = "Hege"
	pers1.age = 20
	pers1.job = "Engineer"
	pers1.salary = 60000

	pers2.name = "Rakesh"
	pers2.age = 22
	pers2.job = "pilot"
	pers2.salary = 50000

	// two method to print the values
	// method 1
	// fmt.Println("Name:", pers1.name)
	// fmt.Println("Age:", pers1.age)
	// fmt.Println("Job:", pers1.job)
	// fmt.Println("Salary:", pers1.salary)

	// method 2
	printperson(pers1)

}

// for function 2

func printperson(pers person) {
	fmt.Println("Name:", pers.name)
	fmt.Println("Age:", pers.age)
	fmt.Println("Job:", pers.job)
	fmt.Println("Salary:", pers.salary)
}
