package main

import "fmt"

func MergeSort(src []int) []int {
	l := len(src)
	if l < 2 {
		return src
	}
	mid := l / 2
	left := src[:mid]
	right := src[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	res := make([]int, 0)
	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		res = append(res, left[:]...)
	} else {
		res = append(res, right[:]...)
	}
	return res
}

func main() {
	arr := []int{3, 1, 2, 5, 6, 43, 4}
	fmt.Println(MergeSort(arr))
}
