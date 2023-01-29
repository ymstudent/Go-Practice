package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 归并排序
func sortList2(head *ListNode) *ListNode {
	return sort(head, nil)
}

// 自顶向下的归并排序，空间复杂度 O(log n)，时间复杂度 O(n log n)
// 可以参考 labuladong 的归并排序详解 ：https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-66994/gui-bing-p-1387f/
func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	// 利用快慢指针，寻找链表中点
	slow, fast := head, head
	if fast != tail {
		slow = slow.Next
		// 注意这里不能直接写成 fast = fast.Next.Next，防止超过 tail，需要每进一步都判断一下
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	// 递归排序，并合并
	return merge(sort(head, mid), sort(mid, tail))
}

// 21 合并2个有序链表 easy
func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

// 自顶向下的归并排序竟然会超时 !!!
// 自低向上的归并排序不会超时，时间复杂度：O(n log n), 空间复杂度：O(1)
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 1、遍历链表，统计链表长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	// 2、初始化 dummyHead
	dummyHead := &ListNode{Next: head}
	// 3、每次将链表拆分为若干个长度为subLength的子链表，并按照每2个子链表一组进行合并
	for subLength := 1; subLength < length; subLength <<= 1 { // subLength每次左移一位(即subLength=subLength*2)
		prev, cur := dummyHead, dummyHead.Next // cur用来记录链表拆分的位置

		for cur != nil { // 如果链表没有拆分完
			// 3.1 拆分出长度为subLength的链表1
			head1 := cur                                        // 第一个链表的头 即cur初始的位置
			for i := 1; i < subLength && cur.Next != nil; i++ { // 拆分出长度为subLength的链表1
				cur = cur.Next
			}

			// 3.2 拆分出长度为subLength的链表2
			head2 := cur.Next                                                 // 第二个链表的头 即链表1尾部的下一个位置
			cur.Next = nil                                                    // 断开第一个链表和第二个链表的链接
			cur = head2                                                       // 第二个链表头 重新赋值给cur
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ { // 再拆分出长度为subLen的链表2
				cur = cur.Next
			}

			// 3.3 再次断开第二个链表最后的next的链接
			var next *ListNode
			if cur != nil {
				next = cur.Next // next 用于记录 拆分完两个链表的结束位置
				cur.Next = nil  // 断开连接
			}

			// 3.4 合并两个长度为subLength的有序链表
			prev.Next = merge(head1, head2)
			for prev.Next != nil { // 将prev 移动到 subLength * 2的位置后去
				prev = prev.Next
			}
			cur = next // next用于记录拆分完的两个链表的结束位置
		}
	}
	return dummyHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)
