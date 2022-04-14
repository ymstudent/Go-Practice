package main

func main() {

}

// 二分查找的局限性：1、依赖顺序表结构(数组)，2、针对的是有序数据
// 适用于数据量较大，插入，删除操作不频繁，一次排序，多次查找的场景
// 二分查找最简单的实现：在一个有序的不含重复元素的数组中查找一个给定的值
func bsearch(a []int, b int) int {
	low, high := 0, len(a)-1

	for low <= high { //注意点1： 循环推出条件为：low <= high, 而不是low < high
		//注意点2：(low + high)/2存在溢出风险，改为方法为：low + (high - low) / 2
		// 更好的做法为：low + ((high - low)>>1), 因为相比除法，计算机处理位运算要快得多
		mid := (low + high) / 2
		if a[mid] == b {
			return mid
		} else if a[mid] < b {
			//注意点3：如果写成low = mid 或high = mid，可能存在死循环的风险
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// 查找第一个值等于给定值的元素
func bsearch1(a []int, v int) int {
	lo, hi := 0, len(a)-1

	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if a[mid] > v {
			hi = mid - 1
		} else if a[mid] < v {
			lo = mid + 1
		} else {
			// 如果是第一元素，或者前一个元素不等于我们要找的，那么mid就是第一个等于给定值的元素，否则继续二分
			if mid == 0 || a[mid-1] != v {
				return mid
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

// 查找最后一个值等于给定值的元素
func bsearch2(a []int, v int) int {
	l := len(a)
	lo, hi := 0, l-1

	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if a[mid] > v {
			hi = mid - 1
		} else if a[mid] < v {
			lo = mid + 1
		} else {
			if mid == l-1 || a[mid+1] != v {
				return mid
			} else {
				lo = mid + 1
			}
		}
	}

	return -1
}

// 查找第一个大于等于给定值的元素
func bsearch3(a []int, v int) int {
	l := len(a)
	lo, hi := 0, l-1

	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if a[mid] >= v {
			if mid == 0 || a[mid-1] < v {
				return mid
			}
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}

// leetcode 33：如果有序数组是一个循环有序数组，比如 4，5，6，1，2，3。求目标值的位置
func search(a []int, target int) int {
	low, high := 0, len(a)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] == target {
			return mid
		}
		// 如果low <= mid, 表示前半部分是个有序数组
		if a[mid] >= a[low] {
			// 判断元素是否在有序数组中，如果在就从有序数组中查找，否就去后半部分继续二分
			if a[mid] > target && a[low] <= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// low > mid 表示后半部分是个有序数组，判断是否在有序数组中，不在去前半部分继续二分
			if a[mid] < target && a[high] >= target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
