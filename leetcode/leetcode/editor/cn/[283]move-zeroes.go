package cn

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	slow, fast, sz := 0, 0, len(nums)
	for fast < sz {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}

//leetcode submit region end(Prohibit modification and deletion)
