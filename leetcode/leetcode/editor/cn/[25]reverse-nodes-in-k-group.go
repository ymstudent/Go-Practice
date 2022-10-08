package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 区间[a, b)内包含k个待反转元素
	a, b := head, head
	for i := 0; i < k; i++ {
		// base case: 不足k个，不需要反转
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 反转前k个元素
	newHead := reverseBetween2(a, b)
	// 递归反转后续链表并连接起来
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 反转链表
func reverse(a *ListNode) *ListNode {
	var pre *ListNode = nil
	cur, next := a, a
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 反转[a, b) 区间内的元素，左闭右开
func reverseBetween2(a, b *ListNode) *ListNode {
	var pre *ListNode = nil
	cur, next := a, a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)
