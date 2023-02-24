package cn

//leetcode submit region begin(Prohibit modification and deletion)

type pair struct {
	x, y int
}

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	// 记录访问过的元素
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 减枝
			return false
		}
		if k == len(word)-1 {
			return true
		}

		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 退回上一步

		for _, dir := range dirs { // 上下左右探索，注意需要跳过已经访问的单元格
			newI, newJ := i+dir.x, j+dir.y
			if 0 <= newI && newI < m && 0 <= newJ && newJ < n && !vis[newI][newJ] && check(newI, newJ, k+1) {
				return true
			}
		}
		return false
	}

	for i, rows := range board {
		for j := range rows { // 以每个单元格为启动，判断是否有满足条件的
			if check(i, j, 0) {
				return true
			}
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
