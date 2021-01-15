package main

import "fmt"

func main() {
	// g := NewGraph(5)

	// g.AddEdges(1, 2)
	// g.AddEdges(1, 3)

	// g.Print()

	g := NewGraph()

	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)

	g.AddEdge(1, 2, false)
	g.AddEdge(1, 3, false)
	g.AddEdge(2, 3, false)

	g.PrintEdges()
}

// Graph structure
type Graph struct {
	Nodes map[int]struct{}
	Edges map[int]map[int]struct{}
}

// NewGraph creates new graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]struct{}),
		Edges: make(map[int]map[int]struct{}),
	}
}

// AddNode adds node to graph
func (g *Graph) AddNode(node int) {
	if _, ok := g.Nodes[node]; !ok {
		g.Nodes[node] = struct{}{}
	}
}

// AddEdge adds edge from u to v in graph
func (g *Graph) AddEdge(u, v int, unDirected bool) {
	if _, ok := g.Nodes[u]; !ok {
		g.Nodes[u] = struct{}{}
	}

	if _, ok := g.Nodes[v]; !ok {
		g.Nodes[v] = struct{}{}
	}

	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = make(map[int]struct{})
	}

	g.Edges[u][v] = struct{}{}

	if unDirected {
		if _, ok := g.Edges[v]; !ok {
			g.Edges[v] = make(map[int]struct{})
		}

		g.Edges[v][u] = struct{}{}
	}
}

// PrintEdges prints all edges of graph
func (g *Graph) PrintEdges() {
	for u, adjacent := range g.Edges {
		for v := range adjacent {
			fmt.Printf("Edge: %d -> %d \n", u, v)
		}
	}
}

// type Graph struct {
// 	Nodes int
// 	Edges [][]int
// }

// func NewGraph(nodes int) *Graph {
// 	return &Graph{
// 		Nodes: nodes,
// 		Edges: make([][]int, nodes),
// 	}
// }

// func (g *Graph) AddEdges(u, v int) {
// 	g.Edges[u] = append(g.Edges[u], v)
// }

// func (g *Graph) Print() {
// 	fmt.Println(g)
// 	for u, adjacent := range g.Edges {
// 		for _, v := range adjacent {
// 			fmt.Printf("Edge: %d -> %d \n", u, v)
// 		}
// 	}
// }
