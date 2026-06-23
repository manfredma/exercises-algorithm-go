// Package common 提供 LeetCode 题目中常用的数据结构。
package common

// TreeNode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode 创建一个叶子节点。
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// BuildTree 按层序（BFS）数组构建二叉树，-1 表示空节点。
// 例：[]int{1, 2, 3, -1, -1, 4, 5} 对应：
//
//	    1
//	   / \
//	  2   3
//	     / \
//	    4   5
func BuildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != -1 {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != -1 {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
