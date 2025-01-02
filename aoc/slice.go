package aoc

import "slices"

func Rotate[T any](values []T, n int) []T {
	values = slices.Clone(values)
	slices.Reverse(values[:n])
	slices.Reverse(values[n:])
	slices.Reverse(values)
	return values
}
