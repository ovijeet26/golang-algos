package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	var current *ListNode
	var next *ListNode
	var prev *ListNode

	current = head
	next = current
	prev = nil

	for current != nil {
		next = current.Next

		current.Next = prev
		prev = current

		current = next
	}

	return prev
}

func reverseRecursive(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
