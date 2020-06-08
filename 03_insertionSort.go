package main

import (
	"fmt"
)

func insertionSort(src []int) []int {
	for idx := range src {
		pre := idx - 1
		for pre >= 0 && src[pre] > src[pre+1] {
			src[pre], src[pre+1] = src[pre+1], src[pre]
			pre--
		}
	}
	return src
}

func insertionSortNew(src []int, step int) []int {
	for idx := range src {
		pre := idx - step
		for pre >= 0 && src[pre] > src[pre+step] {
			src[pre], src[pre+step] = src[pre+step], src[pre]
			pre -= step
		}
	}
	return src
}

func main() {
	fmt.Println(insertionSortNew([]int{4, 2, 6, 7, 1, 3, 5}, 1))
	fmt.Println(insertionSort([]int{}))
	fmt.Println(insertionSort([]int{4}))
	fmt.Println(insertionSort([]int{1, 2}))
	fmt.Println(insertionSort([]int{2, 1}))
}
