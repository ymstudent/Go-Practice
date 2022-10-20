package cn

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	var sb []byte
	// 清洗数据，删除多余的空格
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			// 单词中的字母/数字
			sb = append(sb, s[i])
		} else if len(sb) != 0 && sb[len(sb)-1] != ' ' {
			// 单词之间只保留一个空格
			sb = append(sb, ' ')
		}
	}
	// 末尾如果有空格，清除
	if sb[len(sb)-1] == ' ' {
		sb = sb[:len(sb)-1]
	}

	n := len(sb)
	// 先整体翻转
	reverseByte(sb, 0, n-1)
	// 再把每个单词翻转
	for i := 0; i < n; { // 注意：这里没有i++
		for j := i; j < n; j++ {
			if j+1 == n || sb[j+1] == ' ' {
				// sb[i...j] 是一个单词，翻转
				reverseByte(sb, i, j)
				// 把i置为下一个单词的首字母
				i = j + 2
				break
			}
		}
	}
	return string(sb)
}

func reverseByte(a []byte, i, j int) {
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
