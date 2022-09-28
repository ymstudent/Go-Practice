package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 返回后一个节点不需要判断单双，当为双数时，慢指针刚好到第二个中间节点
	/*if fast == nil { //链表长度为双数，返回第二个中间节点
		return slow.Next
	}*/
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)
