package graph

import (
	"fmt"
	"jotto/util"
)

type CompareOutput struct {
	Overlap bool
	Merge   int32
}

func MinVertex(g *Graph) []*Node {
	res := []*Node{}

	// for i, v := range g.V {

	// }

	i := g.FirstIndex()
	compare(g.V[i], g.V[i+1])

	return res
}

func compare(a *Node, b *Node) {
	diff := a.BinaryRep() ^ b.BinaryRep()
	s := "nah"
	fmt.Println(s[1])
	fmt.Println("DIFF: ", util.InBinary(diff))
}
