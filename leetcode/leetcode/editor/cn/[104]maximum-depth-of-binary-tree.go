package cn

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res int
var depth int

func maxDepth(root *TreeNode) int {
	//res, depth = 0, 0
	//traverse(root)
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	return max(leftMax, rightMax) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	depth++
	// 到达叶子节点，判断是否需要更新res
	if root.Left == nil && root.Right == nil && res < depth {
		res = depth
	}
	traverse(root.Left)
	traverse(root.Right)
	depth--
}

//leetcode submit region end(Prohibit modification and deletion)
