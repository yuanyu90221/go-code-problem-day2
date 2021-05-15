package sol

import (
	"fmt"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

func setUpNextNodeWithVal(l *ListNode, val int) {
	valNode := ListNode{Val: val, Next: nil}
	l.Next = &valNode
}

func ConvertList(a *[]int) *ListNode {
	head := &ListNode{}
	prev := &ListNode{}
	for i, v := range *a {
		// println(i, v)
		cur := ListNode{Val: v, Next: nil}
		if i == 0 {
			head = &cur
			prev = &cur
		} else {
			prev.Next = &cur
			prev = prev.Next
		}
	}
	return head
}

func Transerval(h *ListNode) {
	for ; h != nil; h = h.Next {
		fmt.Printf("%d->", h.Val)
	}
	fmt.Println("nil")
}

func findTargetRegionNode(l *ListNode, start int, end int) (*ListNode, *ListNode, []*ListNode) {
	prevNode := &ListNode{}
	nextNode := &ListNode{}
	resultArray := []*ListNode{}
	visit := 1
	for ; visit <= end; visit++ {
		if start == 1 {
			prevNode = l
		} else if visit == start-1 {
			prevNode = l
		}
		if visit >= start && visit <= end {
			resultArray = append([]*ListNode{l}, resultArray...)
		}
		l = l.Next
	}
	nextNode = l
	return prevNode, nextNode, resultArray
}

func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	resultHead := &ListNode{}
	prevNode, nextNode, stack := findTargetRegionNode(head, m, n)
	prevNode.Next = stack[0]
	stack[len(stack)-1].Next = nextNode
	prev := &ListNode{}
	for i, v := range stack {
		if i == 0 {
			prev = v
		} else {
			prev.Next = v
			prev = v
		}
	}
	if m == 1 {
		return stack[0]
	}
	resultHead = head
	return resultHead
}
