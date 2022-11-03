package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner        = bufio.NewScanner(os.Stdin)
	writer         = bufio.NewWriter(os.Stdout)
	N              int
	parent         []int
	candidates     []*Road // 새롭게 설치될 수 있는 철로들의 정보
	newRoads       []*Road // 새롭게 설치된 철로
	newRoadCount   int     // 새롭게 설치된 철로의 수
	totalRoadCount int     // 중복된 연결을 제외하고 모든 도시를 연결하는 철로의 수
	totalCost      int     // 모든 도시를 연결하기 위해 필요한 철로 설치 비용의 최솟값
)

type Road struct {
	a, b, c int
}

// 난이도: Gold 3
// 메모리: 2308KB
// 시간: 16ms
// 분류: 최소 스패닝 트리
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	parent = make([]int, N+1)

	for i := 1; i <= N; i++ {
		parent[i] = i
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			cost := scanInt()
			// i번 도시 자기자신으로 향하는 철로와
			// 양방향 이동이 가능한 철로의 중복되는 입력은 고려하지 않는다
			if j <= i {
				continue
			}

			if cost < 0 {
				// 이미 설치되어 있는 철로인 경우
				totalCost -= cost
				pi, pj := find(i), find(j)
				if pi != pj {
					// union 작업
					parent[pj] = pi
					totalRoadCount++
				}
			} else {
				// i와 j를 연결하는 설치 비용이 cost인 철로를 후보로 추가
				candidates = append(candidates, &Road{i, j, cost})
			}
		}
	}

	// 모든 도시를 연결하기 위해 필요한 철로 설치 비용의 최솟값을 구하기 위해
	// 가장 적은 설치 비용이 드는 철로를 기준으로 오름차순 정렬
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].c < candidates[j].c
	})
}

func Solve() {
	for totalRoadCount < N-1 && len(candidates) > 0 {
		r := candidates[0]
		candidates = candidates[1:]

		pa, pb := find(r.a), find(r.b)
		if pa != pb {
			// union 작업
			parent[pb] = pa
			totalCost += r.c
			newRoadCount++
			newRoads = append(newRoads, r)
			totalRoadCount++
		}
	}

	fmt.Fprintln(writer, totalCost, newRoadCount)
	for _, r := range newRoads {
		fmt.Fprintln(writer, r.a, r.b)
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
