package dijkstra

type Vertex struct {
	Name       string
	Neighbours []Neighbour
	SPValue    int
	Pred       string
}

type Neighbour struct {
	Name   string
	Weight int
}

type PriorityQueue interface {
	Insert(string)
	ExtractMin(map[string]int) string
}

type Dijkstra struct {
	q        PriorityQueue
	shortest map[string]int
	pred     map[string]string
}

func New(pq PriorityQueue) *Dijkstra {
	return &Dijkstra{
		q:        pq,
		shortest: map[string]int{},
		pred:     map[string]string{},
	}
}

func (d *Dijkstra) ShortestPath(graph map[string]Vertex) map[string]int {
	d.setupShortestAndPred(graph)
	d.setupQ(graph)
	done := len(graph)
	for done != 0 {
		u := d.q.ExtractMin(d.shortest) //TODO
		for _, neighbour := range graph[u].Neighbours {
			relax(d.shortest, d.pred, graph[u], neighbour)
		}
		done--
	}
	return d.shortest
}

func (d *Dijkstra) setupShortestAndPred(graph map[string]Vertex) {
	for k := range graph {
		d.shortest[k] = 99999
		d.pred[k] = "undefined"
	}
	d.shortest["s"] = 0
}

func (d *Dijkstra) setupQ(graph map[string]Vertex) {
	for k := range graph {
		d.q.Insert(k)
	}
}

func relax(shortest map[string]int, pred map[string]string, u Vertex, v Neighbour) {
	if shortest[u.Name]+v.Weight < shortest[v.Name] {
		shortest[v.Name] = shortest[u.Name] + v.Weight
	}
}
