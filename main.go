package main

import (
	"fmt"
	"jotto/exclude"
	"jotto/graph"
	"jotto/util"
	"os"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	inputPath := "./words2.txt"
	outputPath := "./words_output.txt"

	exclude.Exclude(alphabet, inputPath, outputPath)

	data, err := os.ReadFile(outputPath)
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
	util.Check(err)

	fmt.Println("done")
}
