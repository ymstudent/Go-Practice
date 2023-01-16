package cn

//leetcode submit region begin(Prohibit modification and deletion)
func countDifferentSubsequenceGCDs(nums []int) int {
	mx := maxInNums(nums)
	has := getHasSlice(mx, nums)
	var ans int
	for i := 1; i <= mx; i++ {
		g := 0  // 0 和任何数 x 的最大公约数都是 x
		for j := i; j <= mx && g != i; j += i { // 枚举 i 的倍数 j
			if has[j] { // 如果 j 在 nums 中
				g = gcd(g, j) // 更新最大公约数
			}
		}
		if g == i { // 找到一个答案
			ans++
		}
	}
	return ans
}

func maxInNums(nums []int) int {
	var ans int
	for _, v := range nums {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func getHasSlice(mx int, nums []int) []bool {
	has := make([]bool, mx+1)
	for _, v := range nums {
		has[v] = true
	}
	return has
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
