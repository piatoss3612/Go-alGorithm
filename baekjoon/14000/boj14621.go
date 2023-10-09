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
	N, M      int
	univ      []byte // 대학 구성원 정보 (W: 여초, M: 남초)
	parent    []int  // 최소 스패닝 트리 루트 요소
	paths     []Path // 도로 정보
	pathCount int    // 조건을 만족하는 도로의 수
	ans       int    // 최소 거리
)

type Path struct {
	u, v     int
	distance int
}

// 난이도: Gold 3
// 메모리: 1248KB
// 시간: 8ms
// 분류: 최소 스패닝 트리
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	univ = make([]byte, N+1)
	parent = make([]int, N+1)
	for i := 1; i <= N; i++ {
		univ[i] = scanByte()
		parent[i] = i
	}

	paths = make([]Path, M)
	for i := 0; i < M; i++ {
		paths[i] = Path{scanInt(), scanInt(), scanInt()}
	}

	// 거리가 최소인 도로를 기준으로 오름차순 정렬
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].distance < paths[j].distance
	})
}

func Solve() {
	for len(paths) > 0 {
		p := paths[0]
		paths = paths[1:]

		// 남초 대학교와 여초 대학교를 연결하는 도로인 경우
		if univ[p.u] != univ[p.v] {
			// 이미 선택한 도로와 부분 경로가 중복되지 않는 경우
			if union(p.u, p.v) {
				// 도로 p를 최소 스패닝 트리에 포함한다
				pathCount++
				ans += p.distance
			}
		}
	}

	// 모든 학교를 연결할 수 있는 경우 = 최소 스패닝 트리를 구성하는 N-1개의 도로를 선택할 수 있는 경우
	if pathCount == N-1 {
		fmt.Fprintln(writer, ans)
	} else {
		fmt.Fprintln(writer, -1)
	}
}

func find(x int) int {
	if parent[x] == x {
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

func scanByte() byte {
	scanner.Scan()
	return scanner.Bytes()[0]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
