package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/graph"
	"github.com/akaKAIN/algorithms_go/input"
	"log"
)

func main() {
	fileName := "graph_matrix.txt"

	matrix, err := input.UserFileWithMatrix(fileName)
	if err != nil {
		log.Println(err)
	}
	g := graph.GenerateGraphFromMatrix(matrix)
	fmt.Println(g)
	g.SearchFrom("V0")
	for val := range g.VertexMap {
		fmt.Println("::", val)
	}
}
