package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime: 15ms
// Memory Usage: 4.62MB
// Time complexity: O(max(m,n)) -> m and n are the length of l1 and l2
// Space complexity: O(max(m,n)) -> the length of the new list is at most max(m,n)+1
// Category: Linked List
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode // head and tail of the new list
	var carry int            // carry is 0 or 1

	// loop until l1 and l2 are nil
	for l1 != nil || l2 != nil {
		var sum int

		// if l1 or l2 is not nil, add the value to sum
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// if carry is not 0, add it to sum
		sum += carry

		// create a new node with the value of sum % 10
		node := &ListNode{
			Val: sum % 10,
		}

		// if head is nil, set head to node
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}

		// set tail to node
		tail = node

		// set carry to sum / 10
		carry = sum / 10
	}

	// if carry is not 0, create a new node with the value of carry
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}

	return head
}