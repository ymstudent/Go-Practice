package cn

//leetcode submit region begin(Prohibit modification and deletion)
func isValidSudoku(board [][]byte) bool {
	var rows [9][9]int
	var cols [9][9]int
	var subBoxs [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			cols[j][index]++
			subBoxs[i/3][j/3][index]++
			if rows[i][index] > 1 || cols[j][index]> 1 || subBoxs[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
