package main

// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回已排序的链表。

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := head
	for newHead.Next != nil {
		next := newHead.Next
		if newHead.Val == next.Val {
			newHead.Next = next.Next
			continue
		}
		newHead = next
	}
	return head
}
