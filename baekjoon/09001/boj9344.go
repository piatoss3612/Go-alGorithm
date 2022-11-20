package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner       = bufio.NewScanner(os.Stdin)
	writer        = bufio.NewWriter(os.Stdout)
	T, N, M, p, q int
	parent        []int  // 연결된 부모 토시
	roads         []Road // 지을 수 있는 도로들
)

type Road struct {
	u, v, w int
}

// 난이도: Gold 3
// 메모리: 1784KB
// 시간: 12ms
// 분류: 최소 스패닝 트리
// 양방향 도로로 한 도시에서 다른 모든 도시로 이동할 수 있는 가장 짧은 도로망 만들기
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	// T만큼 테스트 케이스 실행
	for i := 1; i <= T; i++ {
		Input()
		Solve()
	}
}

func Input() {
	N, M, p, q = scanInt(), scanInt(), scanInt(), scanInt()
	// 부모 도시를 자기자신으로 초기화
	parent = make([]int, N+1)
	for i := 1; i <= N; i++ {
		parent[i] = i
	}

	// 도로 정보 입력
	roads = make([]Road, M)
	for i := 0; i < M; i++ {
		roads[i] = Road{scanInt(), scanInt(), scanInt()}
	}

	// 비용이 적게 드는(=가장 짧은) 도로를 기준으로 오름차순 정렬
	sort.Slice(roads, func(i, j int) bool {
		return roads[i].w < roads[j].w
	})
}

func Solve() {
	count := 0 // 최소 스패닝 트리를 구성하기 위해 N-1개의 도로를 선택했는지 판별하는 변수
	ans := false
	for len(roads) > 0 && count < N-1 {
		road := roads[0]
		roads = roads[1:]

		pu, pv := find(road.u), find(road.v)
		if pu != pv {
			// 현재 선택된 도로가 연결하는 두 도시의 부모 도시가 아직 연결되어 있지 않다면
			// 유니온 연산 실행 = 현재 도로를 짓기로 결정
			parent[pv] = pu
			// p와 q를 연결하는 도로인지 판별
			if (road.u == p && road.v == q) || (road.u == q && road.v == p) {
				ans = true
			}
			count++
		}
	}

	// 가장 짧은 도로망을 구성하면서 p와 q를 직접 연결하는 도로를 포함한 경우
	if ans {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func find(x int) int {
	if x == parent[x] {
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
