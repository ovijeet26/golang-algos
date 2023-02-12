package array

// Leet code link -> https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {

	size := len(nums)

	if size == 0 {
		return []int{-1, -1}
	}

	searchIndex := binarySearchRec(nums, target, 0, size-1)

	if searchIndex == -1 {

		return []int{-1, -1}
	}

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

func binarySearchRec(data []int, searchKey int, left int, right int) int {

	if left > right {

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
