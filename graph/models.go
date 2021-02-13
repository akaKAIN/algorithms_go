package graph

import (
	"fmt"
	"log"
)

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

// Граф состоящий из списка вершин и ребер (связей)
type Graph struct {
	Vertices []Vertex
	VertexMap map[*Vertex]int
	ShortWay []Vertex
	Queue []VertexName
}

// Высчитываем максимально возможное значение суммы величин ребер графа
func (g *Graph) getMaxValue() int {
	var sum int
	for _, vertex := range g.Vertices {
		for _, val := range vertex.Connections {
			sum += val.Value
		}
	}
	return sum
}

// Наполняем значения вершин максимальными значениями для замены на меньшие значение по алгоритму Дейкстры
func (g *Graph) FillMap() {
	max := g.getMaxValue() + 1
	for _, vertex := range g.Vertices {
		g.VertexMap[&vertex] = max
	}
}

// Получение структуры вершины из графа по значению имени
func (g *Graph) GetVertexByName(name VertexName) (Vertex, error) {
	for i := range g.Vertices {
		if name == g.Vertices[i].Name {
			return g.Vertices[i], nil
		}
	}
	return Vertex{}, fmt.Errorf("Not found vertex by nama\n")
}

// Добавление имени вершины в очередь на обработку
func (g *Graph) addToQueue(vertexName VertexName) {
	g.Queue = append(g.Queue, vertexName)
}

// Получение первого элемента из очереди на обработку и удаление его из самой очереди.
func (g *Graph) deleteFromQueue() (VertexName, error) {
	if len(g.Queue) > 0 {
		vertexName := g.Queue[0]
		g.Queue = g.Queue[1:]
		return vertexName, nil
	}
	return "", fmt.Errorf("Queue is empty: nothig to delete\n")
}

func (g *Graph) SearchFrom(start VertexName) {
	startVertex, err := g.GetVertexByName(start)
	if err != nil {
		log.Fatal("Wrong start vertex name")
	}
	g.VertexMap[&startVertex] = 0
	g.addToQueue(start)
	g.queueSearch()
}

// Перебор вершин в очереди графа по алгоритму Дейкстры (без возврата)
func (g *Graph) queueSearch() {
	for len(g.Queue) > 0 {
		vertexName, err := g.deleteFromQueue()
		if err != nil {
			log.Println(err)
			return
		}
		vertex, err := g.GetVertexByName(vertexName)
		if err != nil {
			log.Println(err)
			return
		}
		for _, bridge := range vertex.Connections {
			bridgeVertex, err := g.GetVertexByName(bridge.Target)
			if err != nil {
				return
			}
			newValue := g.VertexMap[&vertex] + bridge.Value
			if newValue < g.VertexMap[&bridgeVertex] {
				g.addToQueue(bridgeVertex.Name)
				g.VertexMap[&bridgeVertex] = newValue
			}
		}
	}
	log.Println("Way was build")
}

// Формат вывода графа.
func (g Graph) String() string {
	var result string
	for _, vertex := range g.Vertices {
		result += fmt.Sprintf("%v\n", vertex)
	}
	return result
}
