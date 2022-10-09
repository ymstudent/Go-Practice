package cn

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	a := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		if index, ok := a[sub]; ok {
			return []int{i, index}
		} else {
			a[nums[i]] = i
		}
	}
	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)
