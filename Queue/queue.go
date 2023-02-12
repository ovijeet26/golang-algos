package Queue

import (
	"github.com/ovijeet/go/algos/stack"
)

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// Leet code link -> https://leetcode.com/problems/implement-queue-using-stacks/
type MyQueue struct {
	mainStack stack.IntStack
	sideStack stack.IntStack
}

func Constructor() MyQueue {

	return MyQueue{}
}

func (this *MyQueue) Push(x int) {

	this.mainStack.Push(x)
}

func (this *MyQueue) Pop() int {

	for this.mainStack.Size() > 0 {

		poppedItem := this.mainStack.Pop()

		this.sideStack.Push(poppedItem)
	}

	finalPopedItem := this.sideStack.Pop()

	for this.sideStack.Size() > 0 {

		poppedItem := this.sideStack.Pop()

		this.mainStack.Push(poppedItem)
	}

	return finalPopedItem
}

func (this *MyQueue) Peek() int {

	for this.mainStack.Size() > 0 {

		poppedItem := this.mainStack.Pop()

		this.sideStack.Push(poppedItem)
	}

	peekedItem := this.sideStack.Pop()

	this.mainStack.Push(peekedItem)

	for this.sideStack.Size() > 0 {

		poppedItem := this.sideStack.Pop()

		this.mainStack.Push(poppedItem)
	}

	return peekedItem

}

func (this *MyQueue) Empty() bool {

	return this.mainStack.Size() == 0
}

// Regular queue implementation

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(data interface{}) {

	q.items = append(q.items, data)
}

func (q *Queue) Dequeue() interface{} {

	toRemove := q.items[0]

	// Remove first element from slice
	q.items = q.items[1:]
	return toRemove
}

func (q *Queue) Size() int {

	return len(q.items)
}
