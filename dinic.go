package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

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

func scanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func makeScanner(graph_file *string) *bufio.Scanner {
	var scanner *bufio.Scanner
	if *graph_file != "" {
		f, _ := os.Open(*graph_file)
		scanner = bufio.NewScanner(f)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)

	return scanner
}

func scanGraph(scanner *bufio.Scanner, is_undirected *bool) {
	n, m := scanInt(scanner), scanInt(scanner)

	s, t = 0, n-1
	g = make([][]int, n)
	e = make([]Edge, 0)
	d = make([]int, n)
	ptr = make([]int, n)

	for i := 0; i < m; i++ {
		v, u, c := scanInt(scanner), scanInt(scanner), int64(scanInt(scanner))
		addEdge(v-1, u-1, c)
		if *is_undirected {
			addEdge(u-1, v-1, c)
		}
	}
}

func makeWriter(output_file *string) *bufio.Writer {
	var writer *bufio.Writer
	if *output_file != "" {
		f, _ := os.Create(*output_file)
		writer = bufio.NewWriter(f)
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}

	return writer
}

func printEdgeFlows(writer *bufio.Writer, is_undirected *bool, no_empty_edges *bool) {
	if *is_undirected {
		for i := 0; i < len(e); i += 4 {
			flow := e[i].f - e[i+2].f
			if flow == 0 && *no_empty_edges {
				continue
			}
			if flow >= 0 {
				writer.WriteString(fmt.Sprint(i/4+1, " | ", e[i].v+1, " -> ", e[i].u+1, " | flow = ", flow, " | capacity = ", e[i].c, "\n"))
			} else {
				writer.WriteString(fmt.Sprint(i/4+1, " | ", e[i].v+1, " <- ", e[i].u+1, " | flow = ", -flow, " | capacity = ", e[i].c, "\n"))
			}
		}
	} else {
		for i := 0; i < len(e); i += 2 {
			if e[i].f == 0 && *no_empty_edges {
				continue
			}
			writer.WriteString(fmt.Sprint(i/2+1, " | ", e[i].v+1, " -> ", e[i].u+1, " | flow = ", e[i].f, " | capacity = ", e[i].c, "\n"))
		}
	}
}

func printFlow(writer *bufio.Writer, flow int64, edge_flows *bool, is_undirected *bool, no_empty_edges *bool) {
	writer.WriteString(fmt.Sprint("Max flow: ", flow, "\n"))
	if *edge_flows {
		printEdgeFlows(writer, is_undirected, no_empty_edges)
	}
	writer.Flush()
}

func main() {
	graph_file := flag.String("f", "", "file with graph")
	output_file := flag.String("o", "", "output file")
	is_undirected := flag.Bool("u", false, "make graph undirected")
	edge_flows := flag.Bool("e", false, "print edge flows")
	no_empty_edges := flag.Bool("z", false, "do not print empty edges")
	flag.Parse()

	scanGraph(makeScanner(graph_file), is_undirected)
	printFlow(makeWriter(output_file), dinic(), edge_flows, is_undirected, no_empty_edges)
}
