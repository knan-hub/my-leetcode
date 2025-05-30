package main

// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 方法一
	// if list1 == nil {
	// 	return list2
	// }

	// if list2 == nil {
	// 	return list1
	// }

	// if list1.Val <= list2.Val {
	// 	list1.Next = mergeTwoLists(list1.Next, list2)
	// 	return list1
	// }

	// list2.Next = mergeTwoLists(list1, list2.Next)
	// return list2

	// 方法二
	head := &ListNode{}
	cur := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}

	return head.Next
}
