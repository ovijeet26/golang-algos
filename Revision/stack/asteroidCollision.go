package stack

import "fmt"

func asteroidCollision(asteroids []int) []int {

	stack := IntStack{}
	index := 0
	for index < len(asteroids) {

		a := asteroids[index]
		if stack.Len() == 0 {
			stack.Push(a)
			index++
			continue
		}

		stackTop := stack.Peek()

		// if the stack top element is +ve and the current element is -ve. i.e. there is a chance for collision
		if stackTop < 0 && a < 0 {

			diff := absoluteDifference(stackTop, a)
			// if stack top is greater in value than the current asteroid
			if diff > 0 {
				index++
				continue
			}
			// in case both are equal in value, we remove boh of them
			if diff == 0 {
				stack.Pop()
				index++
				continue
			}
			// if stack top is lesser in value than the current asteroid
			if diff < 0 {
				stack.Pop()
				continue
			}
		}

		stack.Push(a)
		index++
	}
	return stack.ArrrayValues()
}

func absoluteDifference(a int, b int) int {

	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	return a - b
}

func AsteroidCaller() {

	fmt.Println(asteroidCollision([]int{10, 2, -5}))
}
