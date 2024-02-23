package linkedlist

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	// Mege 2 side linked lists till we have only 1 left (kind of like merge sort)
	for len(lists) > 1 {

		newList := make([]*ListNode, 0)
		for i := 0; i < len(lists); i += 2 {

			l1 := lists[i]
			var l2 *ListNode
			if i+1 < len(lists) {
				l2 = lists[i+1]
			} else {
				l2 = nil
			}

			mergedList := merge2Lists(l1, l2)
			newList = append(newList, mergedList)
		}
		lists = newList
	}
	return lists[0]
}

func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	ptr1 := l1
	ptr2 := l2

	var newHead *ListNode
	if ptr1.Val < ptr2.Val {
		newHead = ptr1
		ptr1 = ptr1.Next
		newHead.Next = nil
	} else {
		newHead = ptr2
		ptr2 = ptr2.Next
		newHead.Next = nil
	}

	newPtr := newHead

	for ptr1 != nil && ptr2 != nil {

		if ptr1.Val < ptr2.Val {
			newPtr.Next = ptr1
			ptr1 = ptr1.Next
			newPtr = newPtr.Next
			newPtr.Next = nil
		} else {
			newPtr.Next = ptr2
			ptr2 = ptr2.Next
			newPtr = newPtr.Next
			newPtr.Next = nil
		}
	}

	if ptr1 != nil {
		newPtr.Next = ptr1
	}

	if ptr2 != nil {
		newPtr.Next = ptr2
	}
	return newHead
}

func MegeListCaller() {

	lists := []*ListNode{
		CreateList([]int{1, 4, 5}),
		CreateList([]int{1, 3, 4}),
		CreateList([]int{2, 6}),
	}
	fmt.Print(mergeKLists(lists))
}

func CreateList(arr []int) *ListNode {

	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	ptr := head

	for i := 1; i < len(arr); i++ {

		ptr.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		ptr = ptr.Next
	}

	return head
}
