package main

func main() {
	sortedSquares([]int{-4, -1, 0, 3, 10})
}

//有序数组的平方
//地址：https://leetcode.cn/problems/squares-of-a-sorted-array/submissions/
//难度：简单
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare >= rightSquare {
			result[i] = leftSquare
			left++
		} else {
			result[i] = rightSquare
			right--
		}
	}
	return result
}
