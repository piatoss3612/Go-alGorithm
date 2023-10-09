package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Edge struct {
	node int
	dist int
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   [][]Edge
	visited []bool
	ans     = 0
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	graph = make([][]Edge, n+1)
	visited = make([]bool, n+1)

	var a, b, dist int
	for i := 1; i < n; i++ {
		a, b, dist = scanInt(), scanInt(), scanInt()
		graph[a] = append(graph[a], Edge{b, dist})
		graph[b] = append(graph[b], Edge{a, dist})
	}

	for i := 1; i <= n; i++ {
		visited[i] = true
		DFS(i, 0)
		visited[i] = false
	}

	fmt.Fprintln(writer, ans)
}

func DFS(v, sum int) {
	if sum > ans {
		ans = sum
	}

	for _, edge := range graph[v] {
		if !visited[edge.node] {
			visited[edge.node] = true
			DFS(edge.node, sum+edge.dist)
			visited[edge.node] = false
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
