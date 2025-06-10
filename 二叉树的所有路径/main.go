package main

import "strconv"

// 给你一个二叉树的根节点 root ，按任意顺序 ，返回所有从根节点到叶子节点的路径。
// 叶子节点是指没有子节点的节点。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	// 方法一
	// if root == nil {
	// 	return nil
	// }
	// result := []string{}
	// var search func(node *TreeNode, path string)
	// search = func(node *TreeNode, path string) {
	// 	if node.Left == nil && node.Right == nil {
	// 		result = append(result, path)
	// 		return
	// 	}
	// 	if node.Left != nil {
	// 		search(node.Left, path+"->"+strconv.Itoa(node.Left.Val))
	// 	}
	// 	if node.Right != nil {
	// 		search(node.Right, path+"->"+strconv.Itoa(node.Right.Val))
	// 	}
	// }
	// search(root, strconv.Itoa(root.Val))
	// return result

	// 方法二
	if root == nil {
		return nil
	}
	result := []string{}
	nodeQueue := []*TreeNode{root}
	pathQueue := []string{strconv.Itoa(root.Val)}
	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			result = append(result, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return result
}
