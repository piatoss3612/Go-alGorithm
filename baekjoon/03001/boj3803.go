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
	P, R    int
	Parent  []int   // 각 정점의 부모 요소
	Routes  []Route // 입력된 경로 정보
)

type Route struct {
	A, B int
	Len  int
}

// 난이도: Gold 4
// 메모리: 1524KB
// 시간: 12ms
// 분류: 최소 스패닝 트리
// 직, 간접적으로 모든 P개의 정점을 연결하는 최소 비용 구하기 -> P-1개의 간선으로 구성된 최소 스패닝 트리 찾기
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	TC()
}

func TC() {
	for {
		P = scanInt()
		// 종료 조건
		if P == 0 {
			return
		}

		R = scanInt()

		Parent = make([]int, P+1)
		for i := 1; i <= P; i++ {
			Parent[i] = i
		}

		Routes = make([]Route, R)
		for i := 0; i < R; i++ {
			Routes[i] = Route{scanInt(), scanInt(), scanInt()}
		}

		// 거리가 짧은 경로를 기준으로 모든 경로 정보를 오름차순 정렬
		sort.Slice(Routes, func(i, j int) bool {
			return Routes[i].Len < Routes[j].Len
		})

		Solve()
	}
}

func Solve() {
	edges := 0
	ans := 0

	for len(Routes) > 0 {
		r := Routes[0]
		Routes = Routes[1:]

		pa, pb := Find(r.A), Find(r.B)
		// 정점 A와 정점 B가 아직 연결되어 있지 않은 경우
		if pa != pb {
			Parent[pa] = pb // A와 B를 연결
			ans += r.Len    // A와 B 사이의 거리만큼 전체 거리 증가
			edges++         // 간선의 개수 추가
		}

		// 최소 스패닝 트리를 구성한 경우
		if edges == P-1 {
			break // 탐색 종료
		}
	}

	fmt.Fprintln(writer, ans)
}

func Find(x int) int {
	if Parent[x] == x {
		return x
	}

	Parent[x] = Find(Parent[x])
	return Parent[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
