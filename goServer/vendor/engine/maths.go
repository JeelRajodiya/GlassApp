package Engine

import (
	"fmt"
	"strconv"
	"strings"
)

func ismath(input string) bool {
	// a = +
	// s = -
	// m = *
	// d = /
	input = strings.ReplaceAll(input, " ", "")
	if strings.ContainsAny(input, "asmd+/-*") && strings.ContainsAny(input, "0123456789") && !strings.ContainsAny(input, "qwertyuioplkjhgfzxcvbn") {
		return true
	}
	return false
}

func parseMath(input string) []string {
	var parsed = make([]string, 5)

	var wasNumber bool
	var temp string
	var oprations = []string{"a", "d", "s", "m",
		"+", "/", "-", "*"}
	for _, iRune := range input {

		var i = string(iRune)
		if check_membership(i, numbers) {
			wasNumber = true
			temp = temp + i

		} else if wasNumber && check_membership(i, oprations) {
			parsed = append(parsed, temp)
			temp = temp + i
			wasNumber = false
		}
	}

	return parsed
}

func solveMath(input string) []string {
	fmt.Println(parseMath(input))
	if strings.Contains(input, "a") {

		if strings.Count(input, "a") > 1 {
			return []string{" please use a only one time to add numbers"}
		}
		var nums = strings.Split(input, "a")
		var a, _ = strconv.ParseFloat(nums[0], 64)
		var b, _ = strconv.ParseFloat(nums[1], 64)

		var answer = a + b

		return []string{strconv.FormatFloat(answer, 'G', -100, 64)}
	}
	return []string{"S"}
}
