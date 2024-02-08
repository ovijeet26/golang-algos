package main

import (
	"fmt"
	"regexp"
	"strings"

	array "github.com/ovijeet/go/algos/Array"
	dynamicprogramming "github.com/ovijeet/go/algos/DynamicProgramming"
	hard "github.com/ovijeet/go/algos/LeetCode/Hard"
	revision "github.com/ovijeet/go/algos/Revision"
	revisionStack "github.com/ovijeet/go/algos/Revision/Stack"
	revisionString "github.com/ovijeet/go/algos/Revision/String"
	tree "github.com/ovijeet/go/algos/Tree"
	"github.com/ovijeet/go/algos/backtracking"
	"github.com/ovijeet/go/algos/graph"
	"github.com/ovijeet/go/algos/interfacedesign"
	"github.com/ovijeet/go/algos/matrix"
	"github.com/ovijeet/go/algos/search"
	"github.com/ovijeet/go/algos/sort"
	str "github.com/ovijeet/go/algos/string"
)

func main() {
	revisionStack.DecodeCaller()
	revisionStack.ValidCaller()
	revisionString.LongestCharRCaller()
	//revisionString.TypedOutStringCaller()
	revision.MergeSortCaller()
	revision.ThreesumCaller()
	revision.FnLCaller()
	revision.BSCaller()
	revision.LongestCaller()
	revision.QSCaller()
	revision.InvokeRemoveDuplicates()
	array.RemoveDuplicate2Caller()
	array.RemoveDuplicateCaller()
	array.PeakElementCaller()
	hard.RaceCarCaller()
	array.SlidingWindowMaxCaller()
	dynamicprogramming.CoinChangeCaller()
	dynamicprogramming.CountConstructCaller()
	dynamicprogramming.CanConstructCaller()
	dynamicprogramming.BestSumCaller()
	dynamicprogramming.HowSumCaller()
	dynamicprogramming.CanSumCaller()
	dynamicprogramming.UniquePathCaller()
	matrix.WallsAndGateCaller()
	matrix.AreaIslandCaller()
	str.RemoveKDigitsCaller()
	array.ThreeSumCaller()
	graph.ValidTreeCaller()
	tree.KthElementCaller()
	array.SubArrayDiffCaller()
	matrix.WordsSearch2Caller()
	matrix.WordSearchCaller()
	hard.AlienCaller()
	hard.MaxPerformanceCaller()
	tree.WidthTreeCaller()
	str.SpaceRemover()
	str.CharacterReplacementCaller()
	matrix.SpiralMatrixCaller()
	backtracking.SodokuCaller()
	interfacedesign.MonarchyCaller()
	graph.NetworkDelayCaller()
	graph.CanFinishCaller()
	graph.NumOfMinutesCaller()
	matrix.SetZeroCaller()
	graph.SliceCaller()
	matrix.RottingOrangeCaller()
	matrix.NumIslandsCaller()
	matrixCaller()
	treeCaller()
	binarySearchCaller()
	quickSortCaller()
	mergeSortCaller()
	s := sanitizeString("A man, a plan, a canal: Panama")

	fmt.Println(s)

	lengthOfLongestSubstring("")
	backspaceCompare(",", "")
	nums := []int{3, 2, 4}
	target := 6
	res := TwoSum(nums, target)

	fmt.Print(res)
}

func binarySearchCaller() {

	data := []int{2, 3, 4, 5, 6, 1, 1, 65, 23, -21}

	fmt.Println(data)
	//sort.MergeSort(data, 0, len(data)-1)
	//index := search.BinarySearch(data, 66)
	index := search.BinarySearchRecursive(data, 0, len(data)-1, 65)

	fmt.Printf("Search index found : %v", index)
}

func matrixCaller() {

	matrix.BfsCaller()
	matrix.DfsCaller()
}
func treeCaller() {

	a := tree.TreeNode{Value: 1}
	b := tree.TreeNode{Value: 2}
	c := tree.TreeNode{Value: 3}
	d := tree.TreeNode{Value: 4}
	e := tree.TreeNode{Value: 5}
	f := tree.TreeNode{Value: 6}

	//..........
	// e.Left = &a
	// e.Right = &d
	// d.Left = &c
	// d.Right = &f
	// fmt.Print(tree.IsValidBST(&e))
	//..........
	a.Left = &b
	a.Right = &c

	b.Left = &d
	b.Right = &e

	c.Left = &f
	fmt.Println(tree.CountNodes(&a))
	fmt.Println(tree.RightSideViewDFSApproach(&a))
	// Iterative
	tree.DepthFirstValuesIterative(&a)
	fmt.Println(".............")
	// Recursive
	tree.DepthFirstValuesPreOrder(&a)
	fmt.Println(".............")
	tree.DepthFirstValuesPostOrder(&a)
	fmt.Println(".............")
	tree.DepthFirstValuesInOrder(&a)
	fmt.Println(".............")

}

func mergeSortCaller() {

	data := []int{2, 3, 4, 5, 6, 1, 21, 65, 23, -21}

	fmt.Println(data)
	sort.MergeSort(data, 0, len(data)-1)
	fmt.Println(data)
}

func quickSortCaller() {

	data := []int{2, 3, 4, 5, 6, 1, 21, 65, 23, -21}

	fmt.Println(data)
	sort.QuickSort(data, 0, len(data)-1)
	fmt.Println(data)
}

func TwoSum(nums []int, target int) []int {

	var resultIndices []int
	valueIndexMap := make(map[int]int)
	var numberToFind int

	// Populate the value index map with number to find and index values
	for index, value := range nums {

		indexOfTarget, ok := valueIndexMap[value]

		if ok {

			resultIndices = append(resultIndices, indexOfTarget, index)
			return resultIndices
		} else {

			numberToFind = target - value
			valueIndexMap[numberToFind] = index
		}
	}

	return nil
}

func backspaceCompare(s string, t string) bool {

	s = "y#fo##f"
	t = "y#f#o##f"
	s = "bxj##tw"
	t = "bxj###tw"
	s = "bbbextm"
	t = "bbb#extm"
	p1 := len(s) - 1
	p2 := len(t) - 1

	for p1 >= 0 || p2 >= 0 {

		if p1 >= 0 && s[p1] == '#' {
			backcount := 2

			for backcount > 0 {
				backcount--
				p1--
				if p1 >= 0 && s[p1] == '#' {
					backcount += 2
				}
			}
		}

		if p2 >= 0 && t[p2] == '#' {
			backcount := 2

			for backcount > 0 {
				backcount--
				p2--
				if p2 >= 0 && t[p2] == '#' {
					backcount += 2
				}
			}
		}

		if (p1 < 0 && p2 >= 0) || (p1 >= 0 && p2 < 0) {
			return false
		}
		if p1 < 0 && p2 < 0 {
			break
		}
		if s[p1] != t[p2] {
			return false
		} else {
			p1--
			p2--
		}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {

	s = "tmmzuxt"
	longestGlobalCount := 0
	longestLocalCount := 1

	indexHashMap := make(map[byte]int, 0)
	left := 0
	right := 0

	for right <= len(s)-1 {

		index, ok := indexHashMap[s[right]]

		if ok {
			if index >= left {
				left = index + 1
				longestLocalCount = 0
			} else {

				indexHashMap[s[right]] = left
				longestLocalCount = (right - left) + 1
				right++
				if longestLocalCount > longestGlobalCount {
					longestGlobalCount = longestLocalCount
				}
			}
		} else {

			indexHashMap[s[right]] = right

			longestLocalCount = (right - left) + 1
			right++

			if longestLocalCount > longestGlobalCount {
				longestGlobalCount = longestLocalCount
			}
		}
	}

	return longestGlobalCount
}

func sanitizeString(s string) string {

	stringSanitizer := regexp.MustCompile("[^A-Za-z0-9]+")

	s = stringSanitizer.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	return s
}
