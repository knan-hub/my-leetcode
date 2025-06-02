package main

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	newHeadA := headA
	newHeadB := headB
	for newHeadA != newHeadB {
		if newHeadA != nil {
			newHeadA = newHeadA.Next
		} else {
			newHeadA = headB
		}
		if newHeadB != nil {
			newHeadB = newHeadB.Next
		} else {
			newHeadB = headA
		}
	}
	return newHeadA
}
