package dijkstra

type SimplePriorityQueue struct {
	q []string
}

func NewSimplePriorityQueue() *SimplePriorityQueue {
	return &SimplePriorityQueue{q: []string{}}
}

func (s *SimplePriorityQueue) Insert(name string) {
	s.q = append(s.q, name)
}

func (s *SimplePriorityQueue) ExtractMin(shortest map[string]int) string {
	lowest := s.q[0]
	index := 0
	for i := 1; i < len(s.q); i++ {
		if shortest[s.q[i]] < shortest[lowest] {
			lowest = s.q[i]
			index = i
		}
	}
	s.remove(index)
	return lowest
}

func (s *SimplePriorityQueue) IsEmpty() bool {
	return len(s.q) == 0
}

func (s *SimplePriorityQueue) remove(i int) {
	s.q[i] = s.q[len(s.q)-1]
	s.q = s.q[:len(s.q)-1]
}
