package cn

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const SEP = ","
const NULL = "#"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := serialize(root, "")
	//fmt.Println(res)
	return res
}

// 1,2,#,#,3,4,#,#,5,#,#,
func serialize(root *TreeNode, s string) string {
	if root == nil {
		s += NULL + SEP
		return s
	}
	// 前序遍历 中->左->右
	s += strconv.Itoa(root.Val) + SEP
	s = serialize(root.Left, s)
	s = serialize(root.Right, s)
	return s
}

// Deserializes your encoded data to tree.
var nodes []string
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, SEP)
	nodes = make([]string, 0, len(s))
	for _, v := range s {
		nodes = append(nodes, v)
	}
	//fmt.Println(nodes)
	res := deserialize()
	//fmt.Println(res, res.Left, res.Right)
	return res
}

// 解开后序遍历
func deserialize() *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	first := nodes[0]
	nodes = nodes[1:]
	if first == NULL {
		return nil
	}
	rootVal, _ := strconv.Atoi(first)
	root := &TreeNode{Val: rootVal}
	root.Left = deserialize()
	root.Right = deserialize()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
//leetcode submit region end(Prohibit modification and deletion)
