package addtwonumbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		l1 *ListNode
		l2 *ListNode
	}{
		{
			l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
			l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
		},
		{
			l1: &ListNode{Val: 0, Next: nil},
			l2: &ListNode{Val: 0, Next: nil},
		},
		{
			l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}},

			l2: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}},
		},
	}

	for _, test := range tests {
		got := AddTwoNumbers(test.l1, test.l2)
		DisplayList(got)
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	var tests = []struct {
		l1 *ListNode
		l2 *ListNode
	}{
		{
			l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
			l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
		},
		{
			l1: &ListNode{Val: 0, Next: nil},
			l2: &ListNode{Val: 0, Next: nil},
		},
		{
			l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}},

			l2: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}},
		},
	}

	for _, test := range tests {
		got := AddTwoNumbers2(test.l1, test.l2)
		DisplayList(got)
	}
}
