package cn

//leetcode submit region begin(Prohibit modification and deletion)
func decodeMessage(key string, message string) string {
	cur := byte('a')
	rules := make(map[rune]byte)
	for _, v := range key {
		if v != ' ' && rules[v] == 0 {
			rules[v] = cur
			cur++
		}
	}
	m := []byte(message)
	for i, v := range message {
		if v != ' ' {
			m[i] = rules[v]
		}
	}
	return string(m)
}

//leetcode submit region end(Prohibit modification and deletion)
