package Engine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var wordsHolder []string = make([]string, 53751)

func init() {
	file, err := os.Open("Engine/common.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		if line == "" {
			continue
		}
		wordsHolder = append(wordsHolder, line)
	}
}

func check_membership(elem string, arr []string) bool {
	for _, i := range arr {
		if i == elem {
			return true
		}
	}
	return false
}

// commands -|
//			 |-- help--|
//                     |----tools ---------|-- converter
//										   |-- maths
//           |-- ** to **
//           |-- simple maths
//           |--

func gen_result(input string) []string {
	var commands []string = []string{
		"help",
		"tools",
		"converter",
		"maths",
	}

	type card []string

	var helpCards = []string{
		"these are the valid commands",
		"\"tools\" - to see available list of tools",
		"Tool_Name - where Tool_Name is a name of tool. it will give you a details about specific tool",
	}

	var toolsCards = []string{
		"these tools are supported",
		"converter - will convert from any scientific unit from other unit",
		"math - will allow you to do basic maths",
	}

	var converterCards = []string{
		"converter will convert from any scientific unit from other unit",
		"to use converter use the following syntax",
		" \"6mm to cm\" will convert 6mm to cm ",
	}

	var mathsCards = []string{
		"maths will allow you to do simple mathematics",
		" \"2+2\" will give 4",
	}

	var commandMap = map[string]card{
		"help":      helpCards,
		"tools":     toolsCards,
		"converter": converterCards,
		"maths":     mathsCards,
	}

	if check_membership(input, commands) {
		return commandMap[input]
	}

	var processedInput []string = strings.Split(input, " ")
	if len(processedInput) > 1 {
		if processedInput[1] == "to" {
			var converted []string = convert(processedInput)
			return converted
		}
	} else if ismath(input) {
		return solveMath(input)

	}
	fmt.Println(ismath(input))
	return sortWords(input)

}

func Perform_search(input string) []string {
	var result = make([]string, 0)

	result = gen_result(input)
	if len(result) > 100 {
		result = result[:100]
	} else if len(result) == 0 {
		result = []string{" Not Found "}
	}

	return result

}
