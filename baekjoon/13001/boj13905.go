package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, M    int
	s, e    int
	edges   []edge // 섬과 섬 사이에 연결된 다리들
	parent  []int  // 섬들의 연결 정보
	pepero  []int  // 들고 갈 수 있는 최대 빼빼로
	INF     = 987654321
)

// 다리 정보
type edge struct {
	h1, h2, k int
}

// 메모리: 15400KB
// 시간: 168ms
// 최소 신장 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	N, M = scanInt(), scanInt()
	s, e = scanInt(), scanInt()

	parent = make([]int, N+1)
	pepero = make([]int, N+1)
	for i := 1; i <= N; i++ {
		parent[i] = i
		pepero[i] = INF
	}

	edges = make([]edge, M)
	for i := 0; i < M; i++ {
		edges[i] = edge{scanInt(), scanInt(), scanInt()}
	}

	// 무게제한이 가장 큰 다리부터 내림차순으로 정렬
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].k > edges[j].k
	})

	// 무게제한이 가장 큰 다리부터 탐색
	for _, edge := range edges {
		// 섬 h1과 h2의 연결 정보 x, y 불러오기
		x, y := find(edge.h1), find(edge.h2)
		// x와 y가 아직 연결되어 있지 않다면
		if x != y {
			parent[y] = x                      // x와 y를 연결, 순서는 상관 없다
			pepero[x] = min(pepero[x], edge.k) // x가 포함된 경로로 들고갈 수 있는 최대 빼빼로 수 갱신
		}

		// 시작지점 s와 끝나는 지점 e의 연결 정보 x, y 불러오기
		x, y = find(s), find(e)
		// x와 y가 연결되어 있다면
		if x == y {
			// x가 포함된 경로로 들고갈 수 있는 최대 빼빼로 수를 출력하고 종료
			fmt.Fprintln(writer, pepero[x])
			return
		}
	}

	// s에서 e로 가는 경로를 찾지 못한 경우
	fmt.Fprintln(writer, 0)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
