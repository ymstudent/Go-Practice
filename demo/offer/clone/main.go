package main

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

func main() {

}

func Clone(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &RandomListNode{Label: node.Label, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	res := head.Next
	for head.Next != nil {
		temp := head.Next
		head.Next = head.Next.Next
		head = temp
	}
	return res
}
