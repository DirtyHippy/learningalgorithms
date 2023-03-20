package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode returns a new list node.
func NewListNode(v int) *ListNode {
	return &ListNode{v, nil}
}

// AddNext adds a next node to the end of list.
func (l *ListNode) AddNext(v int) {
	for n := l; n != nil; n = n.Next {
		if n.Next == nil {
			new := NewListNode(v)
			n.Next = new
			break
		}
	}
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	t11 := NewListNode(1)
	t11.AddNext(2)
	t11.AddNext(3)
	t11.AddNext(4)
	t11.AddNext(5)

	fmt.Println(middleNode(t11))
}
