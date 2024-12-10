package day15

import (
	"strings"

	"github.com/kungfukennyg/adventofgode/common"
)

const caloriesProp = "calories"

type ingredient struct {
	name       string
	properties map[string]int
}

type recipe struct {
	ingreds []*ingredient
	props   []string
}

func (r recipe) score(amounts []int, calories int) int {
	total := 1
	for _, prop := range r.props {
		if calories == -1 && prop == caloriesProp {
			continue
		}

		v := 0
		for i, ig := range r.ingreds {
			v += ig.properties[prop] * amounts[i]
		}
		if prop == caloriesProp {
			if v != calories {
				return 0
			}
		} else {
			total *= v
			if total <= 0 {
				return 0
			}
		}
	}
	return total
}

func (r recipe) highestScore(teaspoons, calories int) int {
	highestScore := -1
	n := len(r.ingreds)
	var recurse func(int, []int)
	recurse = func(remaining int, cur []int) {
		if len(cur) == n {
			if remaining == 0 {
				highestScore = max(highestScore, r.score(cur, calories))
			}
			return
		}

		for i := 1; i <= remaining; i++ {
			cur = append(cur, i)
			recurse(remaining-i, cur)
			cur = cur[:len(cur)-1]
		}
	}

	recurse(teaspoons, []int{})
	return highestScore
}

func bestRecipe(input string, teaspoons, calories int) int {
	rec := parse(input)
	highest := rec.highestScore(teaspoons, calories)
	return highest
}

func parse(input string) recipe {
	rec := recipe{
		ingreds: []*ingredient{},
		props:   []string{},
	}

	unique := map[string]struct{}{}
	for _, line := range common.Lines(input) {
		parts := strings.Split(line, ":")
		ig := ingredient{name: parts[0], properties: map[string]int{}}

		for _, part := range strings.Split(parts[1], ",") {
			parts = strings.Split(strings.TrimSpace(part), " ")
			prop := parts[0]
			v := common.MustAtoi(parts[1])

			ig.properties[prop] = v
			unique[prop] = struct{}{}
		}

		rec.ingreds = append(rec.ingreds, &ig)
	}

	for k := range unique {
		rec.props = append(rec.props, k)
	}
	return rec
}
