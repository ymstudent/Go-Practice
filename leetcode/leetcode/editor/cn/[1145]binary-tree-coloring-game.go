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

// 1号玩家给x节点上色后，整个树可以分为3部分，x的左子树，x的右子树以及
// 包含x父节点的其他节点，哪颗子树最大，2号玩家就选哪颗
// 由于只能选相邻的未着色节点进行染色，所以只要2号玩家选择在x的临近节点上
// 色(左，右，父)，就能占据一整个部分，那么1号玩家接下来就只能往剩下的2个部分上色
// 设2号玩家最多可以染n2个节点，左子树大小为lsz，右子树为rsz，父节点子树的大小就是 n-1-lsz-rsz
// 因此 n2 = max(max(lsz, rsz), n-1-lsz-rsz)
// 1号玩家染色的节点数为n-n2，获胜条件为 n2 > n-n2, 即2 * n2 > n
var lsz, rsz int

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	lsz, rsz = 0, 0
	dfs(root, x)
	return max3(max3(lsz, rsz), n-1-lsz-rsz)*2 > n
}

func dfs(root *TreeNode, x int) int {
	if root == nil {
		return 0
	}

	ls := dfs(root.Left, x)
	rs := dfs(root.Right, x)

	if root.Val == x {
		lsz, rsz = ls, rs
	}

	return ls + rs + 1
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
