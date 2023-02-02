package cn

//leetcode submit region begin(Prohibit modification and deletion)
var res39 [][]int

func combinationSum(candidates []int, target int) [][]int {
	res39 = make([][]int, 0)
	track := make([]int, 0)
	backtrack39(candidates, track, 0, 0, target)
	return res39
}

func backtrack39(candidates, track []int, start, sum, target int) {
	if sum > target {
		return
	}
	if sum == target {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res39 = append(res39, tmp)
		return
	}
	// 元素如果不能重复使用，则递归时需要传i+1，如果能重复，则传i
	// 如果任意顺序都可以，则一直从0开始
	for i := start; i < len(candidates); i++ {
		track = append(track, candidates[i])
		sum += candidates[i]
		backtrack39(candidates, track, i, sum, target)
		track = track[:len(track)-1]
		sum -= candidates[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
