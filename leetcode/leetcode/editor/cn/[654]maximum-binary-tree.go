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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx, max :=0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			idx = i
			max = nums[i]
		}
	}
	p := &TreeNode{Val: max}
	p.Left = constructMaximumBinaryTree(nums[0:idx])
	// 注意边界，如果idx是切片最后一位，那么idx + 1会出现数组越界
	var rightNms []int
	if idx + 1 < len(nums) {
		rightNms = nums[idx+1:]
	}
	p.Right = constructMaximumBinaryTree(rightNms)
	return p
}
//leetcode submit region end(Prohibit modification and deletion)
