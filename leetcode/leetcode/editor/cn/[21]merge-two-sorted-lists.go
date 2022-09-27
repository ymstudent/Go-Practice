package cn
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l1, l2, res := list1, list2, &ListNode{Val: -1}
	tmp := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
			tmp = tmp.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
			tmp = tmp.Next
		}
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return res.Next
}
//leetcode submit region end(Prohibit modification and deletion)
