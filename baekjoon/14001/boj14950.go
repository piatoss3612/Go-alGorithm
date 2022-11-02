package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	N, M, t   int
	parent    []int  // 연결된 루트 요소
	roads     []Road // 도로 정보
	baseCost  int    // 도시를 정복할 때마다 증가하는 기본 비용
	roadCount int    // 도시를 정복하는데 사용한 도로의 개수
	ans       int    // 모든 도시를 정복하는데 필요한 최소 비용
)

type Road struct {
	a, b, c int
}

// 난이도: Gold 3
// 메모리: 2096KB
// 시간: 948ms -> 24ms
// 분류: 최소 스패닝 트리
// 회고:
// 문제를 읽다가 1번 도시부터 시작해야 되고 도로 비용이 누적되므로 여러가지 조건을 붙여봤는데
// 최소 스패닝 트리에서 결국에는 모든 도시가 연결되기 때문에 순서를 따질 필요가 없다는 것을 깨달았다
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, M, t = scanInt(), scanInt(), scanInt()
	parent = make([]int, N+1)
	for i := 1; i <= N; i++ {
		parent[i] = i
	}
	roads = make([]Road, M)
	for i := 0; i < M; i++ {
		roads[i] = Road{scanInt(), scanInt(), scanInt()}
	}

	// 비용 c가 가장 작은 도로를 기준으로 오름차순 정렬
	sort.Slice(roads, func(i, j int) bool {
		return roads[i].c < roads[j].c
	})
}

func Solve() {
	for roadCount < N-1 && len(roads) > 0 {
		r := roads[0]
		roads = roads[1:]

		if union(r.a, r.b) {
			ans += r.c + baseCost
			baseCost += t
			roadCount++
		}
	}

	fmt.Fprintln(writer, ans)
}

// 948ms걸린 솔루션
func SolveOld() {
	for roadCount < N-1 && len(roads) > 0 {
		for idx, r := range roads {
			pa, pb := find(r.a), find(r.b)
			if pa != pb {
				if pa == 1 {
					parent[pb] = 1
					ans += r.c + baseCost
					baseCost += t
					roadCount++
					roads = append(roads[:idx+1], roads[idx+1:]...)
					break
				} else if pb == 1 {
					parent[pa] = 1
					ans += r.c + baseCost
					baseCost += t
					roadCount++
					roads = append(roads[:idx+1], roads[idx+1:]...)
					break
				}
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func find(x int) int {
	if x == parent[x] {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) bool {
	x, y = find(x), find(y)
	if x != y {
		parent[y] = x
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
