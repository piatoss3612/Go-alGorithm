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
	INF     = 987654321
)

// 메모리: 2996KB
// 시간: 164ms
// DAG(Directed Acyclic Graph) 문제
// 즉, 플로이드 와샬 알고리즘을 수행하고
// a에서 b로 가는 경로가 있다면 a가 먼저 발생한 사건이라는 것을 알 수 있다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	graph := make([][]int, n+1)

	// 그래프 초기화
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i != j {
				graph[i][j] = INF
			}
		}
	}

	var a, b int

	// a가 먼저 일어난 사건, b가 a보다 나중에 일어난 사건
	for i := 1; i <= m; i++ {
		a, b = scanInt(), scanInt()
		graph[a][b] = 1
	}

	// 플로이드 와샬 알고리즘
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	s := scanInt()
	for i := 1; i <= s; i++ {
		a, b = scanInt(), scanInt()

		fmt.Fprintln(writer, graph[a][b], graph[b][a])

		if graph[a][b] < graph[b][a] {
			fmt.Fprintln(writer, -1)
		} else if graph[a][b] > graph[b][a] {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
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
