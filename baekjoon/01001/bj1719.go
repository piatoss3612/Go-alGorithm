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

	// 최소 이동 시간 그래프, 먼저 거쳐야 하는 집하장 그래프 초기화
	graph := make([][]int, n+1)
	ans := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, n+1)
		ans[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			graph[i][j] = INF
		}
		graph[i][i] = 0
	}

	// 최소 이동 시간 그래프는 방향이 없는 그래프
	// 따라서 a->b, b->a 모두 가중치 c를 가진다

	// a -> b로 가기위해 먼저 거쳐야 하는 집하장은 b
	// b -> a로 가기위해 먼저 거쳐야 하는 집하장은 a
	var a, b, c int
	for i := 1; i <= m; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()
		graph[a][b] = c
		graph[b][a] = c
		ans[a][b] = b
		ans[b][a] = a
	}

	// 플로이드 와샬
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				// 최솟값을 찾은 경우
				if graph[i][k]+graph[k][j] < graph[i][j] {
					graph[i][j] = graph[i][k] + graph[k][j]
					// i -> j로 가기 위해 먼저 거쳐야 하는 집하장은
					// i -> k로 가기 위해 먼저 거쳐야 하는 집하장
					ans[i][j] = ans[i][k]
				}
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				fmt.Fprint(writer, "- ")
			} else {
				fmt.Fprintf(writer, "%d ", ans[i][j])
			}
		}
		fmt.Fprintln(writer)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
