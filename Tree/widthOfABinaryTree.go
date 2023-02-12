package tree

import "fmt"

// Leet code link -> https://leetcode.com/problems/maximum-width-of-binary-tree/
// https://www.youtube.com/watch?v=ZbybYvcVLks&ab_channel=takeUforward
func widthOfBinaryTree(root *TreeNode) int {

	// Taking care of edge cases
	if root == nil {

		return 0
	}

	if root.Left == nil && root.Right == nil {

		return 1
	}

	toVisit := Queue{}

	// Track the index of every node. Root being 0.
	rootNode := NodeIndex{
		node:  root,
		index: 0,
	}

	toVisit.Enqueue(rootNode)

	// Track every level
	levelCount := toVisit.Size()
	maxWidth := 0

	for toVisit.Size() > 0 {

		// If a level is over and the queue has elements, then the contents of the queue are indicative of a new level.
		if levelCount == 0 && toVisit.Size() > 0 {

			// We are capturing the first and last elements of the queue array.
			leftElem := toVisit.First().(NodeIndex)
			rightElem := toVisit.Last().(NodeIndex)

			// Width of the current level will be the difference in index + 1.
			currentWidth := (rightElem.index - leftElem.index) + 1

			// Update max width, if necessary.
			if currentWidth > maxWidth {
				maxWidth = currentWidth
			}

			// Update level count.
			levelCount = toVisit.Size()
		}

		// Update level count in case we are ending a level ( root case will not be updated in the upper if block)
		if levelCount == 0 {

			levelCount = toVisit.Size()
		}

		current := toVisit.Dequeue().(NodeIndex)
		levelCount--

		// Insert the left child of the currentNode if it exists, with the proper index.
		// if root = n :
		// Left child -> 2n+1
		if current.node.Left != nil {

			leftNode := NodeIndex{
				node:  current.node.Left,
				index: (current.index * 2) + 1,
			}

			toVisit.Enqueue(leftNode)
		}

		// Insert the right child of the currentNode if it exists, with the proper index.
		// if root = n :
		// Right child -> 2n+1
		if current.node.Right != nil {

			rightNode := NodeIndex{
				node:  current.node.Right,
				index: (current.index * 2) + 2,
			}

			toVisit.Enqueue(rightNode)
		}

	}

	return maxWidth
}

type NodeIndex struct {
	node  *TreeNode
	index int
}

// Regular queue implementation

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(data interface{}) {

	q.items = append(q.items, data)
}

func (q *Queue) Dequeue() interface{} {

	toRemove := q.items[0]

	// Remove first element from slice
	q.items = q.items[1:]
	return toRemove
}

func (q *Queue) Size() int {

	return len(q.items)
}

func (q *Queue) First() interface{} {

	return q.items[0]
}

func (q *Queue) Last() interface{} {

	return q.items[len(q.items)-1]
}

func WidthTreeCaller() {

	a := TreeNode{Value: 1}
	b := TreeNode{Value: 2}
	c := TreeNode{Value: 3}
	d := TreeNode{Value: 4}
	e := TreeNode{Value: 5}
	f := TreeNode{Value: 6}

	a.Left = &b
	a.Right = &c

	b.Left = &d
	c.Right = &e

	d.Left = &f

	width := widthOfBinaryTree(&a)

	fmt.Print(width)

	//			1
	//		2		3
	//	4				5
	//6
}
