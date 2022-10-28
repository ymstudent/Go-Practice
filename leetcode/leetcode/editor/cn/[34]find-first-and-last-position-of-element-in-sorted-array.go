package cn

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	/*left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			// 锁定左侧边界
			right = mid - 1
		}
	}
	if left == len(nums) {
		return []int{-1, -1}
	}
	if nums[left] == target {
		for i := left; i < len(nums); i++ {
			if nums[i] == target {
				right = i
			}
		}
		return []int{left, right}
	}
	return []int{-1, -1}*/
	return []int{left_bound(nums, target), right_bound(nums, target)}
}

func left_bound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			// 锁定左侧边界
			right = mid - 1
		}
	}
	// 防止越界
	if left == len(nums) {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func right_bound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			// 锁定右侧边界
			left = mid + 1
		}
	}
	// 防止left-1数组越界
	if left-1 < 0 {
		return -1
	}
	if nums[left-1] == target {
		return left - 1
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
