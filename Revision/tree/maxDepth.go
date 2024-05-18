package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

	return maxDepthRec(root, 0)
}

func maxDepthRec(node *TreeNode, count int) int {

	if node == nil {
		return count
	}
	count++
	return max(maxDepthRec(node.Left, count), maxDepthRec(node.Right, count))
}

func max(a int, b int) int {

	if a > b {
		return a
	}
	return b
}
