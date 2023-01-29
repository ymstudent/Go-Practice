package cn

//leetcode submit region begin(Prohibit modification and deletion)
func countAsterisks(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			for i = i + 1; i < len(s); i++ {
				if s[i] == '|' {
					break
				}
			}
		}
		if s[i] == '*' {
			ans++
		}
	}
	return ans
}

// 官方解
func countAsterisks2(s string) int {
	var ans int
	valid := true
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			valid = !valid
		}
		if s[i] == '*' && valid {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
