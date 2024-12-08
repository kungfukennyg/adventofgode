package day6

import (
	"fmt"
	"strings"
)

type dir int

const (
	dirNone dir = iota
	dirUp
	dirRight
	dirDown
	dirLeft
)

func (d dir) step(x, y int) (int, int) {
	switch d {
	case dirDown:
		return x, y + 1
	case dirLeft:
		return x - 1, y
	case dirRight:
		return x + 1, y
	case dirUp:
		return x, y - 1
	default:
		panic(fmt.Sprintf("unexpected day6.dir: %#v", d))
	}
}

func (d dir) turn() dir {
	switch d {
	case dirUp:
		return dirRight
	case dirRight:
		return dirDown
	case dirDown:
		return dirLeft
	case dirLeft:
		return dirUp
	default:
		panic(fmt.Sprintf("unexpected day6.dir: %#v", d))
	}
}

func (d dir) String() string {
	switch d {
	case dirDown:
		return "|"
	case dirLeft:
		return "-"
	case dirRight:
		return "-"
	case dirUp:
		return "|"
	default:
		return ""
	}
}

type grid [][]bool

type board struct {
	obstacles      grid
	visited        [][]dir
	startX, startY int
	guardX, guardY int
	guardDir       dir
}

func (b *board) boundsCheck(x, y int) bool {
	return y >= 0 && y < len(b.obstacles) && x >= 0 && x < len(b.obstacles[y])
}

func (b *board) step() bool {
	dx, dy := b.guardDir.step(b.guardX, b.guardY)
	if !b.boundsCheck(dx, dy) {
		return false
	}

	if blocked := b.obstacles[dy][dx]; blocked {
		b.guardDir = b.guardDir.turn()
		return true
	}

	b.guardX = dx
	b.guardY = dy
	b.visited[dy][dx] = b.guardDir

	return true
}

func (b board) String() string {
	var sb strings.Builder
	for y, row := range b.obstacles {
		for x, v := range row {
			if x == b.startX && y == b.startY {
				sb.WriteRune('o')
			} else if v {
				sb.WriteRune('#')
			} else if vv := b.visited[y][x]; vv != dirNone {
				sb.WriteString(vv.String())
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (b *board) clone() *board {
	c := &board{
		obstacles: make(grid, len(b.obstacles)),
		visited:   make([][]dir, len(b.visited)),
		startX:    b.startX,
		startY:    b.startY,
		guardX:    b.guardX,
		guardY:    b.guardY,
		guardDir:  b.guardDir,
	}

	for y, row := range b.obstacles {
		c.obstacles[y] = make([]bool, len(row))
		copy(c.obstacles[y], b.obstacles[y])
	}
	for y, row := range b.visited {
		c.visited[y] = make([]dir, len(row))
		copy(c.visited[y], b.visited[y])
	}

	return c
}

func parse(input string) *board {
	b := board{
		obstacles: [][]bool{},
		visited:   [][]dir{},
		guardDir:  dirUp,
	}
	for y, line := range strings.Split(input, "\n") {
		objs := make([]bool, len(line))
		for x, s := range line {
			switch s {
			case '.':
				continue
			case '#':
				objs[x] = true
			case '^':
				b.guardX, b.guardY = x, y
				b.startX, b.startY = x, y
			}
		}
		b.obstacles = append(b.obstacles, objs)
		b.visited = append(b.visited, make([]dir, len(objs)))
	}
	b.visited[b.startY][b.startX] = b.guardDir
	return &b
}

func (b *board) visitedByGuard() int {
	for b.step() {

	}

	visited := 0
	for _, row := range b.visited {
		for _, v := range row {
			if v != dirNone {
				visited++
			}
		}
	}
	return visited
}

func (b *board) infiniteObstacles() int {
	infiniteObstacles := 0
	for y, row := range b.obstacles {
		for x := range row {
			c := b.clone()
			if c.placeObstacle(x, y) {
				if !c.findsExit() {
					infiniteObstacles++
				}
			}
		}
	}

	return infiniteObstacles
}

func (b *board) findsExit() bool {
	loopPoint := len(b.obstacles) * len(b.obstacles[0])
	steps := 0
	for b.step() {
		steps++
		if steps > loopPoint {
			return false
		}
	}

	return true
}

func (b *board) placeObstacle(x, y int) bool {
	if x == b.startX && y == b.startY {
		return false
	}

	if !b.boundsCheck(x, y) {
		return false
	}

	if v := b.obstacles[y][x]; v {
		return false
	}

	b.obstacles[y][x] = true
	return true
}
