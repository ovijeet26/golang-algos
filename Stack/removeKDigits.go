package stack

import (
	"strconv"
	"strings"
)
func removeKdigits(num string, k int) string {
	stack := Stack{}

	index := 0

    if k>=len(num){
        return "0"
    }
	for index < len(num) {

		currentNum, _ := strconv.Atoi(string(num[index]))
		if stack.Length() == 0{
			stack.Push(string(num[index]))
			index++
			continue
		}

		if currentNum > toNum(stack.Peek()) {
			stack.Push(string(num[index]))
			index++
			continue
		}

		for k > 0 && stack.Length() > 0 && currentNum < toNum(stack.Peek()) {
			stack.Pop()
			k--
		}
		stack.Push(string(num[index]))
		index++
	}

    finalString :=strings.Join(stack.ArrayValues(), "")

    if len(finalString)>k && k>0{
            finalString = finalString[:len(finalString)-k]

    }
    leadingZeroCounter:=0
    i:=0
    for i=0; i<len(finalString);i++{
        if string(finalString[i])!="0"{
            break
        }
        leadingZeroCounter++
    }
    if i==len(finalString){
        return "0"
    }

    finalString = finalString[leadingZeroCounter:]

	return finalString
}


func toNum(s string) int {

	num, _ := strconv.Atoi(s)
	return num
}

//Input: num = "1432219", k = 3  1 432 219
//Output: "1219"
