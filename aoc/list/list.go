package list

import (
	"fmt"
	"iter"
	"strings"
)

type Linked[T comparable] struct {
	head *node[T]
}

type node[T comparable] struct {
	value T
	next  *node[T]
}

func NewLinked[T comparable](head T) *Linked[T] {
	return &Linked[T]{
		head: &node[T]{value: head},
	}
}

func LinkedWithValues[T comparable](elems []T) *Linked[T] {
	l := Linked[T]{}
	var cur *node[T]
	for i, e := range elems {
		if i == 0 {
			l.head = &node[T]{value: e}
			cur = l.head
			continue
		}

		cur.next = &node[T]{value: e}
		cur = cur.next
	}

	return &l
}

func (l *Linked[T]) Iter() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		c := l.head
		i := 0
		for c != nil {
			if !yield(i, c.value) {
				return
			}

			c = c.next
			i++
		}
	}
}

func (l *Linked[T]) Add(v T) {
	if l.head == nil {
		l.head = &node[T]{value: v}
		return
	}

	c := l.head
	for c.next != nil {
		c = c.next
	}
	c.next = &node[T]{value: v}
}

func (l *Linked[T]) Remove(v T) bool {
	c := l.head
	var prev *node[T]
	removed := false
	for c.next != nil {
		if c.value == v {
			if c == l.head {
				l.head = c.next
			}

			prev.next = c.next
			removed = true
		}

		prev = c
		c = c.next
	}

	return removed
}

func (l *Linked[T]) Tail() T {
	c := l.head
	for c.next != nil {
		c = c.next
	}
	return c.value
}

func (l *Linked[T]) Head() T {
	return l.head.value
}

func (l *Linked[T]) String() string {
	var sb strings.Builder
	for i, c := range l.Iter() {
		var s string
		if i+1 == len(s) {
			s = "%v"
		} else {
			s = "%v -> "
		}
		sb.WriteString(fmt.Sprintf(s, c))
	}
	s := sb.String()
	return s[:len(s)-4]
}
