package main

import "fmt"

func main() {
	var arr [2]int

	arr[1] = 6
	arr[0] = 3

	fmt.Println(arr)
}

func init() {
	fmt.Println("This will be invoked first")
}
