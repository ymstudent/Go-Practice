package cqueue

import "errors"

// 利用数组实现一个顺序栈，非并发安全
type ArrayStack struct {
	items []int // 数组
	count int   // 栈中元素个数
	n     int   // 栈的大小
}

func NewArrayStack(n int) *ArrayStack {
	return &ArrayStack{
		items: make([]int, 0, n),
		count: 0,
		n:     n,
	}
}

func (s *ArrayStack) Push(item int) bool {
	if s.count == s.n {
		return false
	}
	s.items[s.count] = item
	s.count++
	return true
}

func (s *ArrayStack) Pop() (int, error) {
	if s.count == 0 {
		return 0, errors.New("stack is empty")
	}
	tmp := s.items[s.count]
	s.count--
	return tmp, nil
}
