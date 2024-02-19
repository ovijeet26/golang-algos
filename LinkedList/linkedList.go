package linkedlist

import "fmt"

type ListNode struct {
	data int
	Next *ListNode
}

type LinkedList struct {
	head   *ListNode
	length int
}

func NewNode(data int) ListNode {

	return ListNode{data: data, Next: nil}
}

// Add a new ListNode at the beginning of a linked list
func (l *LinkedList) Prepend(newNode *ListNode) {

	newNode.Next = l.head
	l.head = newNode

	l.length++
}

func (l *LinkedList) DeleteWithValue(searchKey int) {

	previousToDelete := l.head

	if previousToDelete == nil {
		return
	}

	if previousToDelete.data == searchKey {
		l.head = previousToDelete.Next
		l.length--
		return
	}

	for previousToDelete.Next.Next != nil {

		if previousToDelete.Next.data == searchKey {

			afterDelete := previousToDelete.Next.Next
			previousToDelete.Next = afterDelete
			l.length--
		} else {

			previousToDelete = previousToDelete.Next
		}
	}
}

func (l LinkedList) PrintList() {

	ptrNode := l.head

	if ptrNode == nil {
		return
	}

	for ptrNode.Next != nil {

		fmt.Printf("%v -> ", ptrNode.data)
		ptrNode = ptrNode.Next
	}

	fmt.Print("nil \n")
}

func (l *LinkedList) Reverse() {

	currentNode := l.head
	nextNode := currentNode
	var prevNode *ListNode
	prevNode = nil

	// Take care of border cases
	if currentNode == nil {
		return
	} else if currentNode.Next == nil {
		return
	}

	for currentNode != nil {

		nextNode = currentNode.Next
		currentNode.Next = prevNode

		prevNode = currentNode
		currentNode = nextNode
	}

	l.head = prevNode
}

//1->2->3->4-> nil
func RecursiveReverse(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {

		return head
	}

	nodePtr := RecursiveReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return nodePtr
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//1->2->[3->4->5]->6->nil

//5->4->3->nil

func ReverseBetween(head *ListNode, left int, right int) *ListNode {

	counter := 1
	var startNode *ListNode
	var secondTailNode *ListNode
	var secondHead *ListNode
	var endNode *ListNode

	currentNode := head
	var nextNode *ListNode

	var prevNode *ListNode
	prevNode = nil

	for currentNode != nil {

		if counter < left-1 {

			currentNode = currentNode.Next
			counter++
			continue
		}

		if counter == left-1 {

			startNode = currentNode
		}

		if counter == left {

			secondTailNode = currentNode
		}

		if counter == right {

			secondHead = currentNode
		}

		if counter == right+1 {

			endNode = currentNode
		}

		if counter >= left && counter <= right {

			nextNode = currentNode.Next
			currentNode.Next = prevNode
			prevNode = currentNode
			currentNode = nextNode
		}

		if counter > right+1 {

			break
		}
		currentNode = currentNode.Next
		counter++
	}

	startNode.Next = secondHead
	secondTailNode.Next = endNode

	return head
}

// Leet code link -> https://leetcode.com/problems/linked-list-cycle-ii/
func DetectCycleUsingSet(head *ListNode) *ListNode {

	// Make-shift set data type in Golang.
	// Empty struct takes 0 bits in Golang.
	seenNodes := make(map[*ListNode]struct{}, 0)
	present := struct{}{}

	currentNode := head

	for currentNode != nil {

		_, isPresent := seenNodes[currentNode]

		if isPresent {
			break
		}

		seenNodes[currentNode] = present
		currentNode = currentNode.Next
	}

	return currentNode
}

// Leet code link -> https://leetcode.com/problems/linked-list-cycle-ii/
func DetectCycle(head *ListNode) *ListNode {

	// Use Floyd's hare and tortoise algorithm

	slowPtr := head
	fastPtr := head

	isCycleDetected := false

	for fastPtr != nil && fastPtr.Next != nil && fastPtr.Next.Next != nil {

		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

		if slowPtr == fastPtr {
			isCycleDetected = true
			break
		}
	}

	if isCycleDetected {

		beginingPtr := head
		secondPointer := fastPtr

		for beginingPtr != secondPointer {

			beginingPtr = beginingPtr.Next
			secondPointer = secondPointer.Next
		}
		return beginingPtr

	}
	return nil
}
