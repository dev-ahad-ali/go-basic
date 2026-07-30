package main

import "fmt"

func print(numbers *[3]int) {
	fmt.Println(numbers)
}

// type User struct {
// 	Name   string
// 	Age    int
// 	Salary float64
// }

func main() {
	// pointer or address of memory (ram)
	x := 10

	fmt.Println("x = ", x) // x = 20

	p := &x // ampersand & => address of

	*p = 30

	fmt.Println("x = ", x)
	fmt.Println("Address: ", p)              // p is the address of x
	fmt.Println("Value at the address:", *p) // * =>  value at address

	arr := [3]int{1, 2, 3}
	print(&arr)
}

func init() {
	fmt.Println("This will be invoked first")
}
