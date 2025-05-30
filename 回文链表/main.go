package main

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	slow := dummy
	fast := dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondHalf := slow.Next
	slow.Next = nil
	var reversed *ListNode
	for secondHalf != nil {
		next := secondHalf.Next
		secondHalf.Next = reversed
		reversed = secondHalf
		secondHalf = next
	}
	for head != nil && reversed != nil {
		if head.Val != reversed.Val {
			return false
		}
		head = head.Next
		reversed = reversed.Next
	}
	return true
}
