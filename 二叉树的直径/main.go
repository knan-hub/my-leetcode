package main

// 给你一棵二叉树的根节点，返回该树的直径。
// 二叉树的直径是指树中任意两个节点之间最长路径的长度。这条路径可能经过也可能不经过根节点 root 。
// 两节点之间路径的长度由它们之间边数表示。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var maxHeight func(node *TreeNode) int
	maxHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := maxHeight(node.Left)
		right := maxHeight(node.Right)
		maxDiameter = max(maxDiameter, left+right)
		return 1 + max(left, right)
	}
	maxHeight(root)
	return maxDiameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
