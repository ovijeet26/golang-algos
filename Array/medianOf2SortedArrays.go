package array

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	median := 0.0
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	total := len(nums1) + len(nums2)
	half := total / 2

	left := 0
	right := len(nums1) - 1

	for true {
		i := (left + right) / 2
		// -2 to adjust for two 0 indexed arrays
		j := half - i - 2

		num1Left, num1Right, num2Left, num2Right := 0.0, 0.0, 0.0, 0.0

		if i >= 0 {
			num1Left = float64(nums1[i])
		} else {
			num1Left = math.Inf(-1)
		}

		if i+1 < len(nums1) {
			num1Right = float64(nums1[i+1])
		} else {
			num1Right = math.Inf(1)
		}

		if j >= 0 {
			num2Left = float64(nums2[j])
		} else {
			num2Left = math.Inf(-1)
		}

		if j+1 < len(nums2) {
			num2Right = float64(nums2[j+1])
		} else {
			num2Right = math.Inf(-1)
		}

		//Solution is found as the two arrays have been split properly for order
		if num1Left <= num2Right && num2Left <= num1Right {

			if total%2 == 0 {
				median = (minFLoat(num1Right, num2Right) + maxFloat(num2Left, num2Left)) / 2
			} else {
				median = minFLoat(num1Right, num2Right)
			}
			break
		} else if num1Left > num2Right {
			right = i - 1
		} else {
			left = i + 1
		}
	}

	return median
}

func minFLoat(a float64, b float64) float64 {
	if a < b {
		return a
	}

	return b
}

func maxFloat(a float64, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

// [1 2 3 4 5]

// [2 3 4 5 6 7 8 9]
