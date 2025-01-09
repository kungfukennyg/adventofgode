package day17

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_cpu(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args          args
		wantRegisters map[string]uint64
		wantOutput    string
	}{
		{
			args: args{`Register C: 9
Program: 2,6`},
			wantRegisters: map[string]uint64{"b": 1},
		},
		{
			args: args{`Register A: 10
Program: 5,0,5,1,5,4`},
			wantOutput: "0,1,2",
		},
		{
			args: args{`Register A: 2024
Program: 0,1,5,4,3,0`},
			wantRegisters: map[string]uint64{"a": 0},
			wantOutput:    "4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			args: args{`Register B: 29
Program: 1,7`},
			wantRegisters: map[string]uint64{"b": 26},
		},
		{
			args: args{`Register B: 2024
Register C: 43690
Program: 4,0`},
			wantRegisters: map[string]uint64{"b": 44354},
		},
		{
			args:          args{aoc.Input()},
			wantRegisters: map[string]uint64{},
			wantOutput:    "7,0,3,1,2,6,3,7,1",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cpu := parse(tt.args.input)
			out := cpu.Run()
			if tt.wantRegisters != nil {
				if wantA, ok := tt.wantRegisters["a"]; ok && wantA != cpu.a {
					t.Errorf("cpu.a = %v, want %v", cpu.a, wantA)
				}
				if wantB, ok := tt.wantRegisters["b"]; ok && wantB != cpu.b {
					t.Errorf("cpu.b = %v, want %v", cpu.b, wantB)
				}
				if wantC, ok := tt.wantRegisters["c"]; ok && wantC != cpu.c {
					t.Errorf("cpu.c = %v, want %v", cpu.c, wantC)
				}
			}
			if tt.wantOutput != "" && tt.wantOutput != out {
				t.Errorf("Run() = %v, want %v", out, tt.wantOutput)
			}
		})
	}
}

func Test_inputFromProgram(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want uint64
	}{
		{
			args: args{`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`},
			want: 117440,
		},
		{
			args: args{aoc.Input()},
			want: 109020013201563,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := inputFromProgram(tt.args.input); got != tt.want {
				t.Errorf("inputFromProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}
