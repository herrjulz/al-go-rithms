# Shortest Path in a DAG

Find the shortest path from a source vertex `s` to all other vertices.

Algorithm:

1. `L = TopologicalSort(DAG)`
2. Set `shortest[v] = infinite`, for each vertex `v` except `s`
   Set `shortest[s] = 0`
   Set `pred[v] = null` for each vertex `v`
3. For each vertex `u`, taken the order given by `L`
	a) For each vertex `v` adj. to `u`: `Relax(u,v)`


*Relaxation Step:*

A relaxation step updates the current shortest path from a source node to the current node:

```
if shortest[u] + weight(u,v) < shortest[v], then
  set shortest[v] = shortest[u] + weight(u,v)
  set pred[v] = u
```

