package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := []int{4, 5, 6, 1, 3, 2}
	//b := quickSort(a)
	//fmt.Println(b)

	c := GetK(a, 4)
	//d := qucikSelect(a, 4)
	fmt.Println(c)
}

// 快速排序，不稳定的原地排序算法，时间复杂度：O(nlogn)，最坏为O(nlogn)
func quickSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	return quickSortC(a, 0, len(a)-1)
}

func quickSortC(a []int, low, high int) []int {
	if low >= high {
		return a
	}

	pivot := a[high] //选定一个分区点，一般为第一个元素或最后一个元素
	i := low
	for j := low; j < high; j++ { //遍历数组与分区点比较，大的放分区点右边小的放分区点左边
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	a[i], a[high] = a[high], a[i] //将分区点放到中间(这个位置就是分区点最终的位置)

	quickSortC(a, low, i-1)
	quickSortC(a, i+1, high)

	return a
}

// 思考：利用快排思想，在O(n)内查找第K大的元素

func GetK(a []int, k int) int {
	return quickSearch(a, 0, len(a)-1, k)
}

func quickSearch(a []int, low, high, index int) int {
	i := partition(a, low, high)
	fmt.Println(a)
	q := i + 1
	if q == index {
		return a[i]
	} else if q < index {
		return quickSearch(a, i+1, high, index)
	}

	return quickSearch(a, low, i-1, index)
}

func partition(a []int, low, high int) int {
	pivot := a[high]
	i := low
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	a[i], a[high] = a[high], a[i]

	return i
}

func qucikSelect(a []int, k int) int {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	for l, r := 0, len(a)-1; l < r; {
		v := a[l]
		i, j := l, r+1
		for {
			for i++; i < r && a[i] < v; i++ {
			}
			for j--; j > 1 && a[j] > v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[l], a[j] = a[j], v
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return a[k]
}
