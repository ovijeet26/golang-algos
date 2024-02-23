package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	ptr1, ptr2 := list1, list2

	if ptr1 == nil {
		return ptr2
	}

	if ptr2 == nil {
		return ptr1
	}
	var finalHead *ListNode
	if ptr1.Val < ptr2.Val {
		finalHead = &ListNode{
			Val: ptr1.Val,
		}
		ptr1 = ptr1.Next
	} else {
		finalHead = &ListNode{
			Val: ptr2.Val,
		}
		ptr2 = ptr2.Next
	}
	ctr := finalHead
	for ptr1 != nil && ptr2 != nil {

		if ptr1.Val < ptr2.Val {
			ctr.Next = &ListNode{
				Val: ptr1.Val,
			}
			ptr1 = ptr1.Next
		} else {
			ctr.Next = &ListNode{
				Val: ptr2.Val,
			}
			ptr2 = ptr2.Next
		}

		ctr = ctr.Next
	}

	for ptr1 != nil {
		ctr.Next = &ListNode{
			Val: ptr1.Val,
		}
		ptr1 = ptr1.Next
		ctr = ctr.Next
	}

	for ptr2 != nil {
		ctr.Next = &ListNode{
			Val: ptr2.Val,
		}
		ptr2 = ptr2.Next
		ctr = ctr.Next
	}

	return finalHead
}

func MergeListCaller() {

	listA := &ListNode{
		Val: 1,
	}
	listB := &ListNode{
		Val: 2,
	}
	listC := &ListNode{
		Val:  4,
		Next: nil,
	}

	listX := &ListNode{
		Val: 1,
	}
	listY := &ListNode{
		Val: 3,
	}
	listZ := &ListNode{
		Val:  4,
		Next: nil,
	}

	listA.Next = listB
	listB.Next = listC

	listX.Next = listY
	listY.Next = listZ
	mergeTwoLists(listA, listX)

}
