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
	IsEmpty() bool
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
	for !d.q.IsEmpty() {
		u := d.q.ExtractMin(d.shortest)
		for _, neighbour := range graph[u].Neighbours {
			d.relax(graph[u], neighbour)
		}
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

func (d *Dijkstra) relax(u Vertex, v Neighbour) {
	if d.shortest[u.Name]+v.Weight < d.shortest[v.Name] {
		d.shortest[v.Name] = d.shortest[u.Name] + v.Weight
	}
}
