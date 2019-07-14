package dagshortestpath_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/al-go-rithms/graphs/dags/shortestpath"
	ts "github.com/JulzDiverse/al-go-rithms/graphs/dags/topologicalsort"
)

var _ = Describe("Dagshortestpath", func() {

	Context("For a given DAG", func() {

		var dag [][]ts.Vertex

		BeforeEach(func() {
			dag = [][]ts.Vertex{
				{{2, 5}, {3, 3}},
				{{3, 2}, {4, 6}},
				{{4, 7}, {5, 4}, {6, 2}},
				{{5, -1}, {6, 1}},
				{{6, -2}},
				{{0, 0}},
			}
		})

		It("gets the shortest path", func() {
			shortest := ShortestPath(dag)

			for i, v := range shortest {
				fmt.Printf("Shortest Path Value from source node 1 to %d is %d\n", i+1, v)
			}
		})
	})
})
