package cn
//leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	L, R := make([]int, n), make([]int, n)
	L[0] = 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	R[n-1] = 1
	for i := n-2; i>=0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = L[i] * R[i]
	}
	return res
}

// 空间复杂度O(1)的写法
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	L := make([]int, n)
	L[0] = 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	R := 1
	for i := n-1; i>=0; i-- {
		res[i] = L[i] * R
		R *= nums[i]
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
