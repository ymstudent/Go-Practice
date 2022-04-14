package main

func main() {

}

// 归并排序，稳定的非原地排序算法，时间复杂度：O(nlogn)，空间复杂度：O(n)
func mergeSort(a []int) []int {
	l := len(a)

	if l <= 1 {
		return a
	}

	mid := l / 2

	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	leftL := len(left)
	rightL := len(right)

	i, j := 0, 0

	temp := make([]int, leftL+rightL)
	for i < leftL && j < rightL {
		if left[i] < right[j] {
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			j++
		}
	}

	if i < leftL {
		temp = append(temp, left[i:]...)
	} else {
		temp = append(temp, right[j:]...)
	}

	return temp
}

// 归并排序练习题
// 148 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}

	}

	mid := slow
	return merge2(sort(head, mid), sort(mid, tail))
}

// 合并2个有序链表
func merge2(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	if temp1 != nil && temp2 != nil {
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
