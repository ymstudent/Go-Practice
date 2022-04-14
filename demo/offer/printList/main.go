package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 简单：从尾到头打印链表
// 输入一个链表的头节点，按链表从尾到头的顺序返回每个节点的值（用数组返回）。
func printListFromTailToHead(head *ListNode) []int {
	var res []int
	if head == nil {
		return res
	}

	return reverseLink(head, res)
}

func reverseLink(head *ListNode, res []int) []int {
	if head.Next == nil {
		return append(res, head.Val)
	}
	tmp := reverseLink(head.Next, res)
	return append(tmp, head.Val)
}
