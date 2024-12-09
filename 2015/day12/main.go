package day12

import (
	"encoding/json"
	"regexp"

	"github.com/kungfukennyg/adventofgode/common"
)

var numPattern = regexp.MustCompile("[0-9]+|-[0-9]+")

func sumNumbers(input string) int {
	sum := 0
	for _, n := range numPattern.FindAllString(input, -1) {
		sum += common.MustAtoi(n)
	}
	return sum
}

func sumSkipRed(input string) int {
	var d interface{}
	json.Unmarshal([]byte(input), &d)
	return parse(d)
}

func parse(d interface{}) int {
	o := 0
OUTER:
	switch typ := d.(type) {
	case []interface{}:
		for _, v := range typ {
			o += parse(v)
		}
	case float64:
		o += int(typ)
	case map[string]interface{}:
		for _, v := range typ {
			if v == "red" {
				break OUTER
			}
		}

		for _, v := range typ {
			o += parse(v)
		}
	}
	return o
}
