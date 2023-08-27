package graph

import (
	"fmt"
)

type Graph struct {
	V        []*Node
	E        []*Edge
	Current  *Node
	Alphabet string
}

type Node struct {
	indices []int
}

func (n *Node) String() string {
	output := ""
	for _, x := range n.indices {
		o := fmt.Sprintf(" %d", x)
		output += o
	}
	return output
}

type Edge struct {
	u *Node
	v *Node
}

func (e *Edge) String() string {
	return fmt.Sprintf("E %d-%d", e.u, e.v)
}

func (g *Graph) ExtendWithAscii(b byte) {
	switch {
	case b <= 64:
		return
	case b >= 123:
		return
	case b > 64 && b <= 90:
		b += 32
		fallthrough
	case b < 123:
		letter := int(b) - 97
		g.Current.indices = append(g.Current.indices, letter)
		g.ConnectCurrentWith(letter)
		return

	default:
		fmt.Println("hiiigh")
	}
}

func (g *Graph) ConnectCurrentWith(index int) {
	g.E = append(g.E, &Edge{u: g.V[index], v: g.Current})
}

func CreateGraph(Alphabet string) *Graph {
	G := Graph{V: []*Node{}, Alphabet: Alphabet}

	for i, _ := range Alphabet {
		G.V = append(G.V, &Node{indices: []int{i}})
	}
	G.NewNode()
	return &G
}

func (g *Graph) NewNode() {
	n := &Node{indices: []int{}}
	g.Current = n
	g.V = append(g.V, g.Current)
}
