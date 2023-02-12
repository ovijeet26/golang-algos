package string

import (
	"fmt"
	"strings"
)

//* Given a string of words with lots of spaces between the words , remove all the unnecessary spaces like
// input:  I   live   on     earth
// output: I live on earth
func stringSpacesRemover(input string) string {

	stringArr := strings.Split(input, "")
	finalArr := make([]string, 0)

	for i := 0; i < len(stringArr); i++ {

		if stringArr[i] != " " {
			finalArr = append(finalArr, stringArr[i])
			continue
		}

		if stringArr[i] == " " {

			finalArr = append(finalArr, stringArr[i])

			for stringArr[i+1] == " " {

				i++
			}
		}
	}

	return strings.Join(finalArr, "")
}

func SpaceRemover() {

	fmt.Printf("Original -> %s\n", "Hello from the   other   side    .")

	input := stringSpacesRemover("Hello from the   other   side    .")
	fmt.Printf("Reduced -> %s", input)
}
