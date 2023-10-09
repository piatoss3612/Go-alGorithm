package bj1240

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   map[int][][]int
	visited []bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()

	graph = make(map[int][][]int)

	for i := 1; i < n; i++ {
		a, b, dist := scanInt(), scanInt(), scanInt()
		graph[a] = append(graph[a], []int{b, dist})
		graph[b] = append(graph[b], []int{a, dist})
	}

	for i := 1; i <= m; i++ {
		start, end := scanInt(), scanInt()
		visited = make([]bool, n+1)
		dfs(start, end, 0)
	}
}

func dfs(current, target, length int) {
	if current == target {
		fmt.Fprintln(writer, length)
		return
	}

	visited[current] = true
	for _, vtx := range graph[current] {
		next := vtx[0]
		dist := vtx[1]
		if !visited[next] {
			visited[next] = true
			dfs(next, target, length+dist)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
