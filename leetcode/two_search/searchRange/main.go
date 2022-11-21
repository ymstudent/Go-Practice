package main

import "fmt"

func main() {
	//println(searchRange([]int{1, 1, 2}, 1))
	println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 8))
}

//34：在排序数组中查找元素的第一个和最后一个位置
//地址：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
//难度：中等
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			//等于的时候就需要左右边界了，但是这样写最坏的情况下需要遍历整个数组，这样就变成O(n)了
			st, ed := mid, mid
			for i := mid - 1; i >= 0; i-- {
				if nums[i] != target {
					break
				}
				st = i
			}

			for i := mid + 1; i < len(nums); i++ {
				if nums[i] != target {
					break
				}
				ed = i
			}
			return []int{st, ed}
		}
	}

	return []int{-1, -1}
}

// 这道题起始可以拆分成在数组中寻找第一个大于等于target的元素(左边界)和第一个大于target的元素(右边界)
func searchRange2(nums []int, target int) []int {
	left := searchLeft(nums, target)
	fmt.Println("======================")
	right := searchRight(nums, target)
	println(left)
	println(right)
	//println(binarySearch(nums, target, true))
	//println(binarySearch(nums, target, false))
	if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		/*if nums[mid] == target {
			right = mid - 1 // note: 收紧右侧边界以锁定左侧边界
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}*/
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		fmt.Printf("mid=%d\n", mid)
		fmt.Printf("left=%d\n", left)
		fmt.Printf("right=%d\n", right)
	}
	return left
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		/*if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}*/
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		fmt.Printf("mid=%d\n", mid)
		fmt.Printf("left=%d\n", left)
		fmt.Printf("right=%d\n", right)
	}
	return right
}

func binarySearch(nums []int, target int, lower bool) int {
	left, right, ans := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target || (lower && nums[mid] >= target) {
			right = mid - 1
			ans = right
		} else {
			left = mid + 1
		}
	}
	return ans
}
