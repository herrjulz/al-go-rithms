# Dijkstra Algorithm

Find the shortest path from a single source vertex to all other vertices. 

- The graph can have cycles
- The edge weight must be nonnegative

## Algorithm

- Let `shortest` be a set containing the shortest values for each vertex
- Let `pred` be a set containing the predecessor for each vertex
- Let `Q` be a set of vertices for which the final `shortest` and `pred` values are not known yet

**Idea**:

Iterate over the set `Q` and find the `vertex` with the lowest shortest path value. In the beginning it is always the `source vertex`. Remove this `vertex` from Q and run the relaxation step for each adjecent vertex. The algorithm terminates when Q is empty. 

```
1. shortest[v] = infinite, for all vertex v except s
   shortest[s] = 0
   pred[v] = null, for all vertex v
2. Set Q to contain all vertices
3. While Q is not empty, do:
	1. Find the vertex u in set Q with the lowest shortest value and remove it from Q
	2. For each vertex adjecent to u, do:
		Relax(u,v)
```

The Algorithm requires a `PriorityQueue` (Datastructure) that is defined by the following operations:

- `Insert`: Inserts an Item into the `PriorityQueue`.
- `ExtractMin`: Extracts - removes and returns - the item from the `PriorityQueue` with the lowest shortest value to the caller.

## Loop Invariant

At the start of each iteration of the loop in step 3, `shortest[v] = sp(s, v)` for each vertex `v` not in the set `Q`. That is, for each vertex `v` not in `Q`, the value of `shortest[v]` is the weight of a shortest path from `s` to `v`.

**Explained**:

- In each iteration of the main loop the shortest value for each vertex `v` that is not in set `Q` is the final one. 
- The shortest value of the vertex `u` that was extracted from `Q` in any iteration of the main loop, can never again decrease. Because:
	- The only remaining edges that need to be relaxed, are edges leaving vertices in `Q`. 
	- Every vertex in set Q has s shortest value at least as large as the current lowest vertex `u`.  
	- Since all edge weights are nonnegative, we must have `shortest[u] ≤ shortest[v] + weight(v, u)` for every vertex `v` in `Q`. 
