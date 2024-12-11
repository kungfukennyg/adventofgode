package day4

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_countPattern(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX",
			},
			want: 18,
		},
		{
			args: args{
				input: aoc.Input(),
			},
			want: 2336,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := new(tt.args.input)
			if got := g.countPattern("XMAS"); got != tt.want {
				t.Errorf("countPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crossPattern(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX",
			},
			want: 9,
		},
		{
			args: args{
				input: aoc.Input(),
			},
			want: 1831,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := new(tt.args.input)
			if got := g.crossPattern("MAS", 'A'); got != tt.want {
				t.Errorf("crossPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
