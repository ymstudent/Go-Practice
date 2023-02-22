package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := l + (r-l)>>1
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				cnt++
			}
		}

		if cnt > mid {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
