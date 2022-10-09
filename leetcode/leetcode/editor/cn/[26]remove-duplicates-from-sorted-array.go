package cn
//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	sz := len(nums)
	if sz < 1 {
		return sz
	}
	slow, fast := 0, 1
	for fast < sz {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		fast++
	}
	return slow + 1
}
//leetcode submit region end(Prohibit modification and deletion)
