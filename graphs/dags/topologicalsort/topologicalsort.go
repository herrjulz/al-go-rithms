package main

import (
	"fmt"
)

func main() {
	dag := getDAG()
	fmt.Println("Given DAG")
	PrintDAG(dag)

	fmt.Println("Resulting Linear Order")
	fmt.Println(topologicalSort(dag))
}

type NextStack struct {
	Stack []int
}

func (ns *NextStack) Add(node int) {
	ns.Stack = append(ns.Stack, node)
}

func (ns *NextStack) Pop() int {
	node := ns.Stack[len(ns.Stack)-1]
	ns.Stack = ns.Stack[0 : len(ns.Stack)-1]
	return node
}

func topologicalSort(dag [][]int) []int {
	inDegree := initInDegree(dag)
	next := initNext(inDegree)
	linearOrder := []int{}

	for len(next.Stack) != 0 {
		u := next.Pop()
		linearOrder = append(linearOrder, u+1)
		for _, v := range dag[u] {
			if v != 0 {
				v--
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

func initInDegree(dag [][]int) []int {
	inDegree := make([]int, len(dag))
	for _, nodes := range dag {
		for _, edges := range nodes {
			//Do not increment for nodes with no leaving edge (represented as 0)
			if edges != 0 {
				inDegree[edges-1]++
			}
		}
	}
	return inDegree
}

func PrintDAG(dag [][]int) {
	for i, v := range dag {
		fmt.Printf("%d: %v\n", i+1, v)
	}
}

//Returns the DAG as Adjacency-List
func getDAG() [][]int {
	return [][]int{
		{3},
		{4},
		{4, 5},
		{6},
		{6},
		{7, 11},
		{8},
		{13},
		{10},
		{11},
		{12},
		{13},
		{14},
		{0},
	}
}
