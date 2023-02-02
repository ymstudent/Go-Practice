package cn

//leetcode submit region begin(Prohibit modification and deletion)
var res77 [][]int
func combine(n int, k int) [][]int {
	res77 = make([][]int, 0)
	track := make([]int, 0)
	backtrack77(track, n, k, 1)
	return res77
}

func backtrack77(track []int, n, k, start int)  {
	if len(track) == k {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res77 = append(res77, tmp)
		return
	}

	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack77(track, n, k, i+1)
		track = track[:len(track)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
