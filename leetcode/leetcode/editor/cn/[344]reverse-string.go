package cn

//leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte) {
	lo, hi := 0, len(s) - 1
	for lo < hi {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
