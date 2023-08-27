package main

import (
	"fmt"
	"jotto/graph"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	data, err := os.ReadFile("./words.txt")
	G := graph.CreateGraph(alphabet)

	for i := 0; i < len(data); i++ {
		if data[i] == 10 {
			G.NewNode()
			continue
		}
		if data[i] == 13 {
			continue
		}
		G.ExtendWithAscii(data[i])

	}
	check(err)

	fmt.Println(G.E)
	fmt.Println(G.V)
}
