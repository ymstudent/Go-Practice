package cn

//leetcode submit region begin(Prohibit modification and deletion)
// 式子变换：nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
// 可以转换成：nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
// 设：f(i) = nums[i] - rev(nums[i])
// 那么就存在：f(i) = f(j)
// 用一个哈希表记录nums中每个元素nums[i] - rev(nums[i])的值
// nums[i] - rev(nums[i])出现几次就表示能结成几对
func countNicePairs(nums []int) int {
	var ans int
	m := make(map[int]int)
	for _, num := range nums {
		rev := revNum(num)
		// 获取之前哈希表中num-rev出现的次数
		ans += m[num-rev]
		m[num-rev]++
	}
	return ans % (1e9 + 7)
}

// 反转整数：leetcode第7题，medium
func revNum(x int) int {
	var rev int
	for x != 0 {
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return rev
}

//leetcode submit region end(Prohibit modification and deletion)
