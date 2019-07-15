package main

import (
	"fmt"

	"github.com/JulzDiverse/al-go-rithms/graphs/dijkstra"
)

func main() {
	pq := dijkstra.NewSimplePriorityQueue()
	d := dijkstra.New(pq)

	graph := map[string]dijkstra.Vertex{
		"s": {
			Name: "s",
			Neighbours: []dijkstra.Neighbour{
				{Name: "a", Weight: 6},
				{Name: "b", Weight: 4},
			},
		},
		"a": {
			Name: "a",
			Neighbours: []dijkstra.Neighbour{
				{Name: "c", Weight: 3},
				{Name: "b", Weight: 2},
			},
		},
		"b": {
			Name: "b",
			Neighbours: []dijkstra.Neighbour{
				{Name: "a", Weight: 1},
				{Name: "c", Weight: 9},
				{Name: "d", Weight: 3},
			},
		},
		"c": {
			Name: "c",
			Neighbours: []dijkstra.Neighbour{
				{Name: "d", Weight: 4},
			},
		},
		"d": {
			Name: "d",
			Neighbours: []dijkstra.Neighbour{
				{Name: "c", Weight: 5},
				{Name: "s", Weight: 7},
			},
		},
	}

	fmt.Println("RESULT:", d.ShortestPath(graph))
}
