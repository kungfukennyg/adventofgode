package day10

import (
	"strconv"
	"testing"
)

func Test_lookAndSay(t *testing.T) {
	type args struct {
		input string
		n     int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{input: "211", n: 1},
			want: 4,
		},
		{
			args: args{input: "1", n: 5},
			want: 6,
		},
		{
			args: args{input: "1321131112", n: 40},
			want: 492982,
		},
		{
			args: args{input: "1321131112", n: 50},
			want: 492982,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := lookAndSay(tt.args.input, tt.args.n); len(got) != tt.want {
				t.Errorf("lookAndSay() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
