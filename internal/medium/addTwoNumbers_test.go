package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	t.Run("example 1", func(t *testing.T) {
		l1 := IntArrToLinkedList([]int{2, 4, 3})
		l2 := IntArrToLinkedList([]int{5, 6, 4})

		sum := addTwoNumbers(l1, l2)

		actual := LinkListToArr(sum)
		expected := []int{7, 0, 8}
		assert.Equal(t, expected, actual)
	})

	t.Run("example 2", func(t *testing.T) {
		l1 := IntArrToLinkedList([]int{0})
		l2 := IntArrToLinkedList([]int{0})

		sum := addTwoNumbers(l1, l2)

		actual := LinkListToArr(sum)
		expected := []int{0}
		assert.Equal(t, expected, actual)
	})

	t.Run("example 3", func(t *testing.T) {
		l1 := IntArrToLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
		l2 := IntArrToLinkedList([]int{9, 9, 9, 9})

		sum := addTwoNumbers(l1, l2)

		actual := LinkListToArr(sum)
		expected := []int{8, 9, 9, 9, 0, 0, 0, 1}
		assert.Equal(t, expected, actual)
	})

	t.Run("example 4 (my own)", func(t *testing.T) {
		l1 := IntArrToLinkedList([]int{1, 2, 3})
		l2 := IntArrToLinkedList([]int{1})

		sum := addTwoNumbers(l1, l2)

		actual := LinkListToArr(sum)
		expected := []int{2, 2, 3}
		assert.Equal(t, expected, actual)
	})

	t.Run("fail 1", func(t *testing.T) {
		l1 := IntArrToLinkedList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
		l2 := IntArrToLinkedList([]int{5, 6, 4})

		sum := addTwoNumbers(l1, l2)

		actual := LinkListToArr(sum)
		expected := []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		assert.Equal(t, expected, actual)
	})

	t.Run("fail 2", func(t *testing.T) {
		l1 := IntArrToLinkedList([]int{9, 9, 1})
		l2 := IntArrToLinkedList([]int{1})

		sum := addTwoNumbers(l1, l2)

		actual := LinkListToArr(sum)
		expected := []int{0, 0, 2}
		assert.Equal(t, expected, actual)
	})

}

func IntArrToLinkedList(arr []int) *ListNode {

	var l = NewListNode(arr[0])
	start := l
	for i, num := range arr {
		if i == 0 {
			continue
		}
		l = Add(l, num)
	}
	return start
}

func LinkListToArr(node *ListNode) []int {
	var nums []int
	next := node
	for next != nil {
		nums = append(nums, next.Val)
		next = next.Next
	}
	return nums
}
