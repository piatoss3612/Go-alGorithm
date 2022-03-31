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
	INF     = 10000000
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	graph := make([][]int, n+1)
	graph[0] = make([]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if i == j {
				graph[i][j] = 0
			} else {
				graph[i][j] = INF
			}
		}
	}

	for i := 1; i <= m; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		graph[a][b] = getMin(graph[a][b], c)
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if i == j {
					continue
				}
				graph[i][j] = getMin(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if graph[i][j] == INF {
				fmt.Fprintf(writer, "%d ", 0)
			} else {
				fmt.Fprintf(writer, "%d ", graph[i][j])
			}
		}
		fmt.Fprintln(writer)
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
