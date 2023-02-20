package cn

//leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	// 从后往前，寻找第一个相邻的升序段
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		// 从后往前，寻找第一个大于num[i]的数
		for nums[i] >= nums[k] {
			k--
		}
		// 交换i,k位置
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 将[j,end]翻转
	for i, j := j, len(nums) - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
