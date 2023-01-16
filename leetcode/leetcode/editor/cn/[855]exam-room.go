package cn

//leetcode submit region begin(Prohibit modification and deletion)
type ExamRoom struct {
	n int
	s []int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{n: n, s: []int{}}
}

func (this *ExamRoom) Seat() int {
	student := 0
	idx := 0
	if len(this.s) > 0 {
		// 二分法寻找最合适的位置
		dist := this.s[0]
		pre := -1
		for i, v := range this.s {
			if pre != -1 {
				d := (v - pre) / 2
				if d > dist {
					dist = d
					student = pre + d
					idx = i
				}
			}
			pre = v
		}
		if this.n-1-this.s[len(this.s)-1] > dist {
			student = this.n - 1
			idx = len(this.s)
		}
	}
	this.s = append(this.s, 0)
	copy(this.s[idx+1:], this.s[idx:])
	this.s[idx] = student
	return student
}

func (this *ExamRoom) Leave(p int) {
	idx := 0
	for i := 0; i < len(this.s); i++ {
		if this.s[i] == p {
			idx = i
			break
		}
	}
	this.s = append(this.s[:idx], this.s[idx+1:]...)
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
//leetcode submit region end(Prohibit modification and deletion)
