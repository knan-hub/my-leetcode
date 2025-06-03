package main

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
// 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 方法一
	// if root == nil || root == p || root == q {
	// 	return root
	// }
	// left := lowestCommonAncestor(root.Left, p, q)
	// right := lowestCommonAncestor(root.Right, p, q)
	// if left != nil && right != nil {
	// 	return root
	// }
	// if left != nil {
	// 	return left
	// }
	// return right

	// 方法二
	// var search func(node *TreeNode) *TreeNode
	// search = func(node *TreeNode) *TreeNode {
	// 	if node == nil || node == p || node == q {
	// 		return node
	// 	}
	// 	left := search(node.Left)
	// 	right := search(node.Right)
	// 	if left != nil && right != nil {
	// 		return node
	// 	}
	// 	if left != nil {
	// 		return left
	// 	}
	// 	return right
	// }
	// return search(root)

	// 方法三
	// 用于记录每个节点的父节点
	parentMap := map[*TreeNode]*TreeNode{}
	stack := []*TreeNode{root}
	parentMap[root] = nil
	// 一直遍历直到找到了 p 和 q 的父节点
	for len(stack) > 0 && !(p != nil && parentMap[p] != nil && q != nil && parentMap[q] != nil) {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			parentMap[node.Left] = node
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			parentMap[node.Right] = node
			stack = append(stack, node.Right)
		}
	}
	// 记录 p 的所有祖先
	ancestors := map[*TreeNode]bool{}
	for p != nil {
		ancestors[p] = true
		p = parentMap[p]
	}
	// 向上遍历 q，找到第一个在 p 祖先集合中的节点
	for q != nil {
		if ancestors[q] {
			return q
		}
		q = parentMap[q]
	}
	return nil // 不可能执行到这里，如果树中一定包含 p 和 q
}
