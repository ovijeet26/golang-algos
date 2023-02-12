package search

// Leet code link -> https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRangeOptimized(nums []int, target int) []int {

	size := len(nums)

	if size == 0 {
		return []int{-1, -1}
	}

	// Search any index of the searchKey
	searchIndex := binarySearchRec(nums, target, 0, size-1)

	if searchIndex == -1 {

		return []int{-1, -1}
	}

	// Keep binary searching in both directions, till we don't get any searchkey.
	// The last element before we don't find anything are the boundaries.
	frontCrawler := searchIndex
	backCrawler := searchIndex

	temp := searchIndex

	for frontCrawler != -1 {
		temp = frontCrawler
		frontCrawler = binarySearchRec(nums, target, frontCrawler+1, size-1)
	}
	frontCrawler = temp

	temp = backCrawler

	for backCrawler != -1 {
		temp = backCrawler
		backCrawler = binarySearchRec(nums, target, 0, backCrawler-1)
	}

	backCrawler = temp

	return []int{backCrawler, frontCrawler}
}

func searchRange(nums []int, target int) []int {

	size := len(nums)
	searchIndex := binarySearchRec(nums, target, 0, size-1)

	if searchIndex == -1 {

		return []int{-1, -1}
	}
	var frontCrawler int
	var backCrawler int

	if nums[searchIndex+1] == target {
		frontCrawler = searchIndex + 1
	} else {
		frontCrawler = searchIndex
	}

	if nums[searchIndex-1] == target {
		backCrawler = searchIndex - 1
	} else {
		backCrawler = searchIndex
	}

	for frontCrawler+1 < size && nums[frontCrawler+1] == target {

		frontCrawler++
	}

	for backCrawler-1 > 0 && nums[backCrawler-1] == target {
		backCrawler--
	}

	return []int{backCrawler, frontCrawler}
}

func binarySearchRec(data []int, searchKey int, left int, right int) int {

	if left >= right {

		return -1
	}

	mid := (left + right) / 2

	if searchKey == data[mid] {

		return mid
	} else if searchKey > data[mid] {

		return binarySearchRec(data, searchKey, mid+1, right)
	} else {
		return binarySearchRec(data, searchKey, left, mid-1)
	}
}
