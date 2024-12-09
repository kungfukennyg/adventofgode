package day9

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_shortestDist(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{"London to Dublin = 464\nLondon to Belfast = 518\nDublin to Belfast = 141"},
			want: 605,
		},
		{
			args: args{common.Input()},
			want: 251,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := shortestDist(tt.args.input); got != tt.want {
				t.Errorf("shortestDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestDist(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{"London to Dublin = 464\nLondon to Belfast = 518\nDublin to Belfast = 141"},
			want: 982,
		},
		{
			args: args{common.Input()},
			want: 898,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := longestDist(tt.args.input); got != tt.want {
				t.Errorf("longestDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
