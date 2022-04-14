package cqueue

import "errors"

// 数组实现的顺序队列，非并发安全
type ArrayQueue struct {
	items []int
	n     int
	head  int
	tail  int
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		items: make([]int, 0, capacity),
		n:     capacity,
	}
}

func (q *ArrayQueue) Enqueue(item int) bool {
	if q.tail == q.n {
		return false
	}
	q.items[q.tail] = item
	q.tail++
	return true
}

func (q *ArrayQueue) Dequeue() (int, error) {
	if q.tail == 0 {
		return 0, errors.New("queue is empty")
	}
	tmp := q.items[q.head]
	q.head--
	return tmp, nil
}
