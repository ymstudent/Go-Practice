package main

import "fmt"

func main() {
	a := []int{2, 4, -9, 5, 2, 1, -1, 7}
	fmt.Println(maxSubArr(a))
}

// 最大子数组和，简单
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 遍历数组，元素相加，只要和为正，就累计，直到和为负时，重新从正数的节点开始累计
// 时间复杂度：O(n)
func maxSubArr(nums []int) int {
	sum := 0
	ans := nums[0]
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if sum > ans {
			ans = sum
		}
	}
	return sum
}

//进阶：分治法
