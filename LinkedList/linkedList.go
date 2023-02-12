package linkedlist

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head   *ListNode
	length int
}

func NewNode(data int) ListNode {

	return ListNode{data: data, next: nil}
}

// Add a new ListNode at the beginning of a linked list
func (l *LinkedList) Prepend(newNode *ListNode) {

	newNode.next = l.head
	l.head = newNode

	l.length++
}

func (l *LinkedList) DeleteWithValue(searchKey int) {

	previousToDelete := l.head

	if previousToDelete == nil {
		return
	}

	if previousToDelete.data == searchKey {
		l.head = previousToDelete.next
		l.length--
		return
	}

	for previousToDelete.next.next != nil {

		if previousToDelete.next.data == searchKey {

			afterDelete := previousToDelete.next.next
			previousToDelete.next = afterDelete
			l.length--
		} else {

			previousToDelete = previousToDelete.next
		}
	}
}

func (l LinkedList) PrintList() {

	ptrNode := l.head

	if ptrNode == nil {
		return
	}

	for ptrNode.next != nil {

		fmt.Printf("%v -> ", ptrNode.data)
		ptrNode = ptrNode.next
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
	} else if currentNode.next == nil {
		return
	}

	for currentNode != nil {

		nextNode = currentNode.next
		currentNode.next = prevNode

		prevNode = currentNode
		currentNode = nextNode
	}

	l.head = prevNode
}

//1->2->3->4-> nil
func RecursiveReverse(head *ListNode) *ListNode {

	if head == nil || head.next == nil {

		return head
	}

	nodePtr := RecursiveReverse(head.next)
	head.next.next = head
	head.next = nil
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

			currentNode = currentNode.next
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

			nextNode = currentNode.next
			currentNode.next = prevNode
			prevNode = currentNode
			currentNode = nextNode
		}

		if counter > right+1 {

			break
		}
		currentNode = currentNode.next
		counter++
	}

	startNode.next = secondHead
	secondTailNode.next = endNode

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
		currentNode = currentNode.next
	}

	return currentNode
}

// Leet code link -> https://leetcode.com/problems/linked-list-cycle-ii/
func DetectCycle(head *ListNode) *ListNode {

	// Use Floyd's hare and tortoise algorithm

	slowPtr := head
	fastPtr := head

	isCycleDetected := false

	for fastPtr != nil && fastPtr.next != nil && fastPtr.next.next != nil {

		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next

		if slowPtr == fastPtr {
			isCycleDetected = true
			break
		}
	}

	if isCycleDetected {

		beginingPtr := head
		secondPointer := fastPtr

		for beginingPtr != secondPointer {

			beginingPtr = beginingPtr.next
			secondPointer = secondPointer.next
		}
		return beginingPtr

	}
	return nil
}
