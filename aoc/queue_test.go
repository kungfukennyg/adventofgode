package aoc

import (
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	tests := []struct {
		name       string
		priorities map[string]int
		transform  func(*PriorityQueue[string])
		want       []string
	}{
		{
			name:       "Push",
			priorities: map[string]int{"banana": 3, "apple": 2, "pear": 4},
			transform: func(pq *PriorityQueue[string]) {
				pq.Push("orange", 1)
			},
			want: []string{"pear", "banana", "apple", "orange"},
		},
		{
			name:       "Update",
			priorities: map[string]int{"banana": 3, "apple": 2, "pear": 4},
			transform: func(pq *PriorityQueue[string]) {
				v := "orange"
				idx := pq.Push(v, 1)
				pq.Update(idx, v, 5)
			},
			want: []string{"orange", "pear", "banana", "apple"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewPriorityQueue[string]()
			for k, v := range tt.priorities {
				pq.Push(k, v)
			}
			tt.transform(pq)

			if got := pq.Consume(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPriorityQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
