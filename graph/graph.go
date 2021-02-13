package graph

import "fmt"

func GenerateGraphFromMatrix(matrix [][]int) *Graph {
	g := Graph{}
	for i, row := range matrix {
		name := fmt.Sprintf("V%d", i)
		connections := make([]Bridge, 0)
		vertex := Vertex{
			Name:        name,
			Connections: nil,
			MinValue:    0,
		}
		for j, val := range row {
			if val == 0 {
				continue
			}
			name := fmt.Sprintf("V%d", j)
			bridge := Bridge{
				Value:  val,
				Target: name,
			}
			connections = append(connections, bridge)
		}
		vertex.Connections = connections
		g.Vertices = append(g.Vertices, vertex)
		g.Queue = make([]VertexName, 0)
		g.VertexMap = make(map[VertexName]int)
	}
	return &g
}