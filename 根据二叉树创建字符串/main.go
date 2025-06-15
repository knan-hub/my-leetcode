package main

import "strconv"

// 给你二叉树的根节点 root ，请你采用前序遍历的方式，将二叉树转化为一个由括号和整数组成的字符串，返回构造出的字符串。
// 空节点使用一对空括号对 "()" 表示，转化后需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	// 方法一
	// if root == nil {
	// 	return ""
	// }
	// res := strconv.Itoa(root.Val)
	// if root.Left != nil || root.Right != nil {
	// 	res += "(" + tree2str(root.Left) + ")"
	// }
	// if root.Right != nil {
	// 	res += "(" + tree2str(root.Right) + ")"
	// }
	// return res

	// 方法二
	if root == nil {
		return ""
	}
	s := ""
	var search func(node *TreeNode)
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		s += strconv.Itoa(node.Val)
		if node.Left != nil || node.Right != nil {
			s += "("
			search(node.Left)
			s += ")"
		}
		if node.Right != nil {
			s += "("
			search(node.Right)
			s += ")"
		}
	}
	search(root)
	return s
}
