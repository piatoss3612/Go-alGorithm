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
	INF     = 100000000
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	graph := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i != j {
				graph[i][j] = INF
			}
		}
	}

	for i := 1; i <= m; i++ {
		graph[scanInt()][scanInt()] = 1
	}

	// 플로이드 와샬 알고리즘
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				graph[i][j] = min(graph[i][k]+graph[k][j], graph[i][j])
			}
		}
	}

	// a -> b 경로 또는 b -> a 경로가 있다면 cnt 감소
	for i := 1; i <= n; i++ {
		cnt := n
		for j := 1; j <= n; j++ {
			if graph[i][j] != INF || graph[j][i] != INF {
				cnt -= 1
			}
		}
		fmt.Fprintln(writer, cnt)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
