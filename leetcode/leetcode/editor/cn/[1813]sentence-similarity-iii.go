package cn

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1 := strings.Split(sentence1, " ")
	s2 := strings.Split(sentence2, " ")
	i, n := 0, len(s1)
	j, m := 0, len(s2)
	// 开头有多少相同字符
	for i < n && i < m && s1[i] == s2[i] {
		i++
	}
	// 结尾有多少相同字符
	for j < n-i && j < m-i && s1[n-1-j] == s2[m-1-j] {
		j++
	}
	// 如果开头+结尾的相同字符数等于较短的句子长度，那么相似
	return i+j == min(n, m)
}

func min(n, m int) int {
	if n > m {
		return m
	}
	return n
}

//leetcode submit region end(Prohibit modification and deletion)
