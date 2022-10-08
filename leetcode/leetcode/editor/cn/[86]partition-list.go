package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2  := &ListNode{Val: -1},  &ListNode{Val: -1}
	minList, maxList := dummy1, dummy2
	for head != nil {
		if head.Val >= x {
			maxList.Next = head//&ListNode{Val: head.Val}
			maxList = maxList.Next
		} else {
			minList.Next = head//&ListNode{Val: head.Val}
			minList = minList.Next
		}
		head = head.Next
	}
	maxList.Next = nil
	minList.Next = dummy2.Next
	return dummy1.Next
}

//leetcode submit region end(Prohibit modification and deletion)
