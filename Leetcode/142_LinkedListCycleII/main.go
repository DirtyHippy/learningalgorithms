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

func detectCycle(head *ListNode) *ListNode {
	ptrMap := make(map[*ListNode]struct{})

	for head != nil {
		if _, o := ptrMap[head]; o {
			return head
		}
		ptrMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func main() {
	t1 := NewListNode(1)
	t1.AddNext(2)
	t1.AddNext(3)

	t2 := NewListNode(1)
	t2.AddNext(2)
	t2.AddNext(3)
	t2.Next.Next.Next = t2

	fmt.Println(detectCycle(t1))
}
