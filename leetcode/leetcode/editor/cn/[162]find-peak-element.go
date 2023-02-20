package cn

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
// 推导：
// 爬坡法：在[0,n)中随机选取一个数i，对于nums[i]存在以下几种情况
// 1、nums[i] > nums[i-1] && nums[i] > nums[i+1]，那么i就是我们要找的位置，直接返回
// 2、nums[i-1] < nums[i] < nums[i+1]，这个时候我们处于升序，即上坡中，需要往右走，即 i = i+1
// 3、nums[i-1] > nums[i] > nums[i+1]，这个时候我们处于降序，即下坡中，需要往左走，即 i = i-1
// 4、nums[i] < nums[i-1] && nums[i] < nums[i+1]，这个时候处于谷底，往哪走都行，我们规定一下，这个时候往右走，即 i=i+1
// 综上：当i不是峰值时，nums[i] < nums[i+1]，我们往右走，nums[i] > nums[i+1]，我们往左走
// 根据上面的推导写出解法1
func findPeakElement1(nums []int) int {
	n := len(nums)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, n-1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

// 推导2：
// 当 nums[i] < nums[i+1]，我们往右走到i+1后，位置i左侧的所有位置我们都不可能在后续迭代中走到
// 因为nums[i] < nums[i+1]时，我们没办法向左走
// 所以当 nums[i] < nums[i+1] 时，我们抛弃[l,i]范围，在剩余的[i+1,r]中寻找峰值
// 当 nums[i] > nums[i+1] 时, 我们抛弃[i,r]范围，在剩余的[l, i-1]中寻找
func findPeakElement(nums []int) int {
	n := len(nums)

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	l, r := 0, n-1
	// 注意，这里的二分法没有终止判断，因为肯定有峰值存在
	for {
		mid := l + (r-l)/2
		if get(mid) > get(mid+1) && get(mid) > get(mid-1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
