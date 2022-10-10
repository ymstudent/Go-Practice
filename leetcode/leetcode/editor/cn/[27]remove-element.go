package cn

//leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast, sz := 0, 0, len(nums)
	for fast < sz {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)
