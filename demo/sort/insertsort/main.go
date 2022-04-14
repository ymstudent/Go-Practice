package main

import (
	"fmt"
)

func main() {
	a := []int{4, 5, 6, 1, 3, 2}
	b := insertSort(a)
	fmt.Println(b)
}

// 原地排序， 稳定，时间复杂度：最好 O(n), 最坏 O(n^2), 平均 O(n^2)
// 将数组中的元素区分为2个区间，已排序区间和未排序区间，初始区间只有一个元素，就是数组的第一个元素
// 插入算法的核心思想就是取未排序区间中的元素，在已排序区间中找到合适的位置将其插入，并保证已排序区间数据一直有序
// 插入排序包含2种操作：一种是元素比较，一种是元素移动
func insertSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	for i := 1; i < l; i++ {
		v := a[i]
		j := i - 1
		//fmt.Printf("i = %d\n", i)
		for ; j >= 0; j-- {
			//fmt.Println(j)
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		//fmt.Printf("j = %d\n", j)
		a[j+1] = v
	}
	return a
}
