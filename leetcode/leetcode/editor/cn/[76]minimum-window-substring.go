package cn

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	need, window := make(map[uint8]int), make(map[uint8]int)
	// 构建need
	for i := range t {
		need[t[i]]++
	}
	left, right, valid := 0, 0, 0
	// 记录最小覆盖子串的起始索引及长度
	start, size := 0, math.MaxInt
	for right < len(s) {
		// c是将移入窗口的字符
		c := s[right]
		// 扩大窗口
		right++
		// 进行窗口内的数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			// 某个字符在 window 的数量满足了 need 的需要，就要更新 valid，表示有一个字符已经满足要求
			if window[c] == need[c] {
				valid++
			}
		}
		/*** debug 输出的位置 ***/
		//fmt.Printf("window: [%d, %d)\n", left, right)
		/********************/
		// 判断左侧窗口是否要收缩，当valid与need长度相等时，表示已经完全覆盖了串 t
		for valid == len(need) {
			// 更新最小覆盖子串
			if right-left < size {
				start = left
				size = right - left
			}
			// d是将移除窗口的字符
			d := s[left]
			// 缩小窗口
			left++
			// 进行窗口内一系列数据更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if size == math.MaxInt {
		return ""
	} else {
		// 返回最小覆盖子串
		return s[start : start+size]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
