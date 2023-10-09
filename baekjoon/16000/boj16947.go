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
	subway  [3001][]int // 전철역 연결 정보
	N       int

	// 순환선(사이클) 정보
	checked  [3001]bool // 해당 전철역 탐색 여부
	cycle    [3001]bool // 전철역이 순환선에 포함되는지 여부
	prev     [3001]int  // 탐색 과정의 이전역
	hasCycle = false

	// 순환선까지의 거리 정보
	dist    [3001]int  // 순환선까지의 거리
	visited [3001]bool // 깊이 우선 탐색 방문 여부
)

// 메모리: 1744KB
// 시간: 12ms
// 2호선의 순환선 1개 즉 그래프의 사이클을 1개 찾는 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	for i := 1; i <= N; i++ {
		a, b := scanInt(), scanInt()
		subway[a] = append(subway[a], b)
		subway[b] = append(subway[b], a)
	}

	findCycle(1)
	bfs()

	for i := 1; i <= N; i++ {
		fmt.Fprintf(writer, "%d ", dist[i])
	}
	fmt.Fprintln(writer)
}

func findCycle(x int) {
	checked[x] = true

	for _, next := range subway[x] {
		// 사이클이 이미 발견됬다면 현재 루프 종료
		if hasCycle {
			return
		}

		// x 역과 연결된 다음역이 이미 탐색이 완료되었는데
		// x의 이전역이 아니라면 사이클을 발견한 것
		if checked[next] && next != prev[x] {
			cycle[x] = true
			hasCycle = true

			// 사이클 포함 여부 갱신
			for x != next {
				cycle[prev[x]] = true
				x = prev[x]
			}
			return
		} else if !checked[next] {
			prev[next] = x
			findCycle(next)
		}
	}
}

func bfs() {
	// 순환선에 포함된 지하철역으로부터
	// 포함되지 않은 지하철역까지의 거리 구하기
	q := []int{}
	for i := 1; i <= N; i++ {
		if cycle[i] {
			q = append(q, i)
			visited[i] = true
		}
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]

		for _, next := range subway[x] {
			if !visited[next] {
				visited[next] = true
				dist[next] = dist[x] + 1
				q = append(q, next)
			}
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
