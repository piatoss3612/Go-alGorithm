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
	TC      int
	N, M, W int
	adjs    [][]Adj // 인접 리스트
	upper   []int   // 최단 거리
)

type Adj struct {
	node   int
	weight int
}

const INF = 987654321

// 난이도: Gold 3
// 메모리: 2360KB
// 시간: 24ms
// 분류: 벨만-포드 알고리즘
// 참고: https://www.acmicpc.net/board/view/72995
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	TC = scanInt()
	for i := 1; i <= TC; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	N, M, W = scanInt(), scanInt(), scanInt()

	adjs = make([][]Adj, N+1)

	// 도로 (양방향)
	for i := 1; i <= M; i++ {
		s, e, t := scanInt(), scanInt(), scanInt()
		adjs[s] = append(adjs[s], Adj{e, t})
		adjs[e] = append(adjs[e], Adj{s, t})
	}

	// 웜홀 (단방향)
	for i := 1; i <= W; i++ {
		s, e, t := scanInt(), scanInt(), scanInt()
		adjs[s] = append(adjs[s], Adj{e, -t})
	}

	// 최단 거리 초기화
	upper = make([]int, N+1)
	for i := 1; i <= N; i++ {
		upper[i] = INF
	}
	upper[1] = 0 // 1번을 시작점으로 설정
}

func Solve() {
	//  한 지점에서 출발을 하여서 시간여행을 하기 시작하여
	// 다시 출발을 하였던 위치로 돌아왔을 때, 출발을 하였을 때보다 시간이 되돌아가 있는 경우 => 음수 사이클 찾기

	// 음수 사이클이 존재하면 YES, 아니면 NO
	if hasNegativeCycle() {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

// 음수 사이클 존재 여부 확인
func hasNegativeCycle() bool {
	// N-1번 완화 시도
	for i := 1; i <= N-1; i++ {
		// 완화 실패 = 음수 사이클이 존재하지 않음
		if !relax() {
			return false
		}
	}

	// N-1번 완화 성공 = 음수 사이클이 존재할 수 있음
	// N번째 완화 시도...
	// N번째 완화 성공 = 음수 사이클이 존재
	return relax()
}

// 완화: 모든 간선에 대해 최단 거리를 갱신
// 여기서는 음수 사이클 어디에 있든 존재하는지 여부를 확인하기 위해
// 시작 정점에서 도달 가능한 정점(upper[v] != INF)인지 아닌지는 고려하지 않음
func relax() bool {
	relaxed := false
	for from := 1; from <= N; from++ {
		for _, adj := range adjs[from] {
			to := adj.node
			weight := adj.weight

			if upper[to] > upper[from]+weight {
				upper[to] = upper[from] + weight
				relaxed = true
			}
		}
	}
	return relaxed
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
