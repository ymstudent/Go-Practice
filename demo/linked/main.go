package main

type Node struct {
	Val int
	Next *Node
	
}

type List struct {
	headNode *Node
}



func main()  {

}

func (this *List) isEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

func (this *List) length() int {
	cur := this.headNode
	count := 0

	for cur != nil  {
		count++
		cur = cur.Next
	}

	return count
}

func (this *List) get(index int) int {
	cur := this.headNode
	count := 0

	for cur != nil {
		if count == index {
			return cur.Val
		} else {
			count++
			cur = cur.Next
		}
	}
	return -1
}

func (this *List) addAtHead(val int) {
	node := &Node{Val:val}
	node.Next = this.headNode
	this.headNode = node
}

func (this *List) addAtTail(val int) {
	node := &Node{Val:val}
	if this.isEmpty() {
		this.headNode = node
	} else {
		cur := this.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (this *List) addAtIndex(index, val int) {
	listLength := this.length()

	if index < 0 {
		this.addAtHead(val)
	} else if index >= listLength {
		this.addAtTail(val)
	} else {
		node := &Node{Val:val}
		pre := this.headNode
		count := 0
		for count < index - 1 {
			count++
			pre = pre.Next
		}
		node.Next = pre.Next
		pre.Next = node
	}
}

func (this *List) deleteAtIndex(index int) {
	listLength := this.length()
	if index >= 0 && index < listLength {
		cur := this.headNode
		count := 0

		for cur != nil {
			if count == index - 1 {
				cur.Next = cur.Next.Next
			} else {
				count++
				cur = cur.Next
			}
		}
	}
}