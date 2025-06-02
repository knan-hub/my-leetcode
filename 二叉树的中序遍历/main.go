package main

// 给定一个二叉树的根节点 root ，返回它的中序遍历 。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// 方法一
	// result := make([]int, 0)
	// var search func(node *TreeNode)
	// search = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	search(node.Left)
	// 	result = append(result, node.Val)
	// 	search(node.Right)
	// }
	// search(root)
	// return result

	// 方法二
	// result := make([]int, 0)
	// stack := make([]*TreeNode, 0)
	// cur := root
	// for cur != nil || len(stack) > 0 {
	// 	for cur != nil {
	// 		stack = append(stack, cur)
	// 		cur = cur.Left
	// 	}
	// 	cur = stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]
	// 	result = append(result, cur.Val)
	// 	cur = cur.Right
	// }
	// return result

	// 方法三
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	res := inorderTraversal(root.Left)
	for _, v := range res {
		result = append(result, v)
	}
	result = append(result, root.Val)
	res = inorderTraversal(root.Right)
	for _, v := range res {
		result = append(result, v)
	}
	return result
}
