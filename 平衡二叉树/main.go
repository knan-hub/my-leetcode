package main

// 给定一个二叉树，判断它是否是平衡二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	// 方法一
	// if root == nil {
	// 	return true
	// }
	// left := getHeight(root.Left)
	// right := getHeight(root.Right)
	// if abs(left-right) > 1 {
	// 	return false
	// }
	// return isBalanced(root.Left) && isBalanced(root.Right)

	// 方法二
	return checkHeight(root) != -1
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return 1 + max(getHeight(node.Left), getHeight(node.Right))
}

func checkHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := checkHeight(node.Left)
	if left == -1 {
		return -1
	}
	right := checkHeight(node.Right)
	if right == -1 {
		return -1
	}
	if abs(left-right) > 1 {
		return -1
	}
	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
