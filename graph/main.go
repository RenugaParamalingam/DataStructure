package main

import "fmt"

func main() {
	g := NewGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	g.Print()

	fmt.Println("\nBFS")
	g.BFS()

	fmt.Println("\nDFS")
	g.DFS()
}

// Graph structure
type Graph struct {
	Vertex []int
	Edge   map[int][]int
}

// NewGraph initialize new graph
func NewGraph() *Graph {
	return &Graph{
		Vertex: []int{},
		Edge:   make(map[int][]int),
	}
}

// AddVertex adds new vertex
func (g *Graph) AddVertex(v int) {
	g.Vertex = append(g.Vertex, v)
}

// AddEdge adds new edge
func (g *Graph) AddEdge(u, v int) {
	g.Edge[u] = append(g.Edge[u], v)
}

// Print prints the graph
func (g *Graph) Print() {
	fmt.Print("Edges of graph")
	for v, adjacents := range g.Edge {
		fmt.Printf("\n%d -> ", v)
		for _, a := range adjacents {
			fmt.Printf("%d, ", a)
		}
	}
}

// BFS implements breath first traversal
func (g *Graph) BFS() {
	visited := make(map[int]bool)

	for _, v := range g.Vertex {
		visited[v] = false
	}

	q := []int{g.Vertex[0]}

	for len(q) > 0 {
		current := q[0]

		if !visited[current] {
			visited[current] = true

			fmt.Printf("%d -> ", current)

			adjacents := g.Edge[current]

			q = append(q, adjacents...)
		}

		q = q[1:]
	}
}

// DFS implements depth first traversal
func (g *Graph) DFS() {
	visited := make(map[int]bool)

	for _, v := range g.Vertex {
		visited[v] = false
	}

	s := []int{g.Vertex[0]}

	for len(s) > 0 {
		current := s[len(s)-1]
		s = s[:len(s)-1]

		if !visited[current] {
			visited[current] = true

			fmt.Printf("%d -> ", current)

			adjacents := g.Edge[current]

			s = append(s, adjacents...)
		}
	}
}
