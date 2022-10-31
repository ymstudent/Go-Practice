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

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	// 外部变量必须在函数内初始化一下，不然提交到云端后，多个测试用例执行时结果会累计
	maxDiameter = 0
	_ = maxDep(root)
	return maxDiameter
}

func maxDep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDep(root.Left)
	rightMax := maxDep(root.Right)

	// 每一条二叉树的直径就是一个节点的左右子树的最大深度之和
	myDiameter := leftMax + rightMax
	maxDiameter = max2(maxDiameter, myDiameter)

	return max2(leftMax, rightMax) + 1
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
