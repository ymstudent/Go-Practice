package cn

//leetcode submit region begin(Prohibit modification and deletion)
import sort2 "sort"

var res47 [][]int

func permuteUnique(nums []int) [][]int {
	res47 = make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(nums))
	// 因为包含可重复数字，所以这里要先排序
	sort2.Ints(nums)
	backtrack47(nums, track, used)
	return res47
}

func backtrack47(nums, track []int, used []bool) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res47 = append(res47, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 保证重复的数字只会被填入一次
		if used[i] || i > 0 && !used[i-1] && nums[i] == nums[i-1] {
			continue
		}
		track = append(track, nums[i])
		used[i] = true
		backtrack47(nums, track, used)
		track = track[:len(track)-1]
		used[i] = false
	}
}

//leetcode submit region end(Prohibit modification and deletion)
