package main

import (
	"heartyhardy/graphs"
)

func main() {
	g := new(graphs.Graph)
	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.Print()
}
