package cn

//leetcode submit region begin(Prohibit modification and deletion)
func shipWithinDays(weights []int, days int) int {
	left, right := 0, 1
	for _, v := range weights {
		if v > left {
			left = v
		}
		right += v
	}

	for left < right {
		mid := left + (right-left)>>1
		if f(weights, mid) <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func f(weights []int, x int) int {
	days := 0
	for i := 0; i < len(weights); {
		c := x
		for i < len(weights) {
			if c < weights[i] {
				break
			} else {
				c -= weights[i]
			}
			i++
		}
		days++
	}
	return days
}

//leetcode submit region end(Prohibit modification and deletion)
