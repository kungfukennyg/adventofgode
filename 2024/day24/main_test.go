package day24

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_readRegister(t *testing.T) {
	type args struct {
		input  string
		target string
	}
	tests := []struct {
		args args
		want int64
	}{
		{
			args: args{
				input: `x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -> z00
x01 XOR y01 -> z01
x02 OR y02 -> z02`,
				target: "z",
			},
			want: 4,
		},
		{
			args: args{
				input:  testInput,
				target: "z",
			},
			want: 2024,
		},
		{
			args: args{
				input:  aoc.Input(),
				target: "z",
			},
			want: 57588078076750,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := run(tt.args.input, tt.args.target); got != tt.want {
				t.Errorf("readRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findBug(t *testing.T) {
	type args struct {
		input  string
		target string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				input:  aoc.Input(),
				target: "z",
			},
			want: "kcd,pfn,shj,tpk,wkb,z07,z23,z27",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := findCrossed(tt.args.input, tt.args.target); got != tt.want {
				t.Errorf("findBug() = %v, want %v", got, tt.want)
			}
		})
	}
}

const testInput = `x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj`
