package day14

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_furthestDistance(t *testing.T) {
	type args struct {
		input   string
		seconds int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input:   "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.\nDancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
				seconds: 1000,
			},
			want: 1120,
		},
		{
			args: args{input: aoc.Input(), seconds: 2503},
			want: 2660,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := furthestDistance(tt.args.input, tt.args.seconds); got != tt.want {
				t.Errorf("furthestDistance) = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mostPoints(t *testing.T) {
	type args struct {
		input   string
		seconds int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input:   "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.\nDancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
				seconds: 1000,
			},
			want: 689,
		},
		{
			args: args{input: aoc.Input(), seconds: 2503},
			want: 1256,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := mostPoints(tt.args.input, tt.args.seconds); got != tt.want {
				t.Errorf("mostPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
