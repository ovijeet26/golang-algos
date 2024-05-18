package string

import (
	"fmt"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {

	hashMap := make(map[string][]string)
	res := make([][]string, 0)

	for _, str := range strs {

		code := generateCode(str)

		hashMap[code] = append(hashMap[code], str)
	}

	for _, val := range hashMap {
		res = append(res, val)
	}
	return res
}

func generateCode(s string) string {

	arr := make([]int, 150)

	for _, v := range s {
		intVal := int(v)
		arr[intVal]++
	}

	code := []string{}
	for i, c := range arr {

		if c > 0 {
			val := fmt.Sprintf("%v%v|", i, c)
			code = append(code, val)
		}
	}

	return strings.Join(code, "")
}
