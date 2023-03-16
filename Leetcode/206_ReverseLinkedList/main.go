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

func reverseList(head *ListNode) *ListNode {
	current := head
	var previous, next *ListNode

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}

func main() {
	t11 := NewListNode(1)
	t11.AddNext(2)
	t11.AddNext(3)
	t11.AddNext(4)
	t11.AddNext(5)

	fmt.Println(reverseList(t11))
}
