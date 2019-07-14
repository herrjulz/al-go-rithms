package topologicalsort

import (
	"fmt"
)

//Vertex is a Representation of a Weighted Vertex
type Vertex struct {
	Number int
	Weight int
}

//NextStack is a represntation of a Stack
type NextStack struct {
	Stack []int
}

//Add adds an item to the stack
func (ns *NextStack) Add(node int) {
	ns.Stack = append(ns.Stack, node)
}

//Pop removes an item from the stack
func (ns *NextStack) Pop() int {
	node := ns.Stack[len(ns.Stack)-1]
	ns.Stack = ns.Stack[0 : len(ns.Stack)-1]
	return node
}

//TopologicalSort sorts a DAG
func TopologicalSort(dag [][]Vertex) []int {
	inDegree := initInDegree(dag)
	next := initNext(inDegree)
	linearOrder := []int{}

	for len(next.Stack) != 0 {
		u := next.Pop()
		linearOrder = append(linearOrder, u+1)
		for _, vertex := range dag[u] {
			if vertex.Number != 0 {
				v := vertex.Number - 1
				inDegree[v]--
				if inDegree[v] == 0 {
					next.Add(v)
				}
			}
		}
	}
	return linearOrder
}

func initNext(inDegree []int) NextStack {
	next := NextStack{}
	for i, v := range inDegree {
		if v == 0 {
			next.Add(i)
		}
	}
	return next
}

func initInDegree(dag [][]Vertex) []int {
	inDegree := make([]int, len(dag))
	for _, vertices := range dag {
		for _, vertex := range vertices {
			//Do not increment for nodes with no leaving edge (represented as 0)
			if vertex.Number != 0 {
				inDegree[vertex.Number-1]++
			}
		}
	}
	return inDegree
}

//PrintDAG prints a DAG
func PrintDAG(dag [][]Vertex) {
	for i, v := range dag {
		fmt.Printf("%d: %v\n", i+1, v)
	}
}
