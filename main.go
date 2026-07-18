package main

import "fmt"

func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}

func call() func(x int, y int) {
	return add
}

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	// higher order function
	processOperation(2, 5, add)
	sum := call() // function expression
	sum(2, 8)
}

func init() {
	fmt.Println("This will be invoked first")
}
