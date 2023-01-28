package cn

//leetcode submit region begin(Prohibit modification and deletion)
func waysToMakeFair(nums []int) int {
	var odd1, even1, odd2, even2 int
	for i, num := range nums {
		if i&1 > 0 { // 奇数和
			odd2 += num
		} else { // 偶数和
			even2 += num
		}
	}
	var ans int
	for i, num := range nums {
		if i&1 > 0 {
			odd2 -= num
		} else {
			even2 -= num
		}
		if odd1+even2 == even1+odd2 {
			ans++
		}
		// 注意这里，不能合并到前面去
		if i&1 > 0 {
			odd1 += num
		} else {
			even1 += num
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
