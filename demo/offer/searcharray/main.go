package main

func main() {

}

// 二分查找几个注意点：
// 1、推出循环条件：low <= high
// 2、mid 的取值：(low + high) / 2可能出现和溢出的情况，建议使用：low + ((high - low) >> 1)
// 3、low和high的更新：low = mid + 1; high = mid - 1

// 33，中等，搜索旋转排序数组
func search(a []int, target int) int {
	low, high := 0, len(a)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] == target {
			return mid
		}

		if a[mid] >= a[low] {
			if a[mid] > target && a[low] <= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if a[mid] < target && a[high] >= target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

//153, 中等, 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + ((high - low) >> 1)
		if nums[mid] == nums[high] {
			return nums[mid]
		} else if nums[mid] < nums[high] {
			// high = mid - 1
			// 注意：这里不能再套用mid -1了，错误用例：[4,5,6,7,0,1,2]
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return nums[low]
}

//154, 困难, 寻找旋转排序数组中的最小值，注意：数组中存在重复元素
func finMin2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + ((high - low) >> 1)
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return nums[low]
}
