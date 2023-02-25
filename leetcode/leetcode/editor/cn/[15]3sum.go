package cn

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)

func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 3, 0, 0)
}

func nSumTarget(nums []int, n, start, target int) [][]int {
	sz := len(nums)
	res := make([][]int, 0)
	if n < 2 || sz < n {
		return res
	}
	if n == 2 {
		// 基础是两数之合
		lo, hi := start, sz-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				// 跳过重复元素
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else if sum == target {
				res = append(res, []int{left, right})
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else {
		// n > 2时，递归计算 (n-1)Sum的结果
		for i := start; i < sz; i++ {
			sub := nSumTarget(nums, n-1, i+1, target-nums[i])
			for _, v := range sub {
				// (n-1)Sum加上nums[i]就是nSum
				v = append(v, nums[i])
				res = append(res, v)
			}
			// 跳过重复元素
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
