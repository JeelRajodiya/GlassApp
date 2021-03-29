package Engine

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var numbers []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."}

type Parsed struct {
	to_prefix   string
	to          string
	from        string
	from_prefix string
	value       float64
}

func detect_prefix(unit_name string) (string, string) {
	var new_name string

	_, isunit := values[unit_name]

	if isunit {
		return unit_name, "None"
	}

	for k := range prefixes {

		if strings.HasPrefix(unit_name, k) {

			new_name = strings.Replace(unit_name, k, "", 1)
			return new_name, k
		}
	}
	return unit_name, "None"

}
func parseConvertable(input []string) Parsed {

	var parsed Parsed

	var temp string
	var wasNumber bool = false

	// parse before *to

	for _, i := range input[0] {
		i2 := string(i)
		if check_membership(i2, numbers) {
			wasNumber = true
			temp = temp + i2

		} else if wasNumber {
			wasNumber = false
			value, _ := strconv.ParseFloat(temp, 64) //temp will be unit value
			parsed.value = value
			temp = ""
			temp = temp + i2

		} else if !wasNumber {
			temp = temp + i2
		}
	}
	parsed.from = temp //temp will be the base unit name
	parsed.to = input[2]

	parsed.from, parsed.from_prefix = detect_prefix(parsed.from)
	parsed.to, parsed.to_prefix = detect_prefix(parsed.to)
	return parsed
}

var prefixes = map[string]float64{
	"c":    math.Pow10(-2),
	"mu":   math.Pow10(-6),
	"n":    math.Pow10(-9),
	"m":    math.Pow10(-3),
	"p":    math.Pow10(-12),
	"f":    math.Pow10(-15),
	"d":    math.Pow10(-1),
	"None": 1,
	"da":   math.Pow10(1),
	"h":    math.Pow10(2),
	"K":    math.Pow10(3),
	"M":    math.Pow10(6),
	"G":    math.Pow10(9),
	"T":    math.Pow10(12),
	"P":    math.Pow10(15),
}

var values = map[string]float64{
	"m":    1,
	"J":    1,
	"atm":  1,
	"cal":  4.18,
	"In":   0.0254,
	"Hg":   0.013,
	"torr": 0.00132,
	"bar":  0.987,
	"pa":   9.869 * math.Pow10(-6),
}

func convert(input []string) []string {
	var parsed Parsed = parseConvertable(input)
	var answer = (parsed.value * (values[parsed.from] * prefixes[parsed.from_prefix])) / (values[parsed.to] * prefixes[parsed.to_prefix])
	// 5cm = x mm
	// ( (5 * cm) / mm ) = x
	fmt.Println(parsed)
	return []string{strconv.FormatFloat(answer, 'G', -100, 64)}
}
