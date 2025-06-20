package main

// init 输入一个 n，初始化数据结构
// put 输入一个 int，该数据结构只存储最近 n 个 put 的数字
// get 输出存储数字平均值，没有则为 0
// 要求上除方法时间复杂度为 O(1)，不能使用数组

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Count int
	Total int
	Sum   int
	Head  *ListNode
	Tail  *ListNode
}

func (l *List) init(n int) {
	if n <= 0 {
		panic("List count must be positive")
	}
	l.Count = n
	l.Total = 0
	l.Sum = 0
	l.Head = nil
	l.Tail = nil
}

func (l *List) put(n int) {
	node := &ListNode{
		Val: n,
	}

	if l.Total == l.Count && l.Total > 0 {
		oldHead := l.Head
		l.Head = l.Head.Next
		l.Sum -= oldHead.Val
		oldHead.Next = nil
		l.Total--
	}

	if l.Total == 0 {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}

	l.Sum += n
	l.Total++
}

func (l *List) get() (float64, bool) {
	if l.Total == 0 {
		return 0, false
	}
	return float64(l.Sum) / float64(l.Total), true
}
