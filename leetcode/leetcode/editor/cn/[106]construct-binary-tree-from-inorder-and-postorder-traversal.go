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
var inorderMap2 map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderMap2 = make(map[int]int)

	for k, v := range inorder {
		inorderMap2[v] = k
	}

	return build2(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func build2(inorder, postorder []int, inStart, inEnd, posStart, posEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	rootVal := postorder[posEnd]
	index := inorderMap2[rootVal]
	leftSize := index - inStart

	// 中序： 左：inStart:index-1->中：index->右：index+1:inEnd
	// 后序： 左：->右->后

	root := &TreeNode{Val: rootVal}
	root.Left = build2(inorder, postorder, inStart, index-1, posStart, posStart+leftSize-1)
	root.Right = build2(inorder, postorder, index+1, inEnd, posStart+leftSize, posEnd-1)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
