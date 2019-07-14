package dagshortestpath

import (
	ts "github.com/JulzDiverse/al-go-rithms/graphs/dags/topologicalsort"
)

//ShortestPath calculates the shortest path value from the source node to any other node in the DAG
func ShortestPath(dag [][]ts.Vertex) []int {
	l := ts.TopologicalSort(dag)

	shortest := make([]int, len(l))
	for i := range shortest {
		shortest[i] = 9999
	}
	shortest[0] = 0

	pred := make([]int, len(l))

	for _, u := range l {
		for _, v := range dag[u-1] {
			relax(shortest, pred, u, v.Number, v.Weight)
		}
	}

	return shortest
}

func relax(shortest []int, pred []int, u int, v int, weight int) {
	if v == 0 {
		return
	}

	if shortest[u-1]+weight < shortest[v-1] {
		shortest[v-1] = shortest[u-1] + weight
		pred[v-1] = u
	}
}
