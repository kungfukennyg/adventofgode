package aoc

import (
	"iter"
	"maps"
)

type Set[T comparable] map[T]struct{}

func SetWith[T comparable](value T) Set[T] {
	s := Set[T]{}
	s.Add(value)
	return s
}

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

func (s Set[T]) Remove(t T) bool {
	_, ok := s[t]
	delete(s, t)
	return ok
}

func (s Set[T]) ContainsAll(ts []T) bool {
	for _, t := range ts {
		if !s.Contains(t) {
			return false
		}
	}

	return true
}

func (s Set[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for t := range s {
			if !yield(t) {
				return
			}
		}
	}
}

func (s Set[T]) Intersect(o Set[T]) Set[T] {
	ret := Set[T]{}
	for t := range s {
		if o.Contains(t) {
			ret.Add(t)
		}
	}
	return ret
}

func (s Set[T]) Union(o Set[T]) Set[T] {
	s = maps.Clone(s)
	for t := range o {
		s.Add(t)
	}
	return s
}
