package tree

import "math"

//Leet code link -> https://leetcode.com/problems/validate-binary-search-tree/
func IsValidBST(root *TreeNode) bool {

	isValid := dfs(root, int(math.Inf(-1)), int(math.Inf(0)))

	return isValid
}

func dfs(node *TreeNode, min int, max int) bool {

	if node.Value <= min || node.Value >= max {

		return false
	}

	if node.Left != nil {

		if dfs(node.Left, min, node.Value) == false {

			return false
		}
	}

	if node.Right != nil {

		if dfs(node.Right, node.Value, max) == false {

			return false
		}
	}

	return true
}
