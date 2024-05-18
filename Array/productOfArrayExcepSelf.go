package array

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {

	res := make([]int, len(nums))

	res[0] = 1

	// Compute Prefix Array
	for i := 1; i < len(nums); i++ {

		res[i] = res[i-1] * nums[i-1]
	}

	// Compute suffix plus th final array
	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {

		res[i] = res[i] * suffix
		suffix = suffix * nums[i]
	}
	return res
}

func productExceptSelf_LessOptimized(nums []int) []int {

	res := make([]int, len(nums))
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	prefix[0] = 1
	suffix[len(nums)-1] = 1

	// Compute Prefix Array
	for i := 1; i < len(nums); i++ {

		prefix[i] = prefix[i-1] * nums[i-1]
	}

	// Compute Suffix Array
	for i := len(nums) - 2; i >= 0; i-- {

		suffix[i] = suffix[i+1] * nums[i+1]
	}

	// Compute final result
	for i := 0; i < len(nums); i++ {

		res[i] = prefix[i] * suffix[i]
	}

	return res
}

func Area(shape interface{}) {
	n := shape.(int)
	fmt.Print(n)
	m, ok := shape.(int)
	fmt.Print(m, ok)
}

// [1,2,3,4,5]

// [0,1,2,6,24]
// [120,60,20,5,0]

// [1,2,3]

// p [1,1,2]
// s [6,3,1]

// r [6,3,2]

// r[  ]
