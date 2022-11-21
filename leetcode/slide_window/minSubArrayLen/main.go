package main

import "math"

func main() {
	println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

// 长度最小的子数组
// 地址：https://leetcode.cn/problems/minimum-size-subarray-sum/
// 难度：中等
func minSubArrayLen(target int, nums []int) int {
	i, sum, res := 0, 0, math.MaxInt64
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target { //滑动窗口
			subLen := j - i + 1 // 滑动窗口长度
			if res > subLen {
				res = subLen
			}
			sum -= nums[i]
			i++ //变更起始位置
		}
	}
	if res == math.MaxInt64 {
		res = 0
	}
	return res
}
