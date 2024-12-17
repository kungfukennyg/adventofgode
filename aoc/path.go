package aoc

// Pather is a datastructure that a pathfinding algorithm can be applied to.
type Pather interface {
	// Cost supplies the expense for moving between two vertices.
	Cost(a, b Vec) int
	// Heuristic supplies an estimate of the minimum cost from any vertex to the goal
	Heuristic(a, b Vec) int
	// Neighbors returns the neighboring vertices to p.
	Neighbors(p Vec) []Vec
	// Goal returns whether the current position matches the goal vertex.
	Goal(cur, goal Vec) bool
}

type path []Vec

func (p path) cont(n Vec) path {
	o := make(path, len(p), len(p)+1)
	copy(o, p)
	o = append(o, n)
	return o
}

func (p path) tail() Vec {
	return p[len(p)-1]
}

func (p path) cost(pather Pather) int {
	var c int
	for i := 1; i < len(p); i++ {
		c += pather.Cost(p[i-1], p[i])
	}
	return c
}

// ShortestPath implements an A* pathfinding algorithm to find the shortest path
// from the start to the goal.
func ShortestPath(pather Pather, start Vec, goal Vec) []Vec {
	closed := Set[Vec]{}

	pq := NewPriorityQueue[path]()
	pq.Init()
	pq.Push(path{start}, 0)

	for pq.Len() > 0 {
		p := pq.Pop()
		n := p.tail()
		if closed.Contains(n) {
			continue
		}

		if pather.Goal(n, goal) {
			return p
		}

		closed.Add(n)

		for _, nb := range pather.Neighbors(n) {
			cp := p.cont(nb)
			priority := -(cp.cost(pather) +
				pather.Heuristic(nb, goal))
			pq.Push(cp, priority)
		}
	}

	return nil
}
