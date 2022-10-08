package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	slow, fast := dummy, dummy
	for fast.Next != nil {
		fast = fast.Next
		if n > 0 {
			n--
		} else {
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
