package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow

	mid = reverse(mid)

	ptr1 := head
	ptr2 := mid

	for ptr1 != nil && ptr2 != nil && ptr1 != ptr2 {

		temp1 := ptr1.Next
		temp2 := ptr2.Next

		ptr1.Next = ptr2
		ptr2.Next = temp1

		ptr1 = temp1
		ptr2 = temp2
	}

}

func reverse(head *ListNode) *ListNode {

	var c *ListNode
	var n *ListNode
	var p *ListNode
	c = head
	n = head.Next
	p = nil

	for n != nil {

		n = c.Next

		c.Next = p
		p = c
		c = n
	}

	return p
}
