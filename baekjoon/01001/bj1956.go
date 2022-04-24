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
	INF     = 100000000 // INF 값이 충분히 크지 않으면 틀리는 문제
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	v, e := scanInt(), scanInt()
	graph := make([][]int, v+1)
	graph[0] = make([]int, v+1)

	// 사이클을 가지는 그래프이므로 i==j인 경우에도 INF를 할당했다
	for i := 1; i <= v; i++ {
		graph[i] = make([]int, v+1)
		for j := 1; j <= v; j++ {
			graph[i][j] = INF
		}
	}

	for i := 1; i <= e; i++ {
		graph[scanInt()][scanInt()] = scanInt()
	}

	// 플로이드 와샬 알고리즘
	for k := 1; k <= v; k++ {
		for i := 1; i <= v; i++ {
			for j := 1; j <= v; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	// 사이클의 최솟값이므로 graph[i][i]의 최솟값을 찾는다
	ans := INF
	for i := 1; i <= v; i++ {
		ans = min(ans, graph[i][i])
	}

	if ans == INF {
		fmt.Fprintln(writer, -1)
		return
	}
	fmt.Fprintln(writer, ans)
}

func min(a, b int) int {
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
