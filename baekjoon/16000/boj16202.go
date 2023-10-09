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
	graph   []int
	edges   []edge
	N, M, K int
)

type edge struct {
	a, b int
}

// 메모리: 1180KB
// 시간: 24ms
// 최소 신장 트리 성립 여부에 대한 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M, K = scanInt(), scanInt(), scanInt()
	graph = make([]int, N+1)
	edges = make([]edge, M+1)

	// 각 간선의 가중치는 입력 순서와 같으므로 궂이 저장하지 않아도 된다
	for i := 1; i <= M; i++ {
		edges[i] = edge{
			scanInt(), scanInt(),
		}
	}

	start := 1 // 간선 탐색 시작 위치

	for i := 1; i <= K; i++ {
		// 그래프의 부모 요소 초기화
		for j := 1; j <= N; j++ {
			graph[j] = j
		}

		numOfEdges := 0 // 최소 신장 트리를 구성하는 간선의 수
		totalCost := 0  // 최소 신장 트리의 총 비용

		for j := start; j <= M; j++ {
			// 최소 신장 트리를 구성하기 위해 필요한 간선의 갯수가 갖추어지면 빠르게 탈출
			if numOfEdges == N-1 {
				break
			}

			// j번째 간선을 구성하는 정점들의 부모 요소 탐색
			pa, pb := find(edges[j].a), find(edges[j].b)
			// 아직 연결되어 있지 않다면 union 연산을 해준다
			if pa != pb {
				graph[pb] = pa
				numOfEdges++
				totalCost += j
			}
		}

		// 최소 신장 트리가 구성된 경우
		if numOfEdges == N-1 {
			fmt.Fprintf(writer, "%d ", totalCost)
		} else {
			// 최소 신장 트리를 구성할 수 없는 경우는
			// 이후 시도에서도 모든 점수는 0이 된다
			for j := start; j <= K; j++ {
				fmt.Fprintf(writer, "%d ", 0)
			}
			fmt.Fprintln(writer)
			return
		}

		start++
	}
}

func find(x int) int {
	if graph[x] == x {
		return x
	}
	graph[x] = find(graph[x])
	return graph[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
