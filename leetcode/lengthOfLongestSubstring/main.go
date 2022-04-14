package main

import "fmt"

func main() {
	s := "abcabcbb" //=> "97 98 99 97 98 99 98 98"
	fmt.Println(lengthOfLongestSubstring(s))
}

// 无重复字符串的最长字串， 难度：中等
// 解法：滑动窗口， 时间复杂度：O(n)
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为-1，相当于我们在字符串左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一格字符
			delete(m, s[i-1])
		}
		// m[s[rk+1]] == 0 表示字符还未在map中出现过
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断的移动右指针
			m[s[rk+1]]++
			fmt.Println(m)
			rk++
		}
		fmt.Println("-----------------")
		// 第i到rk个字符是一个极长的无重复字串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
