package main

// 给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	// 方法一
	// memo := map[int]struct{}{}
	// isFind := false
	// var search func(node *TreeNode)
	// search = func(node *TreeNode) {
	// 	if node == nil || isFind {
	// 		return
	// 	}
	// 	num := k - node.Val
	// 	if _, ok := memo[num]; ok {
	// 		isFind = true
	// 		return
	// 	}
	// 	memo[node.Val] = struct{}{}
	// 	search(node.Left)
	// 	search(node.Right)
	// }
	// search(root)
	// return isFind

	// 方法二
	// if root == nil {
	// 	return false
	// }
	// memo := map[int]struct{}{}
	// queue := []*TreeNode{root}
	// for len(queue) > 0 {
	// 	node := queue[0]
	// 	queue = queue[1:]
	// 	if _, ok := memo[k-node.Val]; ok {
	// 		return true
	// 	}
	// 	memo[node.Val] = struct{}{}
	// 	if node.Left != nil {
	// 		queue = append(queue, node.Left)
	// 	}
	// 	if node.Right != nil {
	// 		queue = append(queue, node.Right)
	// 	}
	// }
	// return false

	// 方法三
	if root == nil {
		return false
	}
	leftStack := []*TreeNode{}
	rightStack := []*TreeNode{}
	for node := root; node != nil; node = node.Left {
		leftStack = append(leftStack, node)
	}
	for node := root; node != nil; node = node.Right {
		rightStack = append(rightStack, node)
	}
	next := func(stack *[]*TreeNode, isLeft bool) *TreeNode {
		n := len(*stack)
		if n == 0 {
			return nil
		}
		node := (*stack)[n-1]
		*stack = (*stack)[:n-1]
		currentNode := node
		if isLeft {
			currentNode = currentNode.Right
			for currentNode != nil {
				*stack = append(*stack, currentNode)
				currentNode = currentNode.Left
			}
		} else {
			currentNode = currentNode.Left
			for currentNode != nil {
				*stack = append(*stack, currentNode)
				currentNode = currentNode.Right
			}
		}
		return node
	}
	leftNode := next(&leftStack, true)
	rightNode := next(&rightStack, false)
	for leftNode != nil && rightNode != nil && leftNode != rightNode {
		sum := leftNode.Val + rightNode.Val
		if sum == k {
			return true
		} else if sum > k {
			rightNode = next(&rightStack, false)
		} else {
			leftNode = next(&leftStack, true)
		}
	}
	return false
}
