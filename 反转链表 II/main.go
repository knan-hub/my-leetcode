package main

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回反转后的链表。

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	dummyHead := dummy
	count := left - 1
	for dummy.Next != nil && count > 0 {
		dummy = dummy.Next
		count--
	}
	if count > 0 || dummy.Next == nil {
		return head
	}
	newHead := dummy.Next
	secondHead := newHead
	var tail *ListNode
	count = right - left + 1
	for secondHead != nil && count > 0 {
		next := secondHead.Next
		secondHead.Next = tail
		tail = secondHead
		secondHead = next
		count--
	}
	dummy.Next = tail
	newHead.Next = secondHead
	return dummyHead.Next
}
