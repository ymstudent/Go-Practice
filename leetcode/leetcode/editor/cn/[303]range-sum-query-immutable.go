package cn

//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	Arr []int
}

func Constructor2(nums []int) NumArray {
	numArr := new(NumArray)
	numArr.Arr = make([]int, len(nums)+1) // ps1: 长度需要比原数组长多1位，首位用0填充，相当于链表的哑节点
	for i := 1; i < len(numArr.Arr); i++ {
		numArr.Arr[i] = numArr.Arr[i-1] + nums[i-1]
	}
	return *numArr
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Arr[right+1] - this.Arr[left] // ps2: 用right + 1 -left
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
