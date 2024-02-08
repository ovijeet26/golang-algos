package string

func groupAnagrams(strs []string) [][]string {

	hashMap := make(map[string][]string, 0)
	result := make([][]string, 0)
	a := 'a'
	asciiOfA := int(a)

	for _, str := range strs {

		asciiArr := make([]int, 26)
		for i := 0; i < len(str); i++ {
			// assuming ascii of a is 0
			charIndex := int(str[i]) - asciiOfA
			asciiArr[charIndex] += 1
		}
		key := ""
		for index, count := range asciiArr {

			if count > 0 {
				key += string(index) + string(count)
			}
		}
		_, ok := hashMap[key]
		if ok {
			hashMap[key] = append(hashMap[key], str)
		} else {
			hashMap[key] = []string{str}
		}
	}

	for _, arr := range hashMap {
		result = append(result, arr)
	}

	return result
}
