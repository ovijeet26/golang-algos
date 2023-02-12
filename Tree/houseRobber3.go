package tree

// Leet codde link -> https://leetcode.com/problems/house-robber-iii/
func rob(root *TreeNode) int {

	return max(dfsCost(root))
}

func dfsCost(node *TreeNode) (int, int) {

	if node == nil {

		return 0, 0
	}

	// Calculate the cost of with root and without root for LST and RST
	leftWithRoot, leftWithoutRoot := dfsCost(node.Left)
	rightWithRoot, rightWithoutRoot := dfsCost(node.Right)

	// Calculate the node with root of current node
	nodeCostWithRoot := node.Value + leftWithoutRoot + rightWithoutRoot

	// Calculate the node without root of current node
	nodeCostWithoutRoot := max(leftWithRoot, leftWithoutRoot) + max(rightWithRoot, rightWithoutRoot)

	return nodeCostWithRoot, nodeCostWithoutRoot
}
