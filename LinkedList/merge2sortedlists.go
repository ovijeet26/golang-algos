package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     data int
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
	if ptr1.data < ptr2.data {
		finalHead = &ListNode{
			data: ptr1.data,
		}
		ptr1 = ptr1.Next
	} else {
		finalHead = &ListNode{
			data: ptr2.data,
		}
		ptr2 = ptr2.Next
	}
	ctr := finalHead
	for ptr1 != nil && ptr2 != nil {

		if ptr1.data < ptr2.data {
			ctr.Next = &ListNode{
				data: ptr1.data,
			}
			ptr1 = ptr1.Next
		} else {
			ctr.Next = &ListNode{
				data: ptr2.data,
			}
			ptr2 = ptr2.Next
		}

		ctr = ctr.Next
	}

	for ptr1 != nil {
		ctr.Next = &ListNode{
			data: ptr1.data,
		}
		ptr1 = ptr1.Next
		ctr = ctr.Next
	}

	for ptr2 != nil {
		ctr.Next = &ListNode{
			data: ptr2.data,
		}
		ptr2 = ptr2.Next
		ctr = ctr.Next
	}

	return finalHead
}

func MergeListCaller() {

	listA := &ListNode{
		data: 1,
	}
	listB := &ListNode{
		data: 2,
	}
	listC := &ListNode{
		data: 4,
		Next: nil,
	}

	listX := &ListNode{
		data: 1,
	}
	listY := &ListNode{
		data: 3,
	}
	listZ := &ListNode{
		data: 4,
		Next: nil,
	}

	listA.Next = listB
	listB.Next = listC

	listX.Next = listY
	listY.Next = listZ
	mergeTwoLists(listA, listX)

}
