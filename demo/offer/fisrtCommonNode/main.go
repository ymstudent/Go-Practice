package fisrtCommonNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 简单： 两个链表的第一个公共节点
// 输入两个无环的单向链表，找出它们的第一个公共结点，如果没有公共节点则返回空。
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	m := make(map[*ListNode]struct{})
	for pHead1 != nil {
		m[pHead1] = struct{}{}
		pHead1 = pHead1.Next
	}
	for pHead2 != nil {
		if _, ok := m[pHead2]; ok {
			return pHead2
		}
		pHead2 = pHead2.Next
	}
	return nil
}

// 双指针解法
// 想办法让 pHead1 与 pHead2长度相等，怎么办？
// 让pHead1+pHead2做为pHead2的新长度，pHead2+pHead1作为pHead1的新长度
func FindFirstCommonNode2(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	l1, l2 := pHead1, pHead2
	for l1 != l2 {
		if l1 == nil {
			l1 = pHead2
		} else {
			l1 = l1.Next
		}

		if l2 == nil {
			l2 = pHead1
		} else {
			l2 = l2.Next
		}
	}
	return l1
}
