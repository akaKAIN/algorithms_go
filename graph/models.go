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
	return fmt.Sprintf("<%s> min: %d, соеденена: %v", v.Name, v.MinValue, v.Connections)
}

// Граф состоящий из списка вершин и ребер (связей)
type Graph struct {
	Vertices []Vertex
	VertexMap map[VertexName]int
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
		g.VertexMap[vertex.Name] = max
	}
}

// Получение структуры вершины из графа по значению имени
func (g *Graph) GetVertexByName(name VertexName) (*Vertex, error) {
	for i := range g.Vertices {
		if name == g.Vertices[i].Name {
			return &g.Vertices[i], nil
		}
	}
	return nil, fmt.Errorf("Not found vertex by nama\n")
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

func (g *Graph) SearchFromTo(start, end VertexName) {
	g.FillMap()
	g.VertexMap[start] = 0
	g.addToQueue(start)
	g.queueSearch()
	g.MinWayTo(end)
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
			newValue := g.VertexMap[vertexName] + bridge.Value
			if newValue < g.VertexMap[bridgeVertex.Name] {
				g.addToQueue(bridgeVertex.Name)
				g.VertexMap[bridgeVertex.Name] = newValue
				bridgeVertex.MinValue = newValue
				log.Println(bridgeVertex)
			}
		}
	}
	log.Println("Way was build")
}

func (g *Graph) MinWayTo(end VertexName) {
	if endValue, ok := g.VertexMap[end]; ok {
		if endValue == 0 {
			return
		}

		endVertex, err := g.GetVertexByName(end)
		if err != nil {
			log.Println(err)
			return
		}
		if len(g.ShortWay) == 0 {
			g.ShortWay = append(g.ShortWay, *endVertex)
		}
		for _, bridge := range endVertex.Connections {
			bridgeVertex, err := g.GetVertexByName(bridge.Target)
			if err != nil {
				log.Println(err)
				return
			}
			if bridgeVertex.MinValue == endVertex.MinValue - bridge.Value {
				g.ShortWay = append(g.ShortWay, *bridgeVertex)
				g.MinWayTo(bridgeVertex.Name)
			}
		}
	}

}

// Формат вывода графа.
func (g Graph) String() string {
	var result string
	for _, vertex := range g.Vertices {
		result += fmt.Sprintf("%v\n", vertex)
	}
	return result
}

func (g *Graph) GetShortWayNames() []string {
	way := make([]string, 0)
	for i := range g.ShortWay {
		way = append(way, g.ShortWay[i].Name)
	}
	return way
}
