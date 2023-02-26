package cn

//leetcode submit region begin(Prohibit modification and deletion)

func maxScoreWords(words []string, letters []byte, score []int) int {
	left := [26]int{} // 统计letters中的个字母个数
	for _, c := range letters {
		left[c-'a']++
	}
	var ans int
	var dfs func(i, total int)
	dfs = func(i, total int) {
		if i < 0 { // 终止条件，words遍历完成
			ans = max1255(ans, total) // 记录最大的分数
			return
		}

		// 不选择words[i]
		dfs(i-1, total)

		// 选择words[i]
		for j, c := range words[i] {
			c -= 'a'
			if left[c] == 0 { // 如果left中的剩余字母不够了，需要撤销之前使用的,然后直接返回
				for _, v := range words[i][:j] {
					left[v-'a']++
				}
				return
			}
			left[c]--         // 减少剩余字母
			total += score[c] // 累加得分
		}
		dfs(i-1, total) // 继续判断下一个单词是否要选择

		// 还原现场
		for _, v := range words[i] {
			left[v-'a']++
		}
	}

	dfs(len(words)-1, 0)
	return ans
}

func max1255(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
