# Dinic's algorithm

Dinic's algorithm is a strongly polynomial algorithm for computing the maximum flow in a flow network, conceived in 1970 by Israeli (formerly Soviet) computer scientist Yefim (Chaim) A. Dinitz.

## Description

This app implements algorithm for computing the maximum flow in a flow network with the assumption that maximum edge capacity is less than $2^{32}$. The app combines the Dinic's algorithm with the flow scaling technique and finds max flow in $\mathcal{O}(VE\log C)$ time, where $V$ is the number of vertices, $E$ is the number of edges in the network and $C$ is the maximum edge capacity.

The app is written in Go and supports various modes of work with flow networks.

## Flags

The app supports the following flags:

- `-f` - the path to the file with the flow network description
- `-o` - the path to the file where the result will be written
- `-u` - the flag for undirected graph
- `-e` - the flag for printing flow of each edge to the output
- `-z` - the flag for disabling zero flow edges in the output

## Input format

The input format is the following:

- The first line contains two integers $n$ and $m$ - the number of vertices and the number of edges in the network, respectively.
- The next $m$ lines contain three integers $u$, $v$ and $c$ - the start vertex, the end vertex and the capacity of the edge, respectively.

## Examples
- Normal flow network:
```bash
➜ cat test1.txt
5 7
1 2 2
2 5 5
1 3 6
3 4 2
4 5 1
3 2 3
2 4 1
➜ ./dinic -f test1.txt -e 
Max flow: 6
1 | 1 -> 2 | flow = 2 | capacity = 2
2 | 2 -> 5 | flow = 5 | capacity = 5
3 | 1 -> 3 | flow = 4 | capacity = 6
4 | 3 -> 4 | flow = 1 | capacity = 2
5 | 4 -> 5 | flow = 1 | capacity = 1
6 | 3 -> 2 | flow = 3 | capacity = 3
7 | 2 -> 4 | flow = 0 | capacity = 1
```
- Undirected flow network:
```bash
➜ cat test2.txt
5 6
2 1 2
1 3 1
2 4 2
5 4 4
3 5 2
4 3 1
➜ ./dinic -f test2.txt -e
Max flow: 1
1 | 2 -> 1 | flow = 0 | capacity = 2
2 | 1 -> 3 | flow = 1 | capacity = 1
3 | 2 -> 4 | flow = 0 | capacity = 2
4 | 5 -> 4 | flow = 0 | capacity = 4
5 | 3 -> 5 | flow = 1 | capacity = 2
6 | 4 -> 3 | flow = 0 | capacity = 1
➜ ./dinic -f test2.txt -e -u
Max flow: 3
1 | 2 <- 1 | flow = 2 | capacity = 2
2 | 1 -> 3 | flow = 1 | capacity = 1
3 | 2 -> 4 | flow = 2 | capacity = 2
4 | 5 <- 4 | flow = 2 | capacity = 4
5 | 3 -> 5 | flow = 1 | capacity = 2
6 | 4 -> 3 | flow = 0 | capacity = 1
```
- Flow with many zero edges:
```bash
➜ cat test3.txt
6 7
1 2 1
1 3 1
1 4 1
5 6 1
2 5 1
3 5 1
4 5 1
➜ ./dinic -f test3.txt -e
Max flow: 1
1 | 1 -> 2 | flow = 1 | capacity = 1
2 | 1 -> 3 | flow = 0 | capacity = 1
3 | 1 -> 4 | flow = 0 | capacity = 1
4 | 5 -> 6 | flow = 1 | capacity = 1
5 | 2 -> 5 | flow = 1 | capacity = 1
6 | 3 -> 5 | flow = 0 | capacity = 1
7 | 4 -> 5 | flow = 0 | capacity = 1
➜ ./dinic -f test3.txt -e -z
Max flow: 1
1 | 1 -> 2 | flow = 1 | capacity = 1
4 | 5 -> 6 | flow = 1 | capacity = 1
5 | 2 -> 5 | flow = 1 | capacity = 1
```
  
## Testing
The algorithm passed all tests in the maximum flow problem on [Codeforces](https://codeforces.com/gym/100140).