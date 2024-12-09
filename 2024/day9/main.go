package day9

import (
	"strconv"
	"strings"

	"github.com/kungfukennyg/adventofgode/common"
)

type disk []int

const empty = -1

func (d disk) String() string {
	var sb strings.Builder
	for _, n := range d {
		if n == empty {
			sb.WriteRune('.')
		} else {
			sb.WriteString(strconv.Itoa(n))
		}
	}
	return sb.String()
}

func (d disk) freeBlock(size, wantIdx int) int {
	foundBlock := 0
	idx := -1
	for i, n := range d {
		if i >= wantIdx {
			break
		}
		if n != empty {
			if foundBlock != 0 {
				foundBlock = 0
				idx = -1
			}
		} else {
			if foundBlock == 0 {
				idx = i
			}
			foundBlock++
			if foundBlock == size {
				return idx
			}
		}
	}

	return -1
}

func (d disk) firstEmpty() int {
	for i, n := range d {
		if n == empty {
			return i
		}
	}
	return -1
}

func (d disk) compact() bool {
	for i := len(d) - 1; i >= 0; i-- {
		n := d[i]
		if n == empty {
			continue
		}
		open := d.firstEmpty()
		if open == -1 {
			return false
		}
		if open > i {
			return false
		}

		d[open] = n
		d[i] = empty
	}
	return true
}

func (d disk) compactWholeFiles(highestId int) {
	file := -1
	size := 0
	idx := -1
	for i := len(d) - 1; i >= 0; i-- {
		if highestId < 0 {
			break
		}
		n := d[i]
		if idx == -1 {
			if n != empty && n == highestId {
				file = n
				idx = i
				size++
			}
		} else {
			if n == file {
				size++
			} else {
				open := d.freeBlock(size, idx)
				if open != -1 && open < i {
					for j := range size {
						d[j+open] = file
						d[i+1+j] = empty
					}
				}
				highestId--
				if n != empty && n == highestId {
					idx = i
					file = n
					size = 1
				} else {
					idx = -1
					file = -1
					size = 0
				}
			}
		}
	}
}

func (d disk) checksum() int {
	sum := 0
	for i, n := range d {
		if n == empty {
			continue
		}

		sum += i * n
	}
	return sum
}

func processDisk(input string, wholeFile bool) int {
	d, highestId := parse(input)
	if wholeFile {
		d.compactWholeFiles(highestId)
	} else {
		for d.compact() {

		}
	}

	s := d.String()
	_ = s
	sum := d.checksum()
	return sum
}

func parse(input string) (disk, int) {
	file := true
	d := disk{}
	id := 0
	for _, r := range input {
		n := common.MustAtoi(string(r))
		if file {
			for range n {
				d = append(d, id)
			}
			id++
		} else {
			for range n {
				d = append(d, empty)
			}
		}
		file = !file
	}
	return d, id - 1
}
