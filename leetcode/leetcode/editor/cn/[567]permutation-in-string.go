package cn

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[uint8]int), make(map[uint8]int)
	for i := range s1 {
		need[s1[i]]++
	}

	left, right, valid := 0, 0, 0

	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 窗口收缩的判断条件为，窗口长度大于等于了子串长度
		for right-left >= len(s1) {
			// 满足条件，直接返回true
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
