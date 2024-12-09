package day10

import (
	"bufio"
	"strconv"
	"strings"
)

func lookAndSay(input string, n int) string {
	seq := input
	for range n {
		seq = generate(seq)
	}
	return seq
}

func generate(in string) string {
	buf := bufio.NewScanner(strings.NewReader(in))
	buf.Split(bufio.ScanRunes)
	var sb strings.Builder
	seq := ""
	for buf.Scan() {
		// read until we hit a repeat
		tok := buf.Text()
		if len(seq) > 0 && string(seq[0]) != tok {
			// commit last sequence
			count := len(seq)
			sb.WriteString(strconv.Itoa(count))
			sb.WriteString(string(seq[0]))
			seq = ""
		}
		seq += tok
	}

	if len(seq) > 0 {
		count := len(seq)
		sb.WriteString(strconv.Itoa(count))
		sb.WriteString(string(seq[0]))
	}

	return sb.String()
}
