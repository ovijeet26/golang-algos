package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	slowPtr := head
	fastPtr := head

	isCycleDetected := false

	for fastPtr != nil && fastPtr.Next != nil && fastPtr.Next.Next != nil {

		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

		if slowPtr == fastPtr {
			isCycleDetected = true
			break
		}
	}

	if isCycleDetected {

		beginingPtr := head
		secondPointer := fastPtr

		for beginingPtr != secondPointer {

			beginingPtr = beginingPtr.Next
			secondPointer = secondPointer.Next
		}
		return beginingPtr

	}
	return nil
}

func detectCycleUsingSet(head *ListNode) *ListNode {

	// Make-shift set data type in Golang.
	// Empty struct takes 0 bits in Golang.
	seenNodes := make(map[*ListNode]struct{}, 0)
	present := struct{}{}

	currentNode := head

	for currentNode != nil {

		_, isPresent := seenNodes[currentNode]

		if isPresent {
			break
		}

		seenNodes[currentNode] = present
		currentNode = currentNode.Next
	}

	return currentNode
}
