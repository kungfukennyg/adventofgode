package aoc

type Set[T comparable] map[T]struct{}

// Add adds t to the Set, and returns true if the Set did not already
// contain t.
func (s Set[T]) Add(t T) bool {
	if _, ok := s[t]; ok {
		return false
	}

	s[t] = struct{}{}
	return true
}

func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}
