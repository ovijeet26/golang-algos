package revision

import (
	"fmt"
)

func removeDuplicates(nums []int) int {

	fixedPtr, movingPtr := 0, 1

	for movingPtr < len(nums) {

		if nums[fixedPtr] == nums[movingPtr] {
			movingPtr++
			continue
		}

		swap(nums, fixedPtr+1, movingPtr)
		fixedPtr++
		movingPtr++
	}

	return fixedPtr + 1
}

func swap(nums []int, i int, j int) {

	temp := nums[j]
	nums[j] = nums[i]
	nums[i] = temp
}

// Using hashmap
func removeDuplicates2(nums []int) int {
	cacheMap := make(map[int]bool)
	lastUniqueIndex := 0
	for _, v := range nums {
		_, isPresent := cacheMap[v]
		if !isPresent {
			cacheMap[v] = true
			nums[lastUniqueIndex] = v
			lastUniqueIndex++
		}
	}
	return lastUniqueIndex
}

func InvokeRemoveDuplicates() {

	k := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})

	fmt.Print(k)
}
