package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "GO", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4] // ["is", "a", "GO"]
	fmt.Println(s)

	s1 := s[1:2] // ["a"] len = 1 cap = 4
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

func init() {
	fmt.Println("This will be invoked first")
}
