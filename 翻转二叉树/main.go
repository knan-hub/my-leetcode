package main

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := root.Left
	right := root.Right
	root.Left = invertTree(right)
	root.Right = invertTree(left)
	return root
}
