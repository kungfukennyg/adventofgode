package aoc

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

const inputFilePath = "input.txt"

func Input() string {
	return ReadFile(inputFilePath)
}

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func ContainsOnly(in string, set string) bool {
	for _, s := range in {
		if !strings.ContainsRune(set, s) {
			return false
		}
	}
	return true
}

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func MustAtoi64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func MustUint8(s string) uint8 {
	n, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		panic(err)
	}
	return uint8(n)
}

func MustBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return b
}

func MustText(buf *bufio.Scanner) string {
	if !buf.Scan() {
		if buf.Err() != nil {
			panic(buf.Err())
		}
	}

	return buf.Text()
}

func IndicesOf(str string, find string) []int {
	ids := []int{}
	for i := 0; i < len(str); i++ {
		if len(find) > 1 && i+len(find) >= len(str) {
			continue
		}
		s := str[i : i+len(find)]
		if s == find {
			ids = append(ids, i)
		}
	}

	return ids
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func Ints(s string, delim string) []int {
	parts := strings.Split(s, delim)
	out := []int{}
	for _, s := range parts {
		if len(strings.TrimSpace(s)) == 0 {
			continue
		}

		out = append(out, MustAtoi(s))
	}
	return out
}

func Int64s(s string, delim string) []int64 {
	parts := strings.Split(s, delim)
	out := []int64{}
	for _, s := range parts {
		if len(strings.TrimSpace(s)) == 0 {
			continue
		}

		out = append(out, MustAtoi64(s))
	}
	return out
}

func Uint8s(s string, delim string) []uint8 {
	parts := strings.Split(s, delim)
	out := []uint8{}
	for _, s := range parts {
		if len(strings.TrimSpace(s)) == 0 {
			continue
		}

		out = append(out, MustUint8(s))
	}
	return out
}

func JoinInts(o []int, delim string) string {
	var sb strings.Builder
	for _, i := range o {
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString(delim)
	}

	return strings.TrimSuffix(sb.String(), ",")
}

func JoinUint8s(o []uint8, delim string) string {
	var sb strings.Builder
	for _, i := range o {
		sb.WriteString(strconv.FormatUint(uint64(i), 8))
		sb.WriteString(delim)
	}

	return strings.TrimSuffix(sb.String(), ",")
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
