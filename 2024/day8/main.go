package day8

import (
	"strings"
)

type pos struct {
	x, y int
}

func (p pos) sub(b pos) pos {
	return pos{p.x - b.x, p.y - b.y}
}

func (p pos) add(b pos) pos {
	return pos{p.x + b.x, p.y + b.y}
}

type grid struct {
	antennas    map[string][]pos
	frequencies [][]string
	antinodes   [][]bool
}

func (g grid) String() string {
	var sb strings.Builder
	for y, row := range g.frequencies {
		for x, freq := range row {
			ok := g.antinodes[y][x]
			if freq != "" && ok {
				sb.WriteString("+")
			} else if freq != "" {
				sb.WriteString(freq)
			} else if ok {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (g grid) boundsCheck(p pos) bool {
	x, y := p.x, p.y
	return y >= 0 && y < len(g.frequencies) && x >= 0 && x < len(g.frequencies[y])
}

func (g grid) countAntinodes() int {
	antinodes := 0
	for y, row := range g.antinodes {
		for x, v := range row {
			if v {
				antinodes++
			} else if s := g.frequencies[y][x]; s != "" && len(g.antennas[s]) > 1 {
				antinodes++
			}
		}
	}
	return antinodes
}

func findAntinodesAnyDistance(input string) int {
	g := parse(input)
	for _, positions := range g.antennas {
		for i, a := range positions {
			for j, b := range positions {
				if i == j {
					continue
				}

				diff := a.sub(b)
				above, below := a.add(diff), b.sub(diff)
				for {
					if !g.boundsCheck(above) {
						break
					}

					g.antinodes[above.y][above.x] = true
					above = above.add(diff)
				}
				for {
					if !g.boundsCheck(below) {
						break
					}
					g.antinodes[below.y][below.x] = true
					below = below.sub(diff)
				}
			}
		}
	}
	return g.countAntinodes()
}

func findAntinodes(input string) int {
	g := parse(input)
	for _, positions := range g.antennas {
		for i, a := range positions {
			for j, b := range positions {
				if i == j {
					continue
				}
				diff := a.sub(b)
				above, below := a.add(diff), b.sub(diff)
				if g.boundsCheck(above) {
					g.antinodes[above.y][above.x] = true
				} else if g.boundsCheck(below) {
					g.antinodes[below.y][below.x] = true
				}
			}
		}
	}

	return g.countAntinodes()
}

func parse(input string) grid {
	g := grid{
		antennas:    map[string][]pos{},
		frequencies: [][]string{},
		antinodes:   [][]bool{},
	}
	for y, line := range strings.Split(input, "\n") {
		g.frequencies = append(g.frequencies, make([]string, len(line)))
		g.antinodes = append(g.antinodes, make([]bool, len(line)))
		for x, s := range line {
			if s == '.' {
				continue
			}
			freq := string(s)
			g.frequencies[y][x] = freq
			positions, ok := g.antennas[freq]
			if !ok {
				positions = []pos{}
			}
			positions = append(positions, pos{x, y})
			g.antennas[freq] = positions
		}
	}
	return g
}
