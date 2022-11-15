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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 先用 2，3，4 来思考
	// 左右子树子节点都为nil，所以递归立即返回了
	flatten(root.Left)
	flatten(root.Right)

	// 暂存 左、右子树
	left := root.Left
	right := root.Right

	// 将左子树指向nil
	root.Left = nil
	// 然后将左子树的值付给右子树，这一步完成后根节点为2，左子树为nil，右子树为3
	root.Right = left

	// 从根节点2开始遍历， 遍历完后，p = 3
	p := root
	for p.Right != nil {
		p = p.Right
	}
	// 将原来右子树放到3的右子树节点上，这样树就变成了[2, nil, 3, nil, 4, nil]
	p.Right = right
}

//leetcode submit region end(Prohibit modification and deletion)
