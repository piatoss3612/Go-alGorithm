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
	N, M          int
	optimalParent []int   // 최적의 경로를 선택했을 때의 연결 정보
	worstParent   []int   // 최악의 경로를 선택했을 때의 연결 정보
	optimalRoads  []*Road // 최적의 경로 탐색을 위해 도로를 내리막길을 우선으로 정렬
	worstRoads    []*Road // 최악의 경로 탐색을 위해 도롤를 오르막길을 우선으로 정렬
	optimalCount  int     // 최적의 경로에 포함된 오르막길의 개수
	worstCount    int     // 최악의 경로에 포함된 오르막길의 개수
)

type Road struct {
	a, b, c int // a번 건물과 b건물이 연결되어 있고 형태가 c인 도로
}

// 난이도: Gold 3
// 메모리: 31492KB
// 시간: 204ms
// 분류: 최소 스패닝 트리
// 입구에서 모든 건물로 갈 수 있음이 보장되어 있다
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	optimalParent = make([]int, N+1)
	worstParent = make([]int, N+1)
	for i := 1; i <= N; i++ {
		optimalParent[i] = i
		worstParent[i] = i
	}

	optimalRoads = make([]*Road, 0, M+1)
	worstRoads = make([]*Road, 0, M+1)
	for i := 0; i <= M; i++ {
		r := Road{scanInt(), scanInt(), scanInt()}
		optimalRoads = append(optimalRoads, &r)
		worstRoads = append(worstRoads, &r)
	}

	sort.Slice(optimalRoads, func(i, j int) bool {
		return optimalRoads[i].c > optimalRoads[j].c
	})

	sort.Slice(worstRoads, func(i, j int) bool {
		return worstRoads[i].c < worstRoads[j].c
	})
}

func Solve() {
	for len(optimalRoads) > 0 {
		r := optimalRoads[0]
		optimalRoads = optimalRoads[1:]

		pa, pb := findOptimal(r.a), findOptimal(r.b)
		if pa != pb {
			optimalParent[pb] = pa
			if r.c == 0 {
				optimalCount++
			}
		}
	}

	for len(worstRoads) > 0 {
		r := worstRoads[0]
		worstRoads = worstRoads[1:]

		pa, pb := findWorst(r.a), findWorst(r.b)
		if pa != pb {
			worstParent[pb] = pa
			if r.c == 0 {
				worstCount++
			}
		}
	}

	fmt.Fprintln(writer, worstCount*worstCount-optimalCount*optimalCount)
}

func findOptimal(x int) int {
	if x == optimalParent[x] {
		return x
	}

	optimalParent[x] = findOptimal(optimalParent[x])
	return optimalParent[x]
}

func findWorst(x int) int {
	if x == worstParent[x] {
		return x
	}

	worstParent[x] = findWorst(worstParent[x])
	return worstParent[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
