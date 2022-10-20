package cn

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	upBound, lowBound, leftBound, rightBound := 0, m-1, 0, n-1

	var res []int
	for len(res) < m*n {
		// 在顶部从左向右遍历
		if upBound <= lowBound {
			for j := leftBound; j <= rightBound; j++ {
				res = append(res, matrix[upBound][j])
			}
			upBound++
		}
		// 在右边从上到下遍历
		if leftBound <= rightBound {
			for i := upBound; i <= lowBound; i++ {
				res = append(res, matrix[i][rightBound])
			}
			rightBound--
		}
		// 在下边从右向左遍历
		if upBound <= lowBound {
			for j := rightBound; j >= leftBound; j-- {
				res = append(res, matrix[lowBound][j])
			}
			lowBound--
		}
		// 在左边从下到上遍历
		if leftBound <= rightBound {
			for i := lowBound; i >= upBound; i-- {
				res = append(res, matrix[i][leftBound])
			}
			leftBound++
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
