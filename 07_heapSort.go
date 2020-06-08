package main

import (
	"fmt"
)

/*func minHeap(root, l int, src []int) {
	for {
		child := 2*root + 1
		if child > l {
			break
		}
		if child+1 <= l && src[child+1] < src[child] {
			child += 1
		}
		if src[root] > src[child] {
			src[root], src[child] = src[child], src[root]
			root = child
		} else {
			break
		}
	}
}

func HeapSort(src []int) {
	l := len(src) - 1
	for root := l / 2; root >= 0; root-- {
		minHeap(root, l, src)
	}
	for end := l; end >= 0; end-- {
		if src[0] < src[end] {
			src[0], src[end] = src[end], src[0]
			minHeap(0, end-1, src)
		}
	}
}*/

func maxHeap(root, l int, src []int) {
	for {
		child := 2*root + 1
		if child > l {
			break
		}
		if child+1 <= l && src[child] < src[child+1] {
			child += 1
		}
		if src[root] < src[child] {
			src[root], src[child] = src[child], src[root]
			root = child
		} else {
			break
		}
	}
}

func HeapSort(src []int) {
	l := len(src) - 1
	for root := l / 2; root >= 0; root-- {
		maxHeap(root, l, src)
	}
	for end := l; end >= 0; end-- {
		if src[0] > src[end] {
			src[0], src[end] = src[end], src[0]
			maxHeap(0, end-1, src)
		}
	}
}

func main() {
	var s []int = []int{2, 28, 6, 15, 26, 22, 9, 10, 18, 233, 8}
	HeapSort(s)
	fmt.Println(s)
}
