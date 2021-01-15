package main

import "fmt"

func main() {
	g := NewGraph(5)

	g.AddEdges(1, 2)
	g.AddEdges(1, 3)

	g.PrintEdges()

	fmt.Println("\n arbitary labeled graph")
	ag := NewArbitaryLabeledGraph()

	ag.AddNode(1)
	ag.AddNode(2)
	ag.AddNode(3)

	ag.AddEdge(1, 2, false)
	ag.AddEdge(1, 3, false)
	ag.AddEdge(2, 3, false)

	ag.PrintEdges()
}

// Graph structure
type Graph struct {
	Nodes int
	Edges [][]int
}

// NewGraph creates new graph
func NewGraph(nodes int) *Graph {
	return &Graph{
		Nodes: nodes,
		Edges: make([][]int, nodes),
	}
}

// AddEdges adds edge from u to v in graph
func (g *Graph) AddEdges(u, v int) {
	g.Edges[u] = append(g.Edges[u], v)
}

// PrintEdges prints all edges of graph
func (g *Graph) PrintEdges() {
	fmt.Println(g)

	for u, adjacent := range g.Edges {
		for _, v := range adjacent {
			fmt.Printf("Edge: %d -> %d \n", u, v)
		}
	}
}

// ArbitaryLabeledGraph is graph structure
type ArbitaryLabeledGraph struct {
	Nodes map[int]struct{}
	Edges map[int]map[int]struct{}
}

// NewArbitaryLabeledGraph creates new graph
func NewArbitaryLabeledGraph() *ArbitaryLabeledGraph {
	return &ArbitaryLabeledGraph{
		Nodes: make(map[int]struct{}),
		Edges: make(map[int]map[int]struct{}),
	}
}

// AddNode adds node to graph
func (g *ArbitaryLabeledGraph) AddNode(node int) {
	if _, ok := g.Nodes[node]; !ok {
		g.Nodes[node] = struct{}{}
	}
}

// AddEdge adds edge from u to v in graph
func (g *ArbitaryLabeledGraph) AddEdge(u, v int, unDirected bool) {
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
func (g *ArbitaryLabeledGraph) PrintEdges() {
	fmt.Println(g)
	for u, adjacent := range g.Edges {
		for v := range adjacent {
			fmt.Printf("Edge: %d -> %d \n", u, v)
		}
	}
}
