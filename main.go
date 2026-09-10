package main

import "fmt"

// func a() {
// 	i := 0 // 0

// 	fmt.Println("first", i) // 0

// 	defer fmt.Println("second", i) //  fmt.Println("second", 0)

// 	i = i + 1 // i++ // 1

// 	fmt.Println("third", i) // 1

// 	defer fmt.Println("fourth", i) //  fmt.Println("second", 0)

// }

func sum(a int, b int) (result int) {
	result = a + b

	return
}

func main() {
	res := sum(3, 4)

	fmt.Println(res)
}
