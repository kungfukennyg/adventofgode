package aoc

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(t T) {
	s.values = append(s.values, t)
}

func (s *Stack[T]) Pop() (T, bool) {
	var t T
	if len(s.values) == 0 {
		return t, false
	}

	t = s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return t, true
}

func (s *Stack[T]) Top() (T, bool) {
	var t T
	if len(s.values) == 0 {
		return t, false
	}

	t = s.values[len(s.values)-1]
	return t, true
}

func (s *Stack[T]) Len() int {
	return len(s.values)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}
