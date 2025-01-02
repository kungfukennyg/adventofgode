package seq

import (
	"fmt"
	"iter"
)

type Map[K comparable, V any] struct {
	m       map[K]V
	filters []func(K, V) bool
}

type MapReducer[K comparable, V any, T any] func(T, K, V) T

func NewMap[K comparable, T any](m map[K]T) Map[K, T] {
	return Map[K, T]{m: m}
}

func GroupByValues[K comparable, V comparable](m map[K]V) Map[V, []K] {
	o := map[V][]K{}
	for k, v := range m {
		if _, ok := o[v]; !ok {
			o[v] = []K{k}
		} else {
			o[v] = append(o[v], k)
		}
	}
	return NewMap(o)
}

func (seq Map[K, V]) Values() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
	outer:
		for k, v := range seq.m {
			for _, pred := range seq.filters {
				if !pred(k, v) {
					continue outer
				}
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

func (seq Map[K, V]) Slice() []V {
	ret := make([]V, 0, len(seq.m))
	for _, v := range seq.Values() {
		ret = append(ret, v)
	}
	return ret
}

func (seq Map[K, V]) ForEach(fn func(K, V)) {
	for k, v := range seq.Values() {
		fn(k, v)
	}
}

func (seq Map[K, V]) Filter(pred func(K, V) bool) Map[K, V] {
	seq.filters = append(seq.filters, pred)
	return seq
}

func (seq Map[K, V]) Len() int {
	n := 0
	for range seq.Values() {
		n++
	}
	return n
}

func (seq Map[K, V]) Sum() int {
	return seq.ReduceInt(0, func(sum int, k K, v int) int { return sum + v })
}

func (seq Map[K, V]) ReduceInt(initial int, reducer MapReducer[K, int, int]) int {
	rdc := initial
	for k, v := range seq.Values() {
		vv, ok := any(v).(int)
		if !ok {
			panic(fmt.Errorf("tried to assert 'v' (#%T) is 'int'", v))
		}

		rdc = reducer(initial, k, vv)
	}
	return rdc
}

func (seq Map[K, V]) ReduceString(initial string, reducer MapReducer[K, string, string]) string {
	rdc := initial
	for k, v := range seq.Values() {
		vv, ok := any(v).(string)
		if !ok {
			panic(fmt.Errorf("tried to assert 'v' (#%T) is 'string'", v))
		}

		rdc = reducer(initial, k, vv)
	}
	return rdc
}
