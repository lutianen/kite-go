// 两数相加 https://leetcode.cn/problems/add-two-numbers/description/

package addtwonumbers

import "fmt"

// Signly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var count int = 0
	var prev *ListNode = &ListNode{}
	curr := prev

	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + count
		count = val / 10
		curr.Next = &ListNode{
			Val:  val % 10,
			Next: nil,
		}

		curr = curr.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		val := l1.Val + count
		count = val / 10
		curr.Next = &ListNode{
			Val:  val % 10,
			Next: nil,
		}

		curr = curr.Next
		l1 = l1.Next
	}

	for l2 != nil {
		val := l2.Val + count
		count = val / 10
		curr.Next = &ListNode{
			Val:  val % 10,
			Next: nil,
		}

		curr = curr.Next
		l2 = l2.Next
	}

	if count != 0 {
		curr.Next = &ListNode{
			Val:  count,
			Next: nil,
		}
	}

	return prev.Next
}

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	prev := &ListNode{}
	curr := prev
	count := 0
	for l1 != nil || l2 != nil {
		var (
			l1Val int
			l2Val int
		)
		if l1 == nil {
			l1Val = 0
		} else {
			l1Val = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			l2Val = 0
		} else {
			l2Val = l2.Val
			l2 = l2.Next
		}
		val := l1Val + l2Val + count
		count = val / 10
		curr.Next = &ListNode{
			Val:  val % 10,
			Next: nil,
		}

		curr = curr.Next
	}

	if count != 0 {
		curr.Next = &ListNode{
			Val:  count,
			Next: nil,
		}
	}

	return prev.Next
}

func DisplayList(head *ListNode) {
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d -> ", head.Val)
		} else {
			fmt.Printf("%d", head.Val)
		}

		head = head.Next
	}
	fmt.Println()
}
