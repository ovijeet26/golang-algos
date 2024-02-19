package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	slowPtr := head
	fastPtr := head

	for fastPtr != nil && fastPtr.Next != nil && fastPtr.Next.Next != nil {

		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

		if slowPtr == fastPtr {
			return true
		}
	}

	return false
}

func hasCycle_alternate(head *ListNode) bool {

	// Make-shift set data type in Golang.
	// Empty struct takes 0 bits in Golang.
	seenNodes := make(map[*ListNode]struct{}, 0)
	present := struct{}{}

	currentNode := head

	for currentNode != nil {

		_, isPresent := seenNodes[currentNode]

		if isPresent {
			return true
		}

		seenNodes[currentNode] = present
		currentNode = currentNode.Next
	}

	return false
}
