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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	current := res

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}
	return res.Next
}

func main() {
	t11 := NewListNode(1)
	t11.AddNext(2)
	t11.AddNext(4)

	t12 := NewListNode(1)
	t12.AddNext(3)
	t12.AddNext(4)

	fmt.Println(mergeTwoLists(t11, t12))
}
