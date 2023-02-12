package tree

import (
	"fmt"
	"math"

	stack "github.com/ovijeet/go/algos/stack"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Depth-First values
func DepthFirstValuesIterative(root *TreeNode) {

	if root == nil {

		return
	}

	visitedStack := stack.GenericStack{}
	visitedStack.Push(root)

	for visitedStack.Size() > 0 {

		poppedNode, _ := visitedStack.Pop().(*TreeNode)

		fmt.Println(poppedNode.Value)

		if poppedNode.Left != nil {

			visitedStack.Push(poppedNode.Left)
		}

		if poppedNode.Right != nil {

			visitedStack.Push(poppedNode.Right)
		}
	}

}

func DepthFirstValuesPreOrder(root *TreeNode) {

	if root == nil {

		return
	}

	fmt.Println(root.Value)

	DepthFirstValuesPreOrder(root.Left)
	DepthFirstValuesPreOrder(root.Right)
}

func DepthFirstValuesPostOrder(root *TreeNode) {

	if root == nil {

		return
	}

	DepthFirstValuesPostOrder(root.Left)
	DepthFirstValuesPostOrder(root.Right)
	fmt.Println(root.Value)
}

func DepthFirstValuesInOrder(root *TreeNode) {

	if root == nil {

		return
	}

	DepthFirstValuesInOrder(root.Left)
	fmt.Println(root.Value)
	DepthFirstValuesInOrder(root.Right)
}

// Leet code link -> https://leetcode.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	return depthCount(root, 0)
}

func depthCount(node *TreeNode, maxCount int) int {

	if node == nil {

		return maxCount
	}

	maxCount++

	return int(math.Max(float64(depthCount(node.Left, maxCount)), float64(depthCount(node.Right, maxCount))))
}

func depthCountSimple(node *TreeNode, maxCount int) int {

	if node == nil {

		return maxCount
	}

	maxCount++

	leftCount := depthCount(node.Left, maxCount)
	rightCount := depthCount(node.Right, maxCount)

	if rightCount > leftCount {

		return rightCount
	}

	return leftCount
}
