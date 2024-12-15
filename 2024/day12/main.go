package day12

import (
	"github.com/kungfukennyg/adventofgode/aoc"
)

type region struct {
	r            rune
	points       []aoc.Pos
	uniquePoints aoc.Set[aoc.Pos]
}

func (r *region) add(p aoc.Pos) {
	if !r.uniquePoints.Add(p) {
		return
	}

	r.points = append(r.points, p)
}

func (r *region) contains(p aoc.Pos) bool {
	_, ok := r.uniquePoints[p]
	return ok
}

func (r *region) area() int {
	return len(r.points)
}

var cardinal = []aoc.Dir{aoc.DirUp, aoc.DirRight, aoc.DirDown, aoc.DirLeft}
var diagonals = []aoc.Dir{aoc.DirUpLeft, aoc.DirUpRight, aoc.DirDownLeft, aoc.DirDownRight}

func (r *region) sides(grid aoc.Grid[rune]) int {
	isSame := func(p aoc.Pos, r rune) bool {
		return grid.BoundsCheck(p) && grid.GetByPos(p) == r
	}

	corners := 0
	for _, p := range r.points {
		for _, d := range diagonals {
			corner := p.Step(d)
			steps := d.Decompose()
			a, b := p.Step(steps[0]), p.Step(steps[1])
			if isSame(a, r.r) && isSame(b, r.r) && !isSame(corner, r.r) {
				corners++
			}
			if !isSame(a, r.r) && !isSame(b, r.r) {
				corners++
			}
		}
	}
	return corners
}

func (r *region) perimeter(grid aoc.Grid[rune]) int {
	total := 0
	for _, pos := range r.points {
		sides := 4
		for p, _ := range grid.Neighbors(pos, cardinal) {
			if r.contains(p) {
				sides--
			}
		}
		total += sides
	}
	return total
}

func fencePrice(input string, discount bool) int {
	grid := parse(input)
	regions := getRegions(grid)

	price := 0
	for _, region := range regions {
		if !discount {
			a, p := region.area(), region.perimeter(grid)
			price += a * p
		} else {
			a, s := region.area(), region.sides(grid)
			price += a * s
		}
	}
	return price
}

func getRegions(grid aoc.Grid[rune]) []*region {
	regions := []*region{}
	added := aoc.Set[aoc.Pos]{}
	for pos, r := range grid.Points() {
		if added.Contains(pos) {
			continue
		}

		stack := aoc.Stack[aoc.Pos]{}
		stack.Push(pos)

		region := &region{r: r, points: []aoc.Pos{}, uniquePoints: aoc.Set[aoc.Pos]{}}
		regions = append(regions, region)
		visited := aoc.Set[aoc.Pos]{}

		for !stack.IsEmpty() {
			p, _ := stack.Pop()
			if !visited.Add(p) {
				continue
			}

			v := grid.GetByPos(p)
			region.add(p)
			added.Add(p)

			for p, r := range grid.Neighbors(p, cardinal) {
				if added.Contains(p) {
					continue
				}

				if v != r {
					continue
				}

				stack.Push(p)
			}
		}
	}
	return regions
}

func parse(input string) aoc.Grid[rune] {
	g := aoc.Grid[rune]{}
	for y, row := range aoc.Lines(input) {
		g.AddRow(make([]rune, len(row)))
		for x, p := range row {
			g.Set(x, y, p)
		}
	}
	return g
}
