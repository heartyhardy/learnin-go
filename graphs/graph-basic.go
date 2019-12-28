package graphs

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

}

//AddEdge ...
func (g *Graph) AddEdge(e int) {

}
