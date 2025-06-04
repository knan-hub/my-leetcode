package main

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	// 方法一
	// var check func(left, right *TreeNode) bool
	// check = func(left, right *TreeNode) bool {
	// 	if left == nil && right == nil {
	// 		return true
	// 	}
	// 	if left == nil || right == nil || left.Val != right.Val {
	// 		return false
	// 	}
	// 	return check(left.Left, right.Right) && check(left.Right, right.Left)
	// }
	// return check(root.Left, root.Right)

	// 方法二
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right)
		queue = append(queue, left.Right, right.Left)
	}
	return true
}
