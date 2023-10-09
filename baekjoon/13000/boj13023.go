package bj13023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   [][]int
	visited []bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	graph = make([][]int, n)
	for i := 1; i <= m; i++ {
		a, b := scanInt(), scanInt()
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited = make([]bool, n)

	for i := 0; i < n; i++ {
		dfs(i, 0)
	}

	fmt.Fprintln(writer, 0)
}

func dfs(x, d int) {
	visited[x] = true
	if d == 4 {
		fmt.Fprintln(writer, 1)
		writer.Flush()
		os.Exit(0)
	}

	for _, v := range graph[x] {
		if !visited[v] {
			dfs(v, d+1)
		}
	}
	visited[x] = false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
