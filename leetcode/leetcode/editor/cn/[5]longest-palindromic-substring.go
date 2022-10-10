package cn

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	var res string
	// 遍历字符
	for i := 0; i < len(s); i++ {
		// 寻找以s[i]为中心的最长回文子串
		s1 := palindrome(s, i, i)
		// 寻找以s[i] 和 s[i+1]为中心的最长回文子串
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

// 在 s 中寻找以 s[l] 和 s[r] 为中心的最长回文串
// 当 l == r，就相当于寻找长度为奇数的回文串
// 当输入相邻的 l 和 r，则相当于寻找长度为偶数的回文串。
func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

//leetcode submit region end(Prohibit modification and deletion)
