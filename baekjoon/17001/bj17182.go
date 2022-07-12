package main

import (
	"bufio"
	"fmt"
	_ "math"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	planet     [10][10]int // 행성 간의 이동 시간
	visitedAll int         // 모든 행성을 방문한 경우의 비트마스크
	taken      = 987654321 // 모든 행성을 방문하는데 걸리는 시간의 최솟값
	N, K       int
)

// 메모리: 920KB
// 시간: 28ms
// 비트마스크, 플로이드 와샬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()

	visitedAll = (1 << N) - 1
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			planet[i][j] = scanInt()
		}
		planet[i][i] = 987654321
	}

	// '모든 행성을 방문해야 하는데 이미 방문한 행성도 중복해서 갈 수 있다'라는 조건
	// 경로의 최솟값을 찾는데 경로가 중복될 수 있으므로 플로이드 와샬 알고리즘을 사용할 수 있다

	// 행성 i에서 j로 가는 모든 최단경로(중복된 경로 포함)를 찾고
	// 각 행성을 한 번씩만 방문하는데 걸리는 시간의 최솟값을 찾음으로써 문제를 해결할 수 있다
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				planet[i][j] = min(planet[i][j], planet[i][k]+planet[k][j])
			}
		}
	}

	// K번 행성부터 시작
	rec(1<<K, K, 0)

	fmt.Fprintln(writer, taken)
}

// visited: 지금까지 방문한 행성의 비트마스크
// current: 현재 방문중인 행성
// time: 행성을 방문하는데 걸린 누적 시간
func rec(visited, current, time int) {
	// 모든 행성을 방문한 경우
	if visited&visitedAll == visitedAll {
		taken = min(taken, time)
		return
	}

	for i := 0; i < N; i++ {
		// i번째 행성을 아직 방문하지 않은 경우
		if visited&(1<<i) != (1 << i) {
			// i번째 행성 방문처리, 현재 행성에서 i행성으로 이동하는 시간 추가
			rec(visited|(1<<i), i, time+planet[current][i])
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
