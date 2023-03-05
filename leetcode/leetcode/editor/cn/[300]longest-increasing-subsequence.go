package cn

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max300(dp[i], dp[j]+1)
			}
		}
		ans = max300(ans, dp[i])
	}
	return ans
}

func max300(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
