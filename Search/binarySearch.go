package search

func BinarySearch(data []int, searchKey int) int {

	left := 0
	right := len(data) - 1
	mid := 0

	for left < right {

		mid = (left + right) / 2

		if searchKey == data[mid] {
			return mid
		} else if searchKey < data[mid] {

			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func BinarySearchRecursive(data []int, left int, right int, searchKey int) int {

	mid := (left + right) / 2

	if left >= right {
		return -1
	}

	if data[mid] == searchKey {
		return mid
	}

	if data[mid] > searchKey {

		return BinarySearchRecursive(data, left, mid-1, searchKey)
	} else {

		return BinarySearchRecursive(data, mid+1, right, searchKey)
	}

}
