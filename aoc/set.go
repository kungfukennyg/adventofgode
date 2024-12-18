package aoc

type Set[T comparable] map[T]struct{}

func SetWithValues[T comparable](values []T) Set[T] {
	s := Set[T]{}
	s.AddAll(values)
	return s
}

// Add adds t to the Set, and returns true if the Set did not already
// contain t.
func (s Set[T]) Add(t T) bool {
	if _, ok := s[t]; ok {
		return false
	}

	s[t] = struct{}{}
	return true
}

func (s Set[T]) AddAll(v []T) {
	for _, t := range v {
		s[t] = struct{}{}
	}
}

func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}
