package tree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {

	result := make([]int, 0)
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		result = append(result, root.Val)
		return result
	}

	queue := Queue{}
	queue.Enqueue(root)
	levelCount := 1

	levelArray := make([]int, 0)
	for queue.Size() != 0 {

		current := queue.Dequeue()
		if root.Left != nil {
			queue.Enqueue(root.Left)
		}
		if root.Right != nil {
			queue.Enqueue(root.Right)
		}
		levelArray = append(levelArray, current.Val)

		levelCount--
		if levelCount == 0 {
			result = append(result, levelArray[len(levelArray)-1])
			levelCount = queue.Size()
		}

	}
	return result
}

func RightSideViewCaller() {

	a := TreeNode{Val: 1}
	b := TreeNode{Val: 2}
	c := TreeNode{Val: 3}
	d := TreeNode{Val: 4}
	e := TreeNode{Val: 5}
	f := TreeNode{Val: 6}

	a.Left = &b
	a.Right = &c

	b.Left = &d
	c.Right = &e

	d.Left = &f

	res := rightSideView(&a)

	fmt.Print(res)

	//			1
	//		2		3
	//	4				5
	//6
}
