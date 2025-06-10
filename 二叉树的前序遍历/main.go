package main

// 给你二叉树的根节点 root ，返回它节点值的前序遍历。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// 方法一
	// result := []int{}
	// var search func(node *TreeNode)
	// search = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	result = append(result, node.Val)
	// 	search(node.Left)
	// 	search(node.Right)
	// }
	// search(root)
	// return result

	// 方法二
	result := []int{}
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			result = append(result, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return result
}
