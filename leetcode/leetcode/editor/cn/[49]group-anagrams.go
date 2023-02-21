package cn

//leetcode submit region begin(Prohibit modification and deletion)
import (
	sort2 "sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		// 本质上就是把一个字符串转化为唯一的hash标识
		s := []byte(str)
		sort2.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}
	ans := make([][]string, 0)
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
