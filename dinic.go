package main

import "fmt"

const INF = 1 << 30

type Edge struct {
	v, u int
	c, f int64
}

func (e *Edge) cf() int64 {
	return e.c - e.f
}

var (
	e    []Edge
	g    [][]int
	ptr  []int
	s, t int
	d    []int
)

func addEdge(v, u int, c int64) {
	g[v] = append(g[v], len(e))
	e = append(e, Edge{v, u, c, 0})
	g[u] = append(g[u], len(e))
	e = append(e, Edge{u, v, 0, 0})
}

func bfs(df int64) bool {
	for i := range d {
		d[i] = INF
	}
	d[s] = 0
	q := []int{s}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, i := range g[v] {
			if e[i].cf() >= df && d[e[i].u] > d[v]+1 {
				d[e[i].u] = d[v] + 1
				q = append(q, e[i].u)
			}
		}
	}

	return d[t] != INF
}

func dfs(v int, df int64) bool {
	if v == t {
		return true
	}
	for ; ptr[v] < len(g[v]); ptr[v]++ {
		i := g[v][ptr[v]]
		if d[v]+1 == d[e[i].u] && e[i].cf() >= df && dfs(e[i].u, df) {
			e[i].f += df
			e[i^1].f -= df
			return true
		}
	}
	return false
}

func dinic() (tot_flow int64) {
	const K = 30
	for df := int64(1 << K); df > 0; df >>= 1 {
		for bfs(df) {
			for i := range ptr {
				ptr[i] = 0
			}
			for dfs(s, df) {
				tot_flow += df
			}
		}
	}
	return
}

func scan() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	s, t = 0, n-1
	g = make([][]int, n)
	e = make([]Edge, 0)
	d = make([]int, n)
	ptr = make([]int, n)

	for i := 0; i < m; i++ {
		var v, u int
		var c int64
		fmt.Scanf("%d %d %d\n", &v, &u, &c)
		addEdge(v-1, u-1, c)
	}
}

func main() {
	scan()
	fmt.Println(dinic())
}
