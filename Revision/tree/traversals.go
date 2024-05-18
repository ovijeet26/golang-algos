package tree

import "fmt"

func PostOrder(root *TreeNode) {

	if root == nil {
		return
	}

	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Println(root.Val)
}

func PreOrder(root *TreeNode) {

	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func InOrder(root *TreeNode) {

	if root == nil {
		return
	}

	InOrder(root.Left)
	fmt.Println(root.Val)
	InOrder(root.Right)
}

func TraversalCaller() {

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

	PostOrder(&a)
	fmt.Println("..............")
	PreOrder(&a)
	fmt.Println("..............")
	InOrder(&a)

	//			1
	//		2		3
	//	4				5
	//6
}
