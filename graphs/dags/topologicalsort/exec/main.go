package main

import (
	"fmt"

	ts "github.com/JulzDiverse/al-go-rithms/graphs/dags/topologicalsort"
)

func main() {
	dag := getDAG()
	fmt.Println("Given DAG")
	ts.PrintDAG(dag)

	fmt.Println("Resulting Linear Order")
	fmt.Println(ts.TopologicalSort(dag))
}

//Returns the DAG as Adjacency-List
func getDAG() [][]ts.Vertex {
	return [][]ts.Vertex{
		{{Number: 3, Weight: 0}},
		{{Number: 4, Weight: 0}},
		{{Number: 4, Weight: 0}, {Number: 5, Weight: 0}},
		{{Number: 6, Weight: 0}},
		{{Number: 6, Weight: 0}},
		{{Number: 7, Weight: 0}, {Number: 11, Weight: 0}},
		{{Number: 8, Weight: 0}},
		{{Number: 13, Weight: 0}},
		{{Number: 10, Weight: 0}},
		{{Number: 11, Weight: 0}},
		{{Number: 12, Weight: 0}},
		{{Number: 13, Weight: 0}},
		{{Number: 14, Weight: 0}},
		{{Number: 0, Weight: 0}},
	}
}
