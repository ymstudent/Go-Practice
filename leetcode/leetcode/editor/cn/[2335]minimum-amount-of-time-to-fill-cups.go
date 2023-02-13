package cn

//leetcode submit region begin(Prohibit modification and deletion)
import sort2 "sort"

// 数学推导
// 设3种杯子的数量为 x,y,z, x <= y <= z
// 如果 x+y <= z，装满z时，x,y一定可以装满，总耗时为 z
// 如果 x+y > z, 令 t = x + y - z
// 如果t为偶数，那么先将x,y各装t/2杯，耗时 t/2, 剩下的x+y=z，装完需要耗时 z
// 所以一共耗时 t/2 + z = (x+y-z) / 2 + z = (x+y+z) / 2
// 如果t为奇数，那么先将x,y各装(t-1)/2杯，耗时 (t-1)/2，剩下的x+y=z+1,需要耗时 z+1
// 所以一共耗时 (t-1)/2 + z + 1 = (x+y-z-1)/2 + z + 1= (x+y+z+1)/2
func fillCups(amount []int) int {
	sort2.Ints(amount)
	if amount[0]+amount[1] <= amount[2] {
		return amount[2]
	}
	// 这里为什么不用分情况判断呢？
	// 因为 t 如果为偶数，那么 x+y+z也为偶数，golang 整数相除，向下取整， 1/2 会被忽略
	return (amount[0] + amount[1] + amount[2] + 1) / 2
}

//leetcode submit region end(Prohibit modification and deletion)
