package cn

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var memo map[string]int
var res2  []*TreeNode
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	memo = make(map[string]int)
	res2 = make([]*TreeNode, 0)
	traverse2(root)
	return res2
}

func traverse2(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	// 后序遍历：先遍历左右子树，再进行操作
	left := traverse2(root.Left)
	right := traverse2(root.Right)
	subTree := left + "," + right + "," + strconv.Itoa(root.Val)

	num, _ := memo[subTree]
	if num == 1 {
		res2 = append(res2, root)
	}
	memo[subTree]++
	return subTree
}

//leetcode submit region end(Prohibit modification and deletion)
