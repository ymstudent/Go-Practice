package cn

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// 思路：将二叉树抽象成一个三叉树，遍历每个三叉树节点，将气内部的2个节点串联起来就行
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	traver(root.Left, root.Right)
	return root
}

func traver(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	// 前序位置，将2个节点串联起来
	node1.Next = node2
	// 连接相同父节点的2个子节点
	traver(node1.Left, node1.Right)
	traver(node2.Left, node2.Right)
	// 连接跨域父节点的2个子节点
	traver(node1.Right, node2.Left)
}

//leetcode submit region end(Prohibit modification and deletion)
