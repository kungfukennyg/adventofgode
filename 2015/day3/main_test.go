package day3

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_housesThatGetPresent(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{">"},
			want: 2,
		},
		{
			args: args{"^>v<"},
			want: 4,
		},
		{
			args: args{"^v^v^v^v^v"},
			want: 2,
		},
		{
			args: args{common.Input()},
			want: 2565,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := housesThatGetPresent(tt.args.input); got != tt.want {
				t.Errorf("housesThatGetPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_housesWithRobot(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{"^v"},
			want: 3,
		},
		{
			args: args{"^>v<"},
			want: 3,
		},
		{
			args: args{"^v^v^v^v^v"},
			want: 11,
		},
		{
			args: args{common.Input()},
			want: 2639,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := housesWithRobot(tt.args.input); got != tt.want {
				t.Errorf("housesWithRobot() = %v, want %v", got, tt.want)
			}
		})
	}
}
