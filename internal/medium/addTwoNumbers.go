package medium

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(num int) *ListNode {
	return &ListNode{
		Val:  num,
		Next: nil,
	}
}
func Add(l *ListNode, num int) *ListNode {
	new := &ListNode{
		Val:  num,
		Next: nil,
	}

	l.Next = new
	return new
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumL *ListNode = nil
	var start *ListNode = nil

	carry := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = 0

		digit := sum

		if sum > 9 {
			digit = sum % 10
			carry = sum / 10
		}

		if sumL != nil {
			sumL = Add(sumL, digit)
		} else {
			sumL = NewListNode(digit)
			start = sumL
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	// Remaining Digits
	var remaingDigits *ListNode = nil

	if l1 != nil {
		remaingDigits = l1
	} else {
		remaingDigits = l2
	}

	for remaingDigits != nil {
		sum := remaingDigits.Val + carry
		carry = 0
		digit := sum

		if sum > 9 {
			digit = sum % 10
			carry = sum / 10
		}
		sumL = Add(sumL, digit)

		remaingDigits = remaingDigits.Next
	}

	if carry != 0 {
		sumL = Add(sumL, carry)
	}

	return start
}
