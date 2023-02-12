package linkedlist

//Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// leet code link -> https://leetcode.com/problems/copy-list-with-random-pointer/
func copyRandomList(head *Node) *Node {

	// Maintain a mapping of the clone of every node
	nodeMap := make(map[*Node]*Node)

	// Create the new head
	newHead := new(Node)
	newHead.Val = head.Val

	ptr1 := head
	ptr2 := newHead

	// Map the cloned head with the original head
	nodeMap[ptr1] = ptr2

	// Traverse teh old list and generate the linear cloned list
	for ptr1.Next != nil {

		ptr1 = ptr1.Next

		newNode := new(Node)
		newNode.Val = ptr1.Val

		// Maintain a 1-1 map of each new node with original node
		nodeMap[ptr1] = newNode

		ptr2.Next = newNode
		ptr2 = ptr2.Next
	}

	newPtr := newHead
	oldPtr := head

	// traverse the lista gain and assign the random pointer
	// based on the nodeMap created in previous pass
	for oldPtr != nil {

		newPtr.Random = nodeMap[oldPtr.Random]

		newPtr = newPtr.Next
		oldPtr = oldPtr.Next
	}

	return newHead
}
