package cqueue

type CQueue struct {
	inStack  []int //输入栈
	outStack []int //输出栈
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	// 将元素添加到输入栈尾部
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		if len(this.inStack) == 0 {
			return -1
		}
		this.in2out()
	}
	value := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return value
}

func (this *CQueue) in2out() {
	// 将输入栈的元素依次压入输出栈，这样输出栈的栈顶往栈底的顺序就是队列队首往队尾的顺序
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}
