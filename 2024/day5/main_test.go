package day5

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_correctPageNumbers(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args           args
		correct, wrong int
	}{
		{
			args:    args{testInput},
			correct: 143,
			wrong:   123,
		},
		{
			args:    args{common.Input()},
			correct: 6260,
			wrong:   5346,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if correct, wrong := correctPageNumbers(tt.args.input); correct != tt.correct || wrong != tt.wrong {
				t.Errorf("correctPageNumbers(), correct = %v, want %v, wrong = %v, want %v", correct, tt.correct, wrong, tt.wrong)
			}
		})
	}
}

const testInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
