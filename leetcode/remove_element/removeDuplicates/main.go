package main

func main() {

}

//删除有序数组中的重复项
//地址：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
//难度：简单
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	slow := 0
	for fast := 1; fast < length; fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = fast
		}
	}
	return slow + 1
}
