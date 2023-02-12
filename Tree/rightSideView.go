package tree

import queue "github.com/ovijeet/go/algos/Queue"

// Leet code link -> https://leetcode.com/problems/binary-tree-right-side-view/
func RightSideViewDFSApproach(root *TreeNode) []int {

	rightSideView := new([]int)

	modifiedPreOrder(root, 1, rightSideView)
	return *rightSideView
}

// If we traverse the right sub tree first in pre, order, we will get near about the same
// pattern as the put put. We just have to track the level of the tree,
// and take the first node we traverse at that level.
func modifiedPreOrder(node *TreeNode, level int, rightSideView *[]int) {

	if node == nil {
		return
	}

	if len(*rightSideView) < level {

		*rightSideView = append(*rightSideView, node.Value)
	}

	modifiedPreOrder(node.Right, level+1, rightSideView)
	modifiedPreOrder(node.Left, level+1, rightSideView)
}

func rightSideViewBFSApproach(root *TreeNode) []int {

	if root == nil {

		return []int{}
	}

	visitedQueue := queue.Queue{}
	levelArray := make([]int, 0)
	finalOutput := make([]int, 0)

	visitedQueue.Enqueue(root)
	levelCount := 1

	// Capture each level of the tree in an array and always take the right-most element.
	// Follow level order traversal technique.
	for visitedQueue.Size() > 0 {

		currentNode := visitedQueue.Dequeue().(*TreeNode)
		levelArray = append(levelArray, currentNode.Value)
		levelCount--

		if currentNode.Left != nil {
			visitedQueue.Enqueue(currentNode.Left)
		}

		if currentNode.Right != nil {
			visitedQueue.Enqueue(currentNode.Right)
		}

		if levelCount == 0 {
			// Append the last element of the BFS level. That is the right-most node.
			finalOutput = append(finalOutput, levelArray[len(levelArray)-1])

			// Reset the level
			levelCount = visitedQueue.Size()
			levelArray = make([]int, 0)
		}
	}

	return finalOutput
}
