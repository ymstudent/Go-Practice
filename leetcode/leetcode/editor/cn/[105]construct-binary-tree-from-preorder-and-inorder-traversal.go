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

var inorderMap map[int]int

func buildTree2(preorder []int, inorder []int) *TreeNode {
	inorderMap = make(map[int]int)
	for k, v := range inorder {
		inorderMap[v] = k
	}

	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preorder, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	// 根结点：取前序遍历的第一个值
	rootVal := preorder[preStart]
	// 根据根节点在中序遍历中的位置，区分根节点的左右节点
	index := inorderMap[rootVal]
	// 获取左子树的长度
	leftSize := index - inStart

	// 构造根节点
	root := &TreeNode{Val: rootVal}
	// 构造左右子树
	root.Left = build(preorder, inorder, preStart+1, preStart+leftSize, inStart, index-1)
	root.Right = build(preorder, inorder, preStart+leftSize+1, preEnd, index+1, inEnd)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
