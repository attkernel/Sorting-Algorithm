package main

import (
	"fmt"
)

func CountingSort(src []int) []int {
	s := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		count := 0
		for j := 0; j < len(src); j++ {
			if src[i] > src[j] {
				count++
			}
		}
		s[count] = src[i]
	}
	return s
}

func main() {
	fmt.Println(CountingSort([]int{5, 2, 6, 3, 1, 4}))
}
