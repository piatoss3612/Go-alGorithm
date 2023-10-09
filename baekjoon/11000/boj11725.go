package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	graph := make([][]int, n+1)

	for i := 1; i < n; i++ {
		a, b := scanInt(), scanInt()
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, n+1)
	ans := make([]int, n+1)

	queue := []int{1}
	visited[1] = true

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		for _, v := range graph[front] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
				ans[v] = front
			}
		}
	}

	for i := 2; i <= n; i++ {
		fmt.Fprintln(writer, ans[i])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
