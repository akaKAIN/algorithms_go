package graph

import "fmt"

type VertexName = string

// Связь с другой вершиной
type Bridge struct {
	Value  int        `json:"value"`
	Target VertexName `json:"target"`
}

func (b Bridge) String() string {
	return fmt.Sprintf("<%s>: вес-%d", b.Target, b.Value)
}

// Модель вершины графа
type Vertex struct {
	Name        VertexName `json:"name"`
	MinValue    int        `json:"min_value"`
	Connections []Bridge   `json:"connections"`
}

func (v Vertex) String() string {
	return fmt.Sprintf("<%s>-соеденена: %v", v.Name, v.Connections)
}

// Граф состоящий из списка вершин и связей
type Graph struct {
	Vertices []Vertex
}

func (g Graph) String() string {
	var result string
	for _, vertex := range g.Vertices {
		result += fmt.Sprintf("%v\n", vertex)
	}
	return result
}
