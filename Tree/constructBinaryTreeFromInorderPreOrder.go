package tree

// Leet code link -> https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func buildTree(preorder []int, inorder []int) *TreeNode {

	root := buildTreeRec(preorder, inorder)
	return root
}

// Build the tree recursively
func buildTreeRec(preOrder []int, inOrder []int) *TreeNode {

	// If either the inorder or preorder arary is empty, then we have reached a leaf node's child
	if len(preOrder) == 0 || len(inOrder) == 0 {

		return nil
	}

	node := new(TreeNode)

	// The current node will always be the first element of the preorder traversal
	node.Value = preOrder[0]

	// We want to locate the position of the current node in the inorder traversal.
	// That will give us the middle elemnt, left of which will be the left subtree,
	// and right of which will be the right sub tree.
	inOrderMid := findIndex(inOrder, node.Value)

	node.Left = buildTreeRec(preOrder[1:inOrderMid+1], inOrder[:inOrderMid])
	node.Right = buildTreeRec(preOrder[inOrderMid+1:], inOrder[inOrderMid+1:])

	return node
}

func findIndex(arr []int, val int) int {

	for i := 0; i < len(arr); i++ {

		if arr[i] == val {

			return i
		}
	}

	return -1
}
