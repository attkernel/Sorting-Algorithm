package main

import (
	"fmt"
)

//时间复杂度始终为n^2
func selectionSort(src []int) []int {
	for i := 0; i < len(src); i++ {
		for j := i + 1; j < len(src); j++ {
			if src[i] > src[j] {
				src[i], src[j] = src[j], src[i]
			}
		}
	}
	return src
}

func main() {
	fmt.Println(selectionSort([]int{3, 2, 6, 8, 1, 4, 7, 5}))
}
