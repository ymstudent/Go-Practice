package main
//反转链表
type Linked struct{
	val int
	next *Linked
}


func main()  {

}

func resver(link *Linked) *Linked {
	head := link
	var newHead *Linked
	for link != nil {
		node := head
		head = head.next

		node.next = newHead
		newHead = node
	}
	return newHead
}
