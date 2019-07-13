# Topological Sort

Algorithm:

1. Init InDegree Array
  a) Create Array of size `n`
  b) Init all values with 0
  c) For each vertex `u`
    - For each vertex `v` adjecent to `u`
      - Increment in-degree[v]
1. Init `Next` Array, consisting of all vertices `u`, s.t. `InDegree[u]=0`
1. while `next` is not empty; Do
  a) Delete a vertex from `next`, call it `u`
  b) Add `u` to the end of the linar order
  c) For each vertex `v` adj. to `u`:
    - Decrement InDegree[v]
    - if `InDegree[v]=0`, then insert `v` into `next`
1. return linear order

## Running Time

- The DAG is represented as Adjecency-List
- `Next` is a stack

*Running Time:* `O(n+m)`

Explaination:

- Since `next` is a stack, add/pop is constant
- initialization of the In-Degree Slice with zero values taks `O(n)`
- Setting the In-Degree Slice takes `O(n+m)`:
  - `O(n)`, because of the outer-loop examining each of n vertices
  - `O(m)`, because of the inner-loop visiting each of the `m` edges exactly once over all of the iterations of the outer-loop
- Initialization of the `next` slice takes `O(n)`, since `next` start with at most `n` vertices. 
- Main Loop takes `n`, because each vertex is inserted exactly once.
- Add/Delete from `next` takes constant time
- Inner Main Loop iterates m times all together (incrementing and adding vertices to the result takes constant time)

=> `O(n+m)`

