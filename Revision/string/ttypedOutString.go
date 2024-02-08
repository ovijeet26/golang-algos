package string

func backspaceCompare(s string, t string) bool {

	rightS := len(s) - 1
	rightT := len(t) - 1
	for rightS >= 0 && rightT >= 0 {

		sBackCount := 0
		tBackCount := 0
		for rightS >= 0 && s[rightS] == '#' {
			sBackCount++
			rightS--
		}

		for rightT >= 0 && t[rightT] == '#' {
			tBackCount++
			rightT--
		}

		for (rightS >= 0 && sBackCount > 0) || (rightS >= 0 && s[rightS] == '#') {
			if s[rightS] == '#' {
				sBackCount += 2
			}
			rightS--
			sBackCount--
		}

		for (rightT >= 0 && tBackCount > 0) || (rightT >= 0 && t[rightT] == '#') {
			if t[rightT] == '#' {
				tBackCount += 2
			}
			rightT--
			tBackCount--
		}

		if rightT < 0 && rightS >= 0 || rightS < 0 && rightT >= 0 {
			break
		}

		if rightT < 0 && rightS < 0 {
			break
		}

		if s[rightS] != t[rightT] {
			return false
		}
		rightS--
		rightT--
	}

	if rightS >= 0 {
		if isBlankString(s[:rightS+1]) {
			return true
		}
		return false
	}

	if rightT >= 0 {
		if isBlankString(t[:rightT+1]) {
			return true
		}
		return false
	}

	return true
}

func isBlankString(s string) bool {

	size := len(s)
	if size == 1 {
		if s[0] == '#' {
			return true
		} else {
			return false
		}
	}
	right := size - 1
	backSpace := 0
	for right > 0 {

		if right > 0 && s[right] == '#' {
			backSpace++
			right--
		}

		for backSpace > 0 && right > 0 {
			if s[right] == '#' {
				backSpace += 2
			}
			backSpace--
			right--
		}

		if right > 0 && s[right] != '#' {
			return false
		}
	}
	return true
}
