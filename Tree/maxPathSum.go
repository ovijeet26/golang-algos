package tree

// Leet code link -> https://leetcode.com/problems/binary-tree-maximum-path-sum/
func maxPathSum(root *TreeNode) int {

	currentMaxPath := root.Value

	pathDfs(root, &currentMaxPath)

	return currentMaxPath
}

// Return the Max path sum WITHOUT the split
func pathDfs(node *TreeNode, currentMaxPath *int) int {

	if node == nil {
		return 0
	}

	// Calculate the MAX for the left and right sub-tree

	leftMax := pathDfs(node.Left, currentMaxPath)
	rightMax := pathDfs(node.Right, currentMaxPath)

	// We don't want to take -ve max as left or right subtree.
	// We can get this scenario when both the left and right subtree path sum are -ve.

	leftMax = max(leftMax, 0)
	rightMax = max(rightMax, 0)

	// Max at current iterating node is the sum of it's value and BOTH lefta nd right sub-tree.

	maxAtCurrentNode := node.Value + leftMax + rightMax

	// Calculate the MAX at Node WITH split
	*currentMaxPath = max(*currentMaxPath, maxAtCurrentNode)

	return node.Value + max(leftMax, rightMax)
}

func max(a int, b int) int {

	if a > b {
		return a
	}
	return b
}
