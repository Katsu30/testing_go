package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s)) // 5
	fmt.Println(cap(s)) // 5

	t := s[0:2]
	fmt.Println(t) // [1, 2]

	x := append(s, 6, 7, 8)
	fmt.Println(x)
	y := append(s, x...)
	fmt.Println(y) // [1 2 3 4 5 1 2 3 4 5 6 7 8]

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := a[2:4:4]
	a3 := a[2:4:6]
	println(len(a2)) // 2
	println(cap(a2)) // 2
	println(len(a3)) // 2
	println(cap(a3)) // 4
}
