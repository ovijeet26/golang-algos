package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    if head==nil{
        return nil
    }

    dummyNode:=&ListNode{
        Next:head,
    }

    left:=dummyNode
    right:=head
    for n>0 {
        right=right.Next
        n--
    }


    for right!=nil{
        right=right.Next
        left=left.Next
    }


    left.Next=left.Next.Next
    return dummyNode.Next
}