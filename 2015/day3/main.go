package day3

type dir string

const (
	dirUp    dir = "^"
	dirRight dir = ">"
	dirDown  dir = "v"
	dirLeft  dir = "<"
)

func dirFromStr(s rune) dir {
	switch dir(s) {
	case dirDown:
	case dirLeft:
	case dirRight:
	case dirUp:
	default:
		panic("unexpected day3.dir")
	}
	return dir(s)
}

func (d dir) move(pos *coord) {
	switch d {
	case dirDown:
		pos.y--
	case dirLeft:
		pos.x--
	case dirRight:
		pos.x++
	case dirUp:
		pos.y++
	default:
		panic("unexpected day3.dir")
	}
}

func housesThatGetPresent(input string) int {
	total := 1
	visited := map[coord]struct{}{}
	pos := &coord{x: 0, y: 0}
	visit(pos, visited)
	for _, s := range input {
		dir := dirFromStr(s)
		dir.move(pos)
		if visit(pos, visited) {
			total++
		}
	}

	return total
}

type coord struct {
	x, y int
}

func housesWithRobot(input string) int {
	total := 1
	visited := map[coord]struct{}{}
	santa, robot := &coord{x: 0, y: 0}, &coord{x: 0, y: 0}
	visit(santa, visited)
	visit(robot, visited)
	flip := true
	for _, s := range input {
		var pos *coord
		if flip {
			pos = santa
		} else {
			pos = robot
		}

		dir := dirFromStr(s)
		dir.move(pos)
		if visit(pos, visited) {
			total++
		}
		flip = !flip
	}

	return total
}

func visit(pos *coord, visited map[coord]struct{}) bool {
	if _, ok := visited[*pos]; !ok {
		visited[*pos] = struct{}{}
		return true
	}

	return false
}
