package cn

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
// 超时，去球，不看了，太难了
type MKAverage struct {
	m    int
	k    int
	nums []int
}

func Constructor5(m int, k int) MKAverage {
	return MKAverage{m: m, k: k}
}

func (this *MKAverage) AddElement(num int) {
	this.nums = append(this.nums, num)
}

func (this *MKAverage) CalculateMKAverage() int {
	if len(this.nums) < this.m {
		return -1
	}
	tmp := append([]int{}, this.nums[(len(this.nums)-this.m):]...)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	tmp = tmp[this.k : len(tmp)-this.k]
	var sum int
	for _, v := range tmp {
		sum += v
	}
	return sum / len(tmp)
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
//leetcode submit region end(Prohibit modification and deletion)
