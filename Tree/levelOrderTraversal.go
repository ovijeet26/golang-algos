package tree

import queue "github.com/ovijeet/go/algos/Queue"

// Leet code link -> https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {

	levelArray := make([][]int, 0)
	level1DArray := make([]int, 0)

	if root == nil {

		return levelArray
	}

	visitedQueue := queue.Queue{}

	visitedQueue.Enqueue(root)

	// Initially only the root is in the queue in level 0.
	levelNodeCounter := 1
	for visitedQueue.Size() > 0 {

		currentNode := visitedQueue.Dequeue().(*TreeNode)
		levelNodeCounter--

		level1DArray = append(level1DArray, currentNode.Value)

		if currentNode.Left != nil {
			visitedQueue.Enqueue(currentNode.Left)
		}

		if currentNode.Right != nil {
			visitedQueue.Enqueue(currentNode.Right)
		}

		if levelNodeCounter == 0 {

			levelArray = append(levelArray, level1DArray)
			// Reset the 1D array
			level1DArray = make([]int, 0)

			// Reinitialize level node counter witht he current contents of the queue.
			// This denotes the next level.
			levelNodeCounter = visitedQueue.Size()
		}
	}

	return levelArray
}
