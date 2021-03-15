package main

type graphNode struct {
	id    int
	edges map[int]int
}
type graph struct {
	nodes []*graphNode
}

func New() *graph {
	return &graph{
		nodes: []*graphNode{},
	}
}

func (g graph) addNode(a int) {
	n := graphNode{id: a}
	g.nodes = append(g.nodes, &n)
}

func (g *graph) addEdge(a int, b int, w int) {
	for _, v := range g.nodes {
		if v.id == a {
			v.edges[b] = w
		}
	}

}

func main() {
	g := New()
	g.addNode(1)
	g.addNode(3)
	g.addEdge(1, 3, 5)

}
