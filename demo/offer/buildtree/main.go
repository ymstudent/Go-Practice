package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 概念：
// 前序遍历：先遍历根节点，再递归的遍历它的左子树，最后递归的遍历右子树，即：根，左，右
// 中序遍历：先递归的遍历左子树，再遍历根节点，最后递归的遍历右子树，即：左，根，右
// 后序遍历：先递归的遍历左子树，再递归的遍历右子树，最后遍历根节点，即：左，右，根
// 前、中、后序中的"序"，都是针对当前父节点与其左右子节点的前后处理关系
// 通过上面的分析可以看出，遍历时都是先打印左树，再打印右树，只是前序是根节点最新打印，
// 中序是根节点在中间打印，后序是根节点最后打印

// 输入某二插树的前序遍历与中序遍历结果，请构建二叉树并返回其根节点
func buildTree(preorde, inorder []int) *TreeNode {
	if len(preorde) == 0 {
		return nil
	}
	// 前序遍历的第一个节点为根节点
	root := &TreeNode{preorde[0], nil, nil}
	i := 0
	// 找到中序遍历中的根节点所在的位置，根据根节点位置，我们可以知道左子树有多少个节点，右子树右多少个节点
	for ; i < len(inorder); i++ {
		if inorder[i] == preorde[0] {
			break
		}
	}
	// preorde[1:len(inorder[:i])+1]， 前序列遍历结果中，左子树的所有节点
	// preorde[len(inorder[:i])+1:]， 前序遍历结果中，右子树的所有节点
	root.Left = buildTree(preorde[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorde[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
