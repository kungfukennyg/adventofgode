package day5

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_countNice(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{"ugknbfddgicrmopn\naaa\njchzalrnumimnmhp\nhaegwjzuvuyypxyu\ndvszwmarrgswjxmb"},
			want: 2,
		},
		{
			args: args{common.Input()},
			want: 258,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := countNice(tt.args.input); got != tt.want {
				t.Errorf("countNice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countNiceNewModel(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{"qjhvhtzxzqqjkmpb"},
			want: 1,
		},
		{
			args: args{"xxyxx"},
			want: 1,
		},
		{
			args: args{"uurcxstgmygtbstg"},
			want: 0,
		},
		{
			args: args{"ieodomkazucvgmuy"},
			want: 0,
		},
		{
			args: args{common.Input()},
			want: 53,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := countNiceNewModel(tt.args.input); got != tt.want {
				t.Errorf("countNiceNewModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
