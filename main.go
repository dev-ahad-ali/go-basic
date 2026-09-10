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

func calculate() (result int) {
	fmt.Println("first", result)

	defer func() {
		result = result + 10
		fmt.Println("defer", result)
	}()

	result = 5
	fmt.Println("second", result)

	return
}

func calc() int {
	result := 0
	fmt.Println("first", result)

	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}

	defer show()

	result = 5
	fmt.Println("second", result)

	return result
}

func main() {
	a := calculate()
	b := calc()

	fmt.Println("main first", a)
	fmt.Println("main second", b)
}
