package main

import "fmt"

func insertionSort(src []int, step int) []int {
	for idx := range src {
		pre := idx - step
		for pre >= 0 && src[pre] > src[pre+step] {
			src[pre], src[pre+step] = src[pre+step], src[pre]
			pre -= step
		}
	}
	return src
}

func shellSort(src []int) []int {
	step := (len(src) - 1) / 2
	for step > 0 {
		insertionSort(src, step)
		step = step / 2
	}
	return src
}

func main() {
	fmt.Println(shellSort([]int{4, 2, 6, 1, 7, 5, 3}))
}
