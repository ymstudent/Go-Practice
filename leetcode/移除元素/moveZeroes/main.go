package main

func main() {
	moveZeroes2([]int{1, 0, 1, 0, 1})
}

//移动0
//地址：https://leetcode.cn/problems/move-zeroes/
//难度：简单
func moveZeroes(nums []int) {
	length := len(nums)
	if length < 2 {
		return
	}

	slow := 0
	for fast := 1; fast < length; fast++ {
		if nums[slow] != 0 {
			slow++
		}
		if nums[slow] == 0 && nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}

	return
}

func moveZeroes2(nums []int) {
	length := len(nums)
	if length < 2 {
		return
	}

	slow := 0
	for fast := 0; fast < length; fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	return
}
