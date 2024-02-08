package string

func isAnagram(s string, t string) bool {

	size := len(s)
	if len(t) != size {
		return false
	}
	hashMap := make(map[rune]int, 0)

	for _, char := range s {

		hashMap[char]++
	}

	for _, char := range t {

		_, isPresent := hashMap[char]
		hashMap[char]--
		if !isPresent {
			return false
		}
	}

	for _, value := range hashMap {

		if value != 0 {
			return false
		}
	}

	return true
}
