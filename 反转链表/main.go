package main

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 方法一
	// newHead := &ListNode{
	// 	Val:  -1,
	// 	Next: nil,
	// }

	// for head != nil {
	// 	next := head.Next
	// 	head.Next = newHead.Next
	// 	newHead.Next = head
	// 	head = next
	// }

	// return newHead.Next

	// 方法二
	// var pre *ListNode
	// cur := head

	// for cur != nil {
	// 	next := cur.Next
	// 	cur.Next = pre
	// 	pre = cur
	// 	cur = next
	// }

	// return pre

	// 方法三
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
