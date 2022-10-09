package cn

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum2(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1
	for lo < hi {
		sum := numbers[lo] + numbers[hi]
		if sum < target {
			lo++
		} else if sum > target {
			hi--
		} else if sum == target {
			return []int{lo + 1, hi + 1}
		}
	}
	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)
