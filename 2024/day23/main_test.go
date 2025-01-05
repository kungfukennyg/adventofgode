package day23

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_interconnected(t *testing.T) {
	type args struct {
		input  string
		filter string
		depth  int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input:  testInput,
				filter: "",
				depth:  3,
			},
			want: 12,
		},
		{
			args: args{
				input:  testInput,
				filter: "t",
				depth:  3,
			},
			want: 7,
		},
		{
			args: args{
				input:  aoc.Input(),
				filter: "t",
				depth:  3,
			},
			want: 926,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := interconnected(tt.args.input, tt.args.filter, tt.args.depth); got != tt.want {
				t.Errorf("interconnected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lanPassword(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{input: testInput},
			want: "co,de,ka,ta",
		},
		{
			args: args{input: aoc.Input()},
			want: "az,ed,hz,it,ld,nh,pc,td,ty,ux,wc,yg,zz",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := lanPassword(tt.args.input); got != tt.want {
				t.Errorf("lanPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

const testInput = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`
