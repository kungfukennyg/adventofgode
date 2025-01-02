package day19

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_distinctMolecules(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{`H => HO
H => OH
O => HH

HOH`},
			want: 4,
		},
		{
			args: args{`H => HO
H => OH
O => HH

HOHOHO`},
			want: 7,
		},
		{
			args: args{aoc.Input()},
			want: 0,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := distinctMolecules(tt.args.input); got != tt.want {
				t.Errorf("distinctMolecules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSteps(t *testing.T) {
	type args struct {
		input    string
		startMol string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: `e => H
e => O
H => HO
H => OH
O => HH

HOH`,
				startMol: "e",
			},
			want: 3,
		},
		{
			args: args{
				input: `e => H
e => O
H => HO
H => OH
O => HH

HOHOHO`,
				startMol: "e",
			},
			want: 6,
		},
		{
			args: args{
				input:    aoc.Input(),
				startMol: "e",
			},
			want: 0,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := minSteps(tt.args.input, tt.args.startMol); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
