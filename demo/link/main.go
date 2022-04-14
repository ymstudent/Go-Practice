package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

//206 简单题：单链表的反转
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	/*newHead := &ListNode{Next: nil}

	for head.Next != nil {
		tmp := head
		head = head.Next
		tmp.Next = newHead
		newHead = tmp
	}*/

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

//141 简单题：链表中的环检测
func hasCycle(head *ListNode) bool {
	/*
		lmap := map[*ListNode]struct{}{}

		// 注意这里的判断条件是head != nil; 如果写成head.Next != nil, 当传入的链表为空时，会报空指针错误
		//for head.Next != nil {
		for head != nil {
			if _, ok := lmap[head]; ok {
				return true
			}
			lmap[head] = struct{}{}
			head = head.Next
		}
	*/
	if head == nil || head.Next == nil {
		return false
	}
	//利用快慢指针解决：
	slow, fast := head, head.Next
	// 这里不能用head.Next != nil，因为当链表只有2个元素或者无环时，fast.next.next会报空指针
	/*for head.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}*/

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

//142 中等：链表中的环检测2， 需要返回入环的第一个节点
func detectCycle(head *ListNode) *ListNode {
	// 设链表共有a + b个节点，其中从链表头部到环入口有a个节点
	// 当快慢指针第一次相遇时，快指针走的步数是慢指针的2倍：f = 2s
	// fast 比 slow 多走n个环的长度，即：f = s + nb
	// 两式相减得：f = 2nb，s = nb， 此时快慢指针可能指向环内任意一个节点
	// 那么怎么才能找到环的入口呢？如果让一个指针从链表头部一直向前走并统计步数k,
	// 那么所有走到链表环入口节点的步数是：k = a + nb
	// 知道以上2点知识后，我们就知道，只要让 slow 指针再往前走 a 步就可以到环的入口了
	// 但我们不知道 a 的值，该怎么办？依然是使用双指针法。我们构建一个指针，此指针需要有以下性质：
	// 此指针和slow 一起向前走 a 步后，两者在入口节点重合。那么从哪里走到入口节点需要 a 步？答案是链表头部head。
	// 所以，再快慢指针第一次相遇后，让 slow 指针的位置保持不变，fast 指针重新指向链表头部，
	// 当两个指针再次相遇时，所在的位置就是环的入口
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow { //第一次相遇：
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

//21 简单题 合并2个有序链表：将两个升序链表合并为一个新的升序链表并返回
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{Next: nil} //添加一个哨兵节点
	result := newHead
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			newHead.Next = l2
			l2 = l2.Next
		} else {
			newHead.Next = l1
			l1 = l1.Next
		}
		newHead = newHead.Next
	}
	if l1 == nil {
		newHead.Next = l2
	}
	if l2 == nil {
		newHead.Next = l1
	}

	return result.Next
}

// 递归解决合并有序链表问题
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}

//19 中等难度 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}

//876 简单题 链表的中间结点
// ps 链表长度为偶数，要返回中间节点的前一个，可以增加一个哨兵节点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
