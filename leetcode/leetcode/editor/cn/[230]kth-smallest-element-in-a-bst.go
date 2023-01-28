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
var ans int
var rank int

// 利用 BST（二叉搜索树）的中序遍历结果是有序的（升序）的性质，得到第k个最小元素
func kthSmallest(root *TreeNode, k int) int {
	ans = 0
	rank = 0
	traverse3(root, k)
	return ans
}

func traverse3(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse3(root.Left, k)
	rank++
	if k == rank {
		ans = root.Val
		return
	}
	traverse3(root.Right, k)
}

//leetcode submit region end(Prohibit modification and deletion)
