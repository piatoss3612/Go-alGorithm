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
	N, M    int
	conns   [201][201]int
)

const INF = 987654321

// 난이도: Gold 4
// 메모리: 1304KB
// 시간: 28ms
// 분류: 그래프 탐색, 플로이드-와샬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			conns[i][j] = INF
		}
	}

	for i := 1; i <= M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		conns[a][b] = c
		conns[b][a] = c
	}
}

func Solve() {
	// 플로이드-와샬 알고리즘을 이용해 모든 정점에서 모든 정점으로의 최단 거리를 구한다
	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if i == j {
					continue
				}
				conns[i][j] = min(conns[i][j], conns[i][k]+conns[k][j])
			}
		}
	}

	minComputer, minDist := 0, INF*N

	// 각 정점에서 모든 정점으로의 거리의 합이 가장 작은 정점을 찾는다
	for i := 1; i <= N; i++ {
		dist := 0
		for j := 1; j <= N; j++ {
			dist += conns[i][j]
		}
		if dist < minDist {
			minComputer = i
			minDist = dist
		}
	}

	fmt.Fprintln(writer, minComputer)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
