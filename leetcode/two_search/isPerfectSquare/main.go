package main

func main() {

}

// 有效的完全平方数
// 地址：https://leetcode.cn/problems/valid-perfect-square/
// 难度：简单
func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	left, right := 0, num/2
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
