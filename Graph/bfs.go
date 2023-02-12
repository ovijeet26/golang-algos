package graph

import "fmt"

func sliceIt(s []int) {

	s[1] = -99
	s = append(s, 5)
	s[0] = -99

}

func SliceCaller() {

	s := []int{0, 1, 2, 3, 4}

	fmt.Println(s)

	sliceIt(s)

	fmt.Println(s)
}
