package aoc

type Graph[T comparable] struct {
	Vertices map[string]*Vertex[T]
	Keys     []string
}

type Edge[T comparable] struct {
	Weight int
	Vertex *Vertex[T]
}

type Vertex[T comparable] struct {
	Key   string
	Value T
	Edges map[string]*Edge[T]
}

func (g *Graph[T]) AddVertex(key string, value T) {
	if _, ok := g.Vertices[key]; ok {
		return
	}

	g.Vertices[key] = &Vertex[T]{Key: key, Value: value, Edges: map[string]*Edge[T]{}}
	g.Keys = append(g.Keys, key)
}

func (g *Graph[T]) AddEdge(src, dst string, weight int) {
	from, ok := g.Vertices[src]
	if !ok {
		return
	}

	to, ok := g.Vertices[dst]
	if !ok {
		return
	}

	from.Edges[dst] = &Edge[T]{Weight: weight, Vertex: to}
}

func (g *Graph[T]) GetVertex(key string) (*Vertex[T], bool) {
	v, ok := g.Vertices[key]
	return v, ok
}

func (g *Graph[T]) DFS(start string, search string) []*Vertex[T] {
	nodes := []*Vertex[T]{}
	v, ok := g.Vertices[start]
	if !ok {
		return nodes
	}

	visited := Set[*Vertex[T]]{}
	var backtrack func(start *Vertex[T], search string) []*Vertex[T]
	backtrack = func(cur *Vertex[T], search string) []*Vertex[T] {
		if !visited.Add(cur) {
			return nil
		}

		if cur.Key == search {
			return []*Vertex[T]{cur}
		}

		for _, e := range cur.Edges {
			found := backtrack(e.Vertex, search)
			if found == nil {
				continue
			}

			found = append(found, cur)
			return found
		}

		return nil
	}

	return backtrack(v, search)
}
