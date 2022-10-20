package cn

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	n := len(matrix)
	// 先沿对角线反转
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		matrix[i] = reverseSlice(matrix[i])
	}
}

func reverseSlice(a []int) []int {
	i, j := 0, len(a)-1
	for i < j { // 注意：这里是小于，当出现等于时会发生什么？
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
