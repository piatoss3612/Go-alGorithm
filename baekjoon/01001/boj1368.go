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
	N         int
	parent    []int  // 연결된 루트 요소
	paths     []Path // 다른 논으로부터 물을 끌어오는 경로
	edgeCount int    // 전체 간선의 개수
	totalCost int    // 모든 논에 물을 대는데 필요한 최소비용
)

type Path struct {
	a, b, c int
}

// 난이도: Gold 2
// 메모리: 6148KB
// 시간: 32ms
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
		// i번째 논에 우물을 직접 파는 경우

		// #1 모든 논에 물을 대기 위해서는 최초에는 반드시 한 개의 우물을 직접 파야 한다: 1개의 경우
		// #2 그 이후에는 우물을 직접 파거나 다른 논으로부터 물을 끌어오는 경우 중 선택을 할 수 있다: N-1개의 경우

		// 우물을 직접파는 경우를 다른 논으로부터 물을 끌어오는 경우로 치환,
		// 물을 끌어오는 가상의 논의 번호를 N의 범위에 포함되지 않는 숫자 0(또는 301 이상의 숫자)으로 설정
		// 0과 i번째 논을 연결하는 비용을 구조체로 묶어 새로운 경로로 paths 슬라이스에 추가

		// 이 경우 최소 스패닝 트리를 구성하기 위해 가상의 0번째 논을 반드시 포함해야 하므로
		// 결과로 구성된 최소 스패닝 트리는 #1 최초에 우물을 직접 파는 경우를 반드시 포함하게 된다
		paths = append(paths, Path{0, i, scanInt()})
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			c := scanInt()
			// 중복된 입력 배제
			if j <= i {
				continue
			}
			// i 또는 j번째 논으로부터 물을 끌어오는 경로 추가
			paths = append(paths, Path{i, j, c})
		}
	}

	// 모든 경로를 비용이 가장 적은 경로를 기준으로 오름차순 정렬
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].c < paths[j].c
	})
}

func Solve() {
	// 0번째 논을 포함하여 N+1개의 논을 연결하여
	// 최소 스패닝 트리를 구성하기 위하여 N개의 간선이 필요
	for edgeCount < N && len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]

		if union(path.a, path.b) {
			totalCost += path.c
			edgeCount++
		}
	}
	fmt.Fprintln(writer, totalCost)
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
