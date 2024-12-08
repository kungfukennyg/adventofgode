package day7

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_cpu(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		wire string
		want uint16
	}{
		{
			args: args{common.Input()},
			wire: "a",
			want: 3176,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := newCpu(tt.args.input)
			got := c.read(tt.wire)
			if got != tt.want {
				t.Errorf("%s = %v, want %v\n", tt.wire, got, tt.want)
			}
		})
	}
}
func Test_override(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args     args
		wire     string
		override string
		want     uint16
	}{
		{
			args:     args{common.Input()},
			wire:     "a",
			override: "b",
			want:     3176,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := newCpu(tt.args.input)
			c.override(tt.override, c.read(tt.wire))

			got := c.read(tt.wire)
			if got != tt.want {
				t.Errorf("%s = %v, want %v\n", tt.wire, got, tt.want)
			}
		})
	}
}
