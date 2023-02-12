package tree

import "math"

// Leet code link -> https://leetcode.com/problems/count-complete-tree-nodes/
func CountNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil {

		return 1
	}

	height := getHeightOfBinaryTree(root)

	// Total nodes till level n-1 -> (2^n) -1
	upperSizeOfTree := totalNodesOfTreeAtLevelN(height)

	left := 0
	right := totalNodesOfTreeAtLevelN(height)

	for left < right {

		// round up the mid value unlike conventional binary search
		mid := caluculateMidCeil(left, right)

		if isNodeExist(root, mid, height) {

			left = mid
		} else {

			right = mid - 1
		}
	}

	// Since left is 0 indexed, size will be + 1
	nodeAtLastLevel := left + 1

	numberOfNodes := upperSizeOfTree + nodeAtLastLevel

	return numberOfNodes
}

func getHeightOfBinaryTree(node *TreeNode) int {

	height := 0
	for node.Left != nil {

		node = node.Left
		height++
	}
	return height
}

func isNodeExist(root *TreeNode, indexToFind int, height int) bool {

	left := 0
	right := totalNodesOfTreeAtLevelN(height)
	counter := 0
	for counter < height {

		mid := caluculateMidCeil(left, right)
		if indexToFind >= mid {

			root = root.Right
			left = mid
		} else {

			root = root.Left
			right = mid - 1
		}
		counter++
	}

	return root != nil
}

func totalNodesOfTreeAtLevelN(height int) int {

	// Total nodes of a complete tree at any level is 2^n -1

	return int(math.Pow(2, float64(height))) - 1
}

func caluculateMidCeil(left int, right int) int {

	ceilValue := math.Ceil(float64((left + right)) / 2)
	return int(ceilValue)
}

// func countNodes2(root *TreeNode) int {
// 	nodeCount := 0

// 	dfsCountPointer(root, &nodeCount)
// 	return nodeCount
// }

// func dfsCountPointer(node *TreeNode, count *int) {

// 	if node == nil {
// 		return
// 	}

// 	*count++

// 	dfsCountPointer(node.Left, count)
// 	dfsCountPointer(node.Right, count)
// }

// func dfsCount(node *TreeNode) int {

// 	if node == nil {
// 		return 0
// 	}

// 	return dfsCount(node.Left) + dfsCount(node.Right) + 1
// }
