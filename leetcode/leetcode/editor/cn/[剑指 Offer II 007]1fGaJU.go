package cn

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]

		lo, hi := i+1, n-1 // 注意：这里左边界是从i+1起的，而不是从0
		for lo < hi {
			l, r := nums[lo], nums[hi]
			sum := l + r
			if sum > target {
				for hi > lo && nums[hi] == r {
					hi--
				}
			} else if sum < target {
				for hi > lo && nums[lo] == l {
					lo++
				}
			} else {
				res = append(res, []int{nums[i], nums[lo], nums[hi]})
				for hi > lo && nums[hi] == r {
					hi--
				}
				for hi > lo && nums[lo] == l {
					lo++
				}
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
