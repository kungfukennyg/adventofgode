package aoc

// Graph is a generic graph of vertices and nodes. Graph implicitly supports
// weighted/unweighted and directed/undirected connections. Graph is not
// intended nor safe to be shared between threads.
type Graph[T comparable] struct {
	Vertices map[string]*Vertex[T]
	Keys     []string
}

// Edge is an edge between two vertices in a Graph.
type Edge[T comparable] struct {
	Weight int
	Vertex *Vertex[T]
}

// Vertex is a vertex in a Graph.
type Vertex[T comparable] struct {
	Key   string
	Value T
	Edges map[string]*Edge[T]
}

// Neighbors returns the vertices along each edge of this Vertex.
func (v *Vertex[T]) Neighbors() []*Vertex[T] {
	nbs := make([]*Vertex[T], 0, len(v.Edges))
	for _, e := range v.Edges {
		nbs = append(nbs, e.Vertex)
	}
	return nbs
}

// NewGraph creates a new empty graph.
func NewGraph[T comparable]() *Graph[string] {
	return &Graph[string]{
		Vertices: map[string]*Vertex[string]{},
		Keys:     []string{},
	}
}

// AddVertex adds a vertex to the graph.
func (g *Graph[T]) AddVertex(key string, value T) {
	if _, ok := g.Vertices[key]; ok {
		return
	}

	g.Vertices[key] = &Vertex[T]{Key: key, Value: value, Edges: map[string]*Edge[T]{}}
	g.Keys = append(g.Keys, key)
}

// AddEdge adds an edge from the src vertex to the dst vertex. Both vertices
// must exist in the graph.
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

// GetVertex returns the vertex matching key and true, if the vertex is present,
// or nil and false otherwise.
func (g *Graph[T]) GetVertex(key string) (*Vertex[T], bool) {
	v, ok := g.Vertices[key]
	return v, ok
}

// DFS performas a depth-first search of the graph, starting at start and
// ending at search, and returns the traversed path of vertices. DFS returns
// an empty slice if either the start or the end vertex does not exist.
func (g *Graph[T]) DFS(start, search string) []*Vertex[T] {
	v, ok := g.Vertices[start]
	if !ok {
		return []*Vertex[T]{}
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

	if found := backtrack(v, search); found != nil {
		return found
	}

	return []*Vertex[T]{}
}

// Cliques finds all maximal cliques, or sets of vertices that all connect to each
// other.
//
// See: https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func (g *Graph[T]) Cliques() []Set[*Vertex[T]] {
	return g.CliquesFunc(func(r, p, x Set[*Vertex[T]]) bool {
		// if p and x are empty, we have no new vertices that will fit in r.
		// therefore, r is a maximal set.
		return len(p) == 0 && len(x) == 0
	})
}

// CliquesN finds all cliques, or sets of vertices that all connect to each
// other, of length n.
//
// See: https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func (g *Graph[T]) CliquesN(n int) []Set[*Vertex[T]] {
	return g.CliquesFunc(func(r, p, x Set[*Vertex[T]]) bool {
		return len(r) == n
	})
}

// CliquesFunc finds all cliques, or sets of vertices that all connect to each
// other, that match the provided predicate function.
//
// See: https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func (g *Graph[T]) CliquesFunc(pred func(r, p, x Set[*Vertex[T]]) bool) []Set[*Vertex[T]] {
	cliques := []Set[*Vertex[T]]{}
	var backtrack func(r, p, x Set[*Vertex[T]])
	backtrack = func(r, p, x Set[*Vertex[T]]) {
		if pred(r, p, x) {
			cliques = append(cliques, r)
		}

		if len(p) == 0 && len(x) == 0 {
			return
		}

		for v := range p {
			nbs := SetWithValues(v.Neighbors())
			backtrack(r.Union(SetWith(v)), p.Intersect(nbs), x.Intersect(nbs))
			p.Remove(v)
			x.Add(v)
		}
	}

	r, p, x := Set[*Vertex[T]]{}, Set[*Vertex[T]]{}, Set[*Vertex[T]]{}
	for _, v := range g.Vertices {
		p.Add(v)
	}

	backtrack(r, p, x)
	return cliques
}
