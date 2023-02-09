package cn

//leetcode submit region begin(Prohibit modification and deletion)
import (
	sort2 "sort"
)

func merge(intervals [][]int) [][]int {
	sort2.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] >= cur[0] {
			prev[1] = max5(prev[1], cur[1])
		} else {
			ans = append(ans, prev)
			prev = cur
		}
	}
	// 注意，需要补偿最后一个数组
	ans = append(ans, prev)
	return ans
}

func max5(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
