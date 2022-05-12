package main

import "fmt"

func main() {
	println(mySqrt(63))
}

// x 的平方根
// 地址：https://leetcode.cn/problems/sqrtx/
// 难度：简单
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 0, x/2
	ans := -1
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			ans = mid //todo 为什么是mid，而不是left
			left = mid + 1
		}
		fmt.Printf("mid=%d\n", mid)
		fmt.Printf("ans=%d\n", ans)
		fmt.Printf("left=%d\n", left)
		fmt.Printf("right=%d\n", right)
	}
	return ans
}
