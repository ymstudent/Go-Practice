package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 前序：中左右
// 中序：左中右
// 后续：左右中

var postorderIndex map[int]int

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	postorderIndex = make(map[int]int, len(postorder))
	for k, v := range postorder {
		postorderIndex[v] = k
	}
	return build3(preorder, postorder, 0, len(preorder)-1, 0, len(postorder)-1)
}

func build3(preorder, postorder []int, preStart, preEnd, postStart, postEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	// 根结点：取前序遍历的第一个值
	rootVal := preorder[preStart]
	// 左节点：前序遍历的第二个值
	leftRootVal := preorder[preStart+1]
	// 根据左节点在后序遍历中的位置，确定左子树长度
	index := postorderIndex[leftRootVal]
	// 获取左子树的长度
	leftSize := index - postStart + 1

	// 构造根节点
	root := &TreeNode{Val: rootVal}
	// 构造左右子树
	root.Left = build(preorder, postorder, preStart+1, preStart+leftSize, postStart, index)
	root.Right = build(preorder, postorder, preStart+leftSize+1, preEnd, index+1, postEnd-1)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
