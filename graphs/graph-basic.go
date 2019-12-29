package graphs

import "fmt"

//Adder ...
type Adder interface {
	AddVertex(int)
	AddEdge(int)
}

//Graph ...
type Graph struct {
	matrix map[int][]int
}

//AddVertex ...
func (g *Graph) AddVertex(v int) {
	if len(g.matrix) == 0 {
		g.matrix = make(map[int][]int)
	}
	if _, ok := g.matrix[v]; !ok {
		g.matrix[v] = make([]int, 0)
	}
}

//AddEdge ...
func (g *Graph) AddEdge(v1 int, v2 int) {
	if _, ok := g.matrix[v1]; !ok {
		g.AddVertex(v1)
	}
	if _, ok := g.matrix[v2]; !ok {
		g.AddVertex(v2)
	}
	g.matrix[v1] = append(g.matrix[v1], v2)
	g.matrix[v2] = append(g.matrix[v2], v1)
}

//Print ...
func (g *Graph) Print() {
	for k, v := range g.matrix {
		fmt.Println(".", k)
		for _, e := range v {
			fmt.Print(" -> ", e)
		}
		fmt.Println("")
	}
}
