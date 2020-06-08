package main

import (
	"fmt"
)

func quickSort(src []int) {
	_quickSort(src, 0, len(src)-1)
}

func _quickSort(src []int, left, right int) {
	if left < right {
		partitionIndex := partition(src, left, right)
		_quickSort(src, left, partitionIndex-1)
		_quickSort(src, partitionIndex+1, right)
	}
}

func partition(src []int, left, right int) int {
	pivot := left
	index := left + 1
	for i := index; i <= right; i++ {
		if src[i] < src[pivot] {
			src[i], src[index] = src[index], src[i]
			index++
		}
	}
	src[pivot], src[index-1] = src[index-1], src[pivot]
	return index - 1
}

func main() {
	var s []int = []int{4, 2, 5, 6, 3, 1, 7}
	quickSort(s)
	fmt.Println(s)
}
