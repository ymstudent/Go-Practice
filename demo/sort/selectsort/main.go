package main

import "fmt"

func main() {
	a := []int{4, 5, 6, 1, 3, 2}
	b := selectSort(a)
	fmt.Println(b)
}

// 原地排序， 不稳定，时间复杂度：最好 O(n^2), 最坏 O(n^2), 平均 O(n^2)
// 选择排序，将数组分为已排序和未排序2部分，从未排序的部分中找到最小的元素，将其放入已排序区间的末尾
func selectSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	for i := 0; i < l; i++ {
		min := i
		for j := i; j < l; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}

	return a
}
