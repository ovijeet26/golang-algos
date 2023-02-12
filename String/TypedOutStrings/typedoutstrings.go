package typedoutstrings

//Given two strings s and t, return true if they are equal when both are typed into empty text editors.
//'#' means a backspace character.
//Note that after backspacing an empty text, the text will continue empty.

func backspaceCompare(s string, t string) bool {

	p1 := len(s) - 1
	p2 := len(t) - 1

	for p1 >= 0 && p2 >= 0 {

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
		if s[p1] != t[p2] {
			return false
		} else {
			p1--
			p2--
		}
	}
	return true
}

func backspaceCompare1(s string, t string) bool {

	p1 := len(s) - 1
	p2 := len(t) - 1

	for p1 >= 0 || p2 >= 0 {

		if (p1 >= 0 && string(s[p1]) == "#") || string(t[p2]) == "#" {

			// Check for backspace case for the first string. Also consider back-to-back deletions e.g. a###b
			if string(s[p1]) == "#" {
				backcount := 2
				for backcount > 0 {
					p1--
					backcount--
					if p1 >= 0 && string(s[p1]) == "#" {

						backcount += 2
					}
				}
			}

			// Check for backspace case for the second string. Also consider back-to-back deletions e.g. a###b
			if string(t[p2]) == "#" {
				backcount := 2
				for backcount > 0 {
					p2--
					backcount--
					if p2 >= 0 && string(t[p2]) == "#" {

						backcount += 2
					}
				}
			}
		} else {
			if s[p1] != t[p2] {
				return false
			} else {
				p1--
				p2--
			}
		}
	}
	if p1 == p2 || (p1 < 0 && p2 < 0) {
		return true
	} else {
		return false
	}
}
