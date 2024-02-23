package linkedlist

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

	nodeMap := make(map[*Node]*Node, 0)
	ptr := head

	for ptr != nil {

		newNode := &Node{
			Val:    ptr.Val,
			Next:   nil,
			Random: nil,
		}
		nodeMap[ptr] = newNode
		ptr = ptr.Next
	}

	ptr = head

	for ptr != nil {
		nodeMap[ptr].Next = nodeMap[ptr.Next]
		nodeMap[ptr].Random = nodeMap[ptr.Random]
		ptr = ptr.Next
	}

	return nodeMap[head]
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
