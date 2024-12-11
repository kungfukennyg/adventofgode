package day13

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

const test = `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`

func Test_bestSeats(t *testing.T) {
	type args struct {
		input     string
		personOpt opt
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{input: test},
			want: 330,
		},
		{
			args: args{input: aoc.Input()},
			want: 618,
		},
		{
			args: args{input: test, personOpt: withPerson("me", 0)},
			want: 286,
		},
		{
			args: args{input: aoc.Input(), personOpt: withPerson("me", 0)},
			want: 601,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var got int
			if tt.args.personOpt == nil {
				got = bestSeats(tt.args.input)
			} else {
				got = bestSeats(tt.args.input, tt.args.personOpt)

			}
			if got != tt.want {
				t.Errorf("bestSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}
