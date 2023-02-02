package cn

//leetcode submit region begin(Prohibit modification and deletion)
var res46 [][]int

func permute(nums []int) [][]int {
	res46 = make([][]int, 0)        // 保存结果
	track := make([]int, 0)         // 记录路径
	used := make([]bool, len(nums)) // 路径中的元素被标记为true，避免重复使用
	backtrack(nums, track, used)
	return res46
}

func backtrack(nums, track []int, used []bool) {
	if len(track) == len(nums) {
		// 这里要用深拷贝，不然track可能会被后续的append操作影响
		tmp := make([]int, len(track))
		copy(tmp, track)
		res46 = append(res46, tmp)
		return
	}

	// 任意顺序，也就是track中的数可以重复，所以每次从0开始
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		track = append(track, nums[i])
		used[i] = true
		backtrack(nums, track, used)
		track = track[:len(track)-1]
		used[i] = false
	}
}

//leetcode submit region end(Prohibit modification and deletion)
