package day1

import (
	"strconv"
	"testing"
)

func Test_hqDistance(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "R2, L3",
			want:  5,
		},
		{
			input: "R2, R2, R2",
			want:  2,
		},
		{
			input: "R5, L5, R5, R3",
			want:  12,
		},
		{
			input: input,
			want:  252,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := hqDistance(tt.input)
			if got != tt.want {
				t.Errorf("hqDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstPlaceVisitedTwice(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "R8, R4, R4, R8",
			want:  4,
		},
		{
			input: input,
			want:  143,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := firstPlaceVisitedTwice(tt.input)
			if got != tt.want {
				t.Errorf("firstPlaceVisitedTwice() = %v, want %v", got, tt.want)
			}
		})
	}
}

const input = `L3, R1, L4, L1, L2, R4, L3, L3, R2, R3, L5, R1, R3, L4, L1, L2, R2, R1, L4, L4, R2, L5, R3, R2, R1, L1, L2, R2, R2, L1, L1, R2, R1, L3, L5, R4, L3, R3, R3, L5, L190, L4, R4, R51, L4, R5, R5, R2, L1, L3, R1, R4, L3, R1, R3, L5, L4, R2, R5, R2, L1, L5, L1, L1, R78, L3, R2, L3, R5, L2, R2, R4, L1, L4, R1, R185, R3, L4, L1, L1, L3, R4, L4, L1, R5, L5, L1, R5, L1, R2, L5, L2, R4, R3, L2, R3, R1, L3, L5, L4, R3, L2, L4, L5, L4, R1, L1, R5, L2, R4, R2, R3, L1, L1, L4, L3, R4, L3, L5, R2, L5, L1, L1, R2, R3, L5, L3, L2, L1, L4, R4, R4, L2, R3, R1, L2, R1, L2, L2, R3, R3, L1, R4, L5, L3, R4, R4, R1, L2, L5, L3, R1, R4, L2, R5, R4, R2, L5, L3, R4, R1, L1, R5, L3, R1, R5, L2, R1, L5, L2, R2, L2, L3, R3, R3, R1`
