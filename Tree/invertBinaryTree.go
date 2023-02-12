package tree

//Leet code Link -> https://leetcode.com/problems/invert-binary-tree/
func invertTree(root *TreeNode) *TreeNode {

	// if root is nil, return
	if root == nil {
		return nil
	}

	leftChild := invertTree(root.Left)
	rightChild := invertTree(root.Right)

	root.Left = rightChild
	root.Right = leftChild

	return root
}
