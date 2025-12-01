package day1

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_simulate(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
			want: 3,
		},
		{
			input: aoc.Input(),
			want:  1048,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := simulate(tt.input)
			if got != tt.want {
				t.Errorf("simulate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simulatePartTwo(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
			want: 6,
		},
		{
			input: aoc.Input(),
			want:  6498,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := simulatePartTwo(tt.input)
			if got != tt.want {
				t.Errorf("simulate() = %v, want %v", got, tt.want)
			}
		})
	}
}
