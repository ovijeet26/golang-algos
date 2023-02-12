package hard

import "fmt"

// Leet code link -> https://leetcode.com/problems/race-car/
func racecar(target int) int {

	// TODO INCOMPLETE
	return actions(0, 1, target, 0)
}

func actions(position int, speed int, target int, move int) int {

	if position == target {

		return move
	}

	if (position < target && speed > 0) || (position > target && speed < 0) {

		move = actions(position+speed, speed*2, target, move+1)
		return move
	}

	if (position > target && speed > 0) || (position < target && speed < 0) {

		if speed < 0 {

			speed = 1
		} else {
			speed = -1
		}
		move = actions(position, speed, target, move+1)
		return move
	}

	return move
}

func RaceCarCaller() {

	fmt.Println(racecar(4))
}
