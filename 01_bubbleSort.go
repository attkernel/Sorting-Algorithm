package main

import (
	"fmt"
)

func BubbleSort(src []int) []int {
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src)-1-i; j++ {
			if src[j] > src[j+1] {
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
	return src
}

//正序时最快 反序时最慢
func main() {
	fmt.Println(BubbleSort([]int{4, 2, 1, 5, 7, 3, 6}))
}
