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

func MustText(buf *bufio.Scanner) string {
	if !buf.Scan() {
		if buf.Err() != nil {
			panic(buf.Err())
		}
	}

	return buf.Text()
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

func Permutate[T any](arr []T) [][]T {
	var helper func([]T, int)
	res := [][]T{}

	helper = func(arr []T, n int) {
		if n == 1 {
			tmp := make([]T, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
