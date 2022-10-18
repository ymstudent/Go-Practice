package cn

//leetcode submit region begin(Prohibit modification and deletion)

type Difference struct {
	diff []int
}

// Difference 构造差分数组，区间操作将在这个数组上进行
func (d *Difference) Difference(nums []int) {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	d.diff = diff
}

// 给闭区间 [i, j] 增加 val (可以是负数)
func (d *Difference) increment(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

// 返回结果数组
func (d *Difference) result() []int {
	// 根据差分数据反推出原始数组
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

func carPooling(trips [][]int, capacity int) bool {
	// 根据提示，最多有1000个车站，所以数组长度定位1001
	nums := make([]int, 1001)
	// 构造差分解法
	df := new(Difference)
	df.Difference(nums)

	for _, v := range trips {
		// 乘客数量
		val := v[0]
		// 第 v[1] 站乘客上车
		i := v[1]
		// 第 v[2] 站乘客已下车
		// 即乘客在车上的区间是[v[1], v[2] - 1]
		j := v[2] - 1
		// 进行区间操作
		df.increment(i, j, val)
	}

	res := df.result()

	// 客车自始自终都不应该超载
	for i := 0; i < len(res); i++ {
		if capacity < res[i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
