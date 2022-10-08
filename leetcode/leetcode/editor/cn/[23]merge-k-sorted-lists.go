package cn

import "container/heap"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 使用内置的最小堆实现优先级队列
type minHeap []*ListNode
func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	// 初始化优先级队列
	priorityQueue := new(minHeap)
	// 将k个链表的头节点加入最小堆
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(priorityQueue, lists[i])
		}
	}

	dummy := new(ListNode)
	p := dummy
	for priorityQueue.Len() > 0 {
		// 获取最小节点，接到结果链表中
		tmp := heap.Pop(priorityQueue).(*ListNode)
		if tmp.Next != nil {
			heap.Push(priorityQueue, tmp.Next)
		}
		p.Next = tmp
		// 指针不断前进
		p = p.Next
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
