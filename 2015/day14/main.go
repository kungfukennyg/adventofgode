package day14

import (
	"fmt"
	"math"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type reindeer struct {
	name     string
	speed    int
	airTime  int
	restTime int

	distance  int
	stateTime int
	resting   bool
	points    int
}

func (r *reindeer) step() {
	if r.resting {
		r.stateTime++
		if r.stateTime >= r.restTime {
			r.resting = false
			r.stateTime = 0
		}
	} else {
		r.distance += r.speed
		r.stateTime++
		if r.stateTime >= r.airTime {
			r.resting = true
			r.stateTime = 0
		}
	}
}

func scoreLeading(deer []*reindeer) {
	highest := 0
	for _, r := range deer {
		if r.distance > highest {
			highest = r.distance
		}
	}
	for _, r := range deer {
		if r.distance == highest {
			r.points++
		}
	}
}

func mostPoints(input string, seconds int) int {
	deer := parse(input)

	last := -1
	for i := range seconds {
		for _, r := range deer {
			r.step()
		}
		scoreLeading(deer)
		_ = i
		last = i
	}

	fmt.Println(last)
	highest := math.MinInt64
	for _, d := range deer {
		highest = max(d.points, highest)
	}
	return highest
}

func furthestDistance(input string, seconds int) int {
	deer := parse(input)

	distances := []int{}
	for _, d := range deer {
		dist := d.airTime * d.speed
		cycleLen := d.airTime + d.restTime
		cycles := seconds / cycleLen
		dist *= cycles
		remaining := seconds % cycleLen
		if remaining > d.airTime {
			dist += d.speed * d.airTime
		} else {
			dist += d.speed * remaining
		}
		distances = append(distances, dist)
	}

	highest := 0
	for _, d := range distances {
		highest = max(highest, d)
	}
	return highest
}

func parse(input string) []*reindeer {
	deer := []*reindeer{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		r := reindeer{
			name:     parts[0],
			speed:    aoc.MustAtoi(parts[3]),
			airTime:  aoc.MustAtoi(parts[6]),
			restTime: aoc.MustAtoi(parts[13]),
		}
		deer = append(deer, &r)
	}
	return deer
}
