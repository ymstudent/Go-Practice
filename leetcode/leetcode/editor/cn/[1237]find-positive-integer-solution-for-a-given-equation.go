package cn

import (
	sort2 "sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */
// 1、暴力解法，枚举
func findSolution(customFunction func(int, int) int, z int) [][]int {
	var ans [][]int
	for x := 1; x <= 1000; x++ {
		for y := 1; y <= 1000; y++ {
			tmp := customFunction(x, y)
			if tmp == z {
				ans = append(ans, []int{x, y})
			}
			if tmp > z {
				break
			}
		}
	}
	return ans
}

// 2 二分法, 枚举x，然后在[1, 1000]中枚举y，使得customFunction(x,y) == z
func findSolution2(customFunction func(int, int) int, z int) [][]int {
	var ans [][]int
	for x := 1; x <= 1000; x++ {
		y := 1 + sort2.Search(999, func(y int) bool {
			return customFunction(x, y+1) > z
		})
		if customFunction(x, y) == z {
			ans = append(ans, []int{x, y})
		}
	}
	return ans
}

// 3 同向双指针法，定义指针x,y，初始时令 x = 1，y = 1000
// 如果 f(x, y) == z，将x,y 加入结果
// 如果 f(x, y) > z, 此时如果继续增大x，则 f(x+1, y) 也必定大于 z, 所以只能在x不变的情况下，将y减小
// 如果 f(x, y) < z，同理，只能在y不变的情况下，将x增大
func findSolution3(customFunction func(int, int) int, z int) [][]int {
	var ans [][]int
	x, y := 1, 1000
	for x <= 1000 && y > 0 {
		tmp := customFunction(x, y)
		if tmp == z {
			ans = append(ans, []int{x, y})
			x++
			y--
		} else if tmp > z {
			y--
		} else {
			x++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
