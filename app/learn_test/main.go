package main

import (
	"app/learn_test/alib"
	"fmt"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}
