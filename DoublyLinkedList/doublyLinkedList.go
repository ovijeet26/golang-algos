package doublylinkedlist

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// Leet code https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/
func Flatten(root *Node) *Node {

	if root == nil || (root.Child == nil && root.Next == nil) {
		return nil
	}

	currentNode := root

	// Traverse the main list.
	for currentNode != nil {

		if currentNode.Child != nil {

			nextToCurrentNode := currentNode.Next

			childNode := currentNode.Child

			// Traverse and get the tail node of the child list.
			for childNode.Next != nil {

				childNode = childNode.Next
			}

			childNode.Next = nextToCurrentNode

			// Make sure that the next node to the current node actually exists.
			// i.e. currentNode isn't the last node of the list
			if nextToCurrentNode != nil {

				nextToCurrentNode.Prev = childNode
			}

			currentNode.Next = currentNode.Child
			currentNode.Child.Prev = currentNode

			// Make the child node of the parent as nil, since we have already absorbed them.
			currentNode.Child = nil
		}

		currentNode = currentNode.Next
	}

	return root
}
