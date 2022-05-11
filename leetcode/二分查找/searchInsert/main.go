package main

func main() {
	println(searchInsert([]int{1, 3, 5, 6}, 2))
}

//35：搜索插入位置
//地址：https://leetcode.cn/problems/search-insert-position/
//难度：简单
func searchInsert(nums []int, target int) int {
	ans, left, right := 0, 0, len(nums)-1 //区间左闭右闭, [left, right]
	// 利用二分查找寻找第一个小于等于target的下标
	for left <= right { //左闭右闭区间对应的条件就是 <=
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			ans = mid
			right = mid - 1
		}
	}
	return ans
}
