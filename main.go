package main

import "fmt"

func changeSlice(n []int) []int {
	n[0] = 10
	n = append(n, 11)
	return n
}

func main() {
	arr := [6]string{"This", "is", "a", "GO", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4] // ["is", "a", "GO"]
	fmt.Println(s)

	s1 := s[1:2] // ["a"] len = 1 cap = 4
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	sl := []int{1, 2, 3} // slice literal
	fmt.Println("Slice:", sl, "len:", len(sl), "capacity:", cap(sl))

	sm := make([]int, 3) // [0, 0, 0], len = 3,  cap = 3
	sm[0] = 5            // [5, 0, 0], len = 3,  cap = 3

	fmt.Println(sm)
	fmt.Println(len(sm))
	fmt.Println(cap(sm))

	sm1 := make([]int, 3, 5) // [0, 0, 0], len = 3, cap = 5
	sm1[0] = 5               // [5, 0, 0], len = 3,  cap = 5
	sm1[2] = 10              // [5, 0, 10], len = 3,  cap = 5

	fmt.Println(sm1)
	fmt.Println(len(sm1))
	fmt.Println(cap(sm1))

	var sn []int             // empty slice or nil slice []
	sn = append(sn, 1, 2, 3) // [1]
	fmt.Println(sn)

	// interview examples
	var x []int      // [], len = 0, cap = 0
	x = append(x, 1) // [1], len = 1, cap = 1
	x = append(x, 2) // [1,2], len = 2 , cap = 2
	x = append(x, 3)

	y := x

	x = append(x, 4)
	y = append(y, 5)

	x[0] = 10

	fmt.Println(x) // [10, 2, 3, 5]
	fmt.Println(y) // [10, 2, 3, 5]

	p := []int{1, 2, 3, 4, 5}
	p = append(p, 6)
	p = append(p, 7)

	l := p[4:]

	q := changeSlice(l)

	fmt.Println(p)
	fmt.Println(q)

}

func init() {
	fmt.Println("This will be invoked first")
}
