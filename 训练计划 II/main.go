package main

// 给定一个头节点为 head 的链表用于记录一系列核心肌群训练项目编号，请查找并返回倒数第 cnt 个训练项目编号。

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlan(head *ListNode, cnt int) *ListNode {
	if head == nil {
		return nil
	}
	slow := &ListNode{
		Val:  -1,
		Next: head,
	}
	fast := &ListNode{
		Val:  -1,
		Next: head,
	}
	for cnt > 0 && fast != nil {
		fast = fast.Next
		cnt--
	}
	if cnt > 0 && fast == nil {
		return nil
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
