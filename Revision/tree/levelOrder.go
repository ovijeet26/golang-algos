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
func levelOrder(root *TreeNode) [][]int {

	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := Queue{}

	queue.Enqueue(root)
	levelArray := make([]int, 0)
	levelElems := queue.Size()

	for queue.Size() != 0 {
		currenNode := queue.Dequeue()
		levelArray = append(levelArray, currenNode.Val)
		levelElems--
		if currenNode.Left != nil {
			queue.Enqueue(currenNode.Left)
		}
		if currenNode.Right != nil {
			queue.Enqueue(currenNode.Right)
		}

		if levelElems == 0 {
			result = append(result, levelArray)
			levelArray = make([]int, 0)
			levelElems = queue.Size()
		}
	}
	return result
}

type Queue struct {
	items []*TreeNode
}

func (q *Queue) Enqueue(item *TreeNode) {

	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() *TreeNode {

	toRemove := q.items[0]

	// Remove first element from slice
	q.items = q.items[1:]
	return toRemove
}

func (q *Queue) Size() int {

	return len(q.items)
}

func LevelOrderCaller() {

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

	width := levelOrder(&a)

	fmt.Print(width)

	//			1
	//		2		3
	//	4				5
	//6
}
