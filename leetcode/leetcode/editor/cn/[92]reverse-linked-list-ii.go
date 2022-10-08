package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// base case
	if left == 1 {
		return reverseN(head, right)
	}
	// 敲进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 反转链表前n个节点
var successor *ListNode //后置节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next //记录第n+1个节点
		return head
	}
	tmp := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return tmp
}

func reverseN2(a, b *ListNode) *ListNode {
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
