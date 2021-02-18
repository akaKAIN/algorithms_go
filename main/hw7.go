package main

import (
	"fmt"
	"github.com/akaKAIN/algorithms_go/graph"
	"github.com/akaKAIN/algorithms_go/input"
	"log"
	"sort"
)

func main() {
	fileName := "graph_matrix.txt"

	matrix, err := input.UserFileWithMatrix(fileName)
	if err != nil {
		log.Println(err)
	}
	g := graph.GenerateGraphFromMatrix(matrix)
	g.SearchFromTo("V0", "V4")
	shortWay := g.GetShortWayNames()
	sort.Strings(shortWay)
	fmt.Println(shortWay)
}
