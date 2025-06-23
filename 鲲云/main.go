package main

// 实现一个 无锁队列 (Lock-Free Queue ) ，使用原子操作而非互斥锁

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type Node struct {
	value interface{}
	next  unsafe.Pointer // *Node
}

type LockFreeQueue struct {
	head unsafe.Pointer // *Node
	tail unsafe.Pointer // *Node
}

func NewLockFreeQueue() *LockFreeQueue {
	dummy := &Node{}
	return &LockFreeQueue{
		head: unsafe.Pointer(dummy),
		tail: unsafe.Pointer(dummy),
	}
}

func (q *LockFreeQueue) Enqueue(value interface{}) {
	newNode := unsafe.Pointer(&Node{value: value})

	for {
		tail := loadAtomic(&q.tail)
		next := loadAtomic(&tail.next)

		if tail == loadAtomic(&q.tail) {
			if next == nil {
				if casAtomic(&tail.next, nil, newNode) {
					casAtomic(&q.tail, unsafe.Pointer(tail), newNode)
					return
				}
			} else {
				casAtomic(&q.tail, unsafe.Pointer(tail), unsafe.Pointer(next))
			}
		}
	}
}

func (q *LockFreeQueue) Dequeue() (interface{}, bool) {
	for {
		head := loadAtomic(&q.head)
		tail := loadAtomic(&q.tail)
		next := loadAtomic(&head.next)

		if head == loadAtomic(&q.head) {
			if head == tail {
				if next == nil {
					return nil, false
				}
				casAtomic(&q.tail, unsafe.Pointer(tail), unsafe.Pointer(next))
			} else {
				value := next.value
				if casAtomic(&q.head, unsafe.Pointer(head), unsafe.Pointer(next)) {
					// 避免内存泄漏
					head.value = nil
					return value, true
				}
			}
		}
	}
}

func loadAtomic(p *unsafe.Pointer) *Node {
	return (*Node)(atomic.LoadPointer(p))
}

func casAtomic(p *unsafe.Pointer, old, new unsafe.Pointer) bool {
	return atomic.CompareAndSwapPointer(p, old, new)
}

func main() {
	q := NewLockFreeQueue()

	q.Enqueue("Go")
	q.Enqueue("Rust")
	q.Enqueue("C++")

	v, ok := q.Dequeue()
	fmt.Println(v, ok) // Go true

	v, ok = q.Dequeue()
	fmt.Println(v, ok) // Rust true

	v, ok = q.Dequeue()
	fmt.Println(v, ok) // C++ true

	v, ok = q.Dequeue()
	fmt.Println(v, ok) // <nil> false
}
