package main

import "fmt"

func main() {
	a := []int{4, 5, 6, 1, 3, 2}
	b := bubbleSort(a)
	fmt.Println(b)
}

// 原地排序， 稳定，时间复杂度：最好 O(n), 最坏 O(n^2), 平均 O(n^2)
func bubbleSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	for i := 1; i < l; i++ {
		for j := 0; j < l-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}
