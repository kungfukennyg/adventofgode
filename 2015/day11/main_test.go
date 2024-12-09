package day11

import (
	"strconv"
	"testing"
)

func Test_nextValidPassword(t *testing.T) {
	type args struct {
		pass        string
		straightLen int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{pass: "abcdefgh", straightLen: 3},
			want: "abcdffaa",
		},
		{
			args: args{pass: "ghijklmn", straightLen: 3},
			want: "ghjaabcc",
		},
		{
			args: args{pass: "cqjxjnds", straightLen: 3},
			want: "cqjxxyzz",
		},
		{
			args: args{pass: "cqjxxyzz", straightLen: 3},
			want: "cqkaabcc",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := nextValidPassword(tt.args.pass, tt.args.straightLen); got != tt.want {
				t.Errorf("nextValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valid(t *testing.T) {
	type args struct {
		pass string
		n    int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{pass: "hijklmmn", n: 3},
			want: false,
		},
		{
			args: args{pass: "abbceffg", n: 3},
			want: false,
		},
		{
			args: args{pass: "abbcegjk", n: 3},
			want: false,
		},
		{
			args: args{pass: "abcdffaa", n: 3},
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := valid(tt.args.pass, tt.args.n); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
