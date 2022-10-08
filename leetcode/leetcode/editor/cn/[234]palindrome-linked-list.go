package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// 求中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	// 反转后半部分链表
	slowR := reverse2(slow)
	// 遍历反转后的链表，并与前半部分对比
	newL := head
	for slowR != nil {
		if newL.Val != slowR.Val {
			return false
		}
		slowR = slowR.Next
		newL = newL.Next
	}
	return true
}

func reverse2(a *ListNode) *ListNode {
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

//leetcode submit region end(Prohibit modification and deletion)
