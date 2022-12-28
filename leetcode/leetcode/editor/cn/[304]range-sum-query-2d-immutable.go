package cn

//leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	preSum [][]int
}

func Constructor3(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	numMatrix := new(NumMatrix)
	numMatrix.preSum = make([][]int, m+1)
	numMatrix.preSum[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		if len(numMatrix.preSum[i]) == 0 {
			numMatrix.preSum[i] = make([]int, n+1)
		}
		for j := 1; j <= n; j++ {
			numMatrix.preSum[i][j] = numMatrix.preSum[i-1][j] + numMatrix.preSum[i][j-1] + matrix[i-1][j-1] - numMatrix.preSum[i-1][j-1]
		}
	}
	return *numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row1][col1] + this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)
