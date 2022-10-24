package cn

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	window := make(map[uint8]int)
	left, right, res := 0, 0, 0

	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		// 收缩条件为当前字符在窗口内的数量大于1
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
