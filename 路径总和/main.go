package main

// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
// 叶子节点是指没有子节点的节点。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	// 方法一
	// if root == nil {
	// 	return false
	// }
	// var search func(node *TreeNode, sum int) bool
	// search = func(node *TreeNode, sum int) bool {
	// 	if node == nil {
	// 		return false
	// 	}
	// 	sum -= node.Val
	// 	if node.Left == nil && node.Right == nil {
	// 		return sum == 0
	// 	}
	// 	return search(node.Left, sum) || search(node.Right, sum)
	// }
	// return search(root, targetSum)

	// 方法二
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
