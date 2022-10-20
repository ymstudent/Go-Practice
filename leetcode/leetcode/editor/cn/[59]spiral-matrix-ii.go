package cn

//leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	upBound, lowBound, leftBound, rightBound := 0, n-1, 0, n-1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	num := 0
	for num < n*n { // 注意：只能是小于，如果有等于符合会陷入死循环
		if upBound <= lowBound {
			// 填入顶部
			for j := leftBound; j <= rightBound; j++ {
				num++
				res[upBound][j] = num
			}
			upBound++
		}

		if leftBound <= rightBound {
			// 填入右边
			for i := upBound; i <= lowBound; i++ {
				num++
				res[i][rightBound] = num
			}
			rightBound--
		}
		if upBound <= lowBound {
			// 填入底部
			for j := rightBound; j >= leftBound; j-- {
				num++
				res[lowBound][j] = num
			}
			lowBound--
		}
		if leftBound <= rightBound {
			// 填入左边
			for i := lowBound; i >= upBound; i-- {
				num++
				res[i][leftBound] = num
			}
			leftBound++
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
