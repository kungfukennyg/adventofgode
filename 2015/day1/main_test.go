package day1

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_findFloor(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{common.Input()},
			want: 0,
		},
		{
			args: args{"(())"},
			want: 0,
		},
		{
			args: args{"()()"},
			want: 0,
		},
		{
			args: args{"((("},
			want: 3,
		},
		{
			args: args{"(()(()("},
			want: 3,
		},
		{
			args: args{"))((((("},
			want: 3,
		},
		{
			args: args{"())"},
			want: -1,
		},
		{
			args: args{"))("},
			want: -1,
		},
		{
			args: args{")))"},
			want: -3,
		},
		{
			args: args{")())())"},
			want: -3,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := findFloor(tt.args.input); got != tt.want {
				t.Errorf("findFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstEnterBasement(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{")"},
			want: 1,
		},
		{
			args: args{"()())"},
			want: 5,
		},
		{
			args: args{common.Input()},
			want: 1797,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := firstEnterBasement(tt.args.input); got != tt.want {
				t.Errorf("firstEnterBasement() = %v, want %v", got, tt.want)
			}
		})
	}
}
