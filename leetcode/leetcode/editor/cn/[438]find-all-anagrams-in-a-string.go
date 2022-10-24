package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	var res []int
	need, window := make(map[uint8]int), make(map[uint8]int)
	for i := range p {
		need[p[i]]++
	}

	left, right, valid := 0, 0, 0

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
