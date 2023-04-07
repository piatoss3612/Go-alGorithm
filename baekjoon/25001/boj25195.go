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

	N, M, S int
	edges   [100001][]int // 간선 정보
	visited [100001]bool  // 방문 여부
)

// 난이도: Gold 4
// 메모리: 9404KB
// 시간: 52ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= M; i++ {
		u, v := scanInt(), scanInt()
		edges[u] = append(edges[u], v) // u -> v 단방향 간선
	}

	S = scanInt()
	for i := 1; i <= S; i++ {
		visited[scanInt()] = true // 팬클럽 곰곰이가 있는 정점을 방문처리 함으로써 불필요한 탐색을 줄임
	}
}

func Solve() {
	// 시작 정점에 팬클럽 곰곰이가 있으면 항상 팬클럽과 만나게 되므로 바로 'Yes' 종료
	if visited[1] {
		fmt.Fprintln(writer, "Yes")
		return
	}

	q := []int{}
	// 시작 정점 1을 큐에 넣고 방문처리
	q = append(q, 1)
	visited[1] = true

	// BFS
	for len(q) > 0 {
		x := q[0]
		q = q[1:]

		// 현재 정점 x와 연결되어 있는 정점이 없는 경우
		// 1-[중간 경로]-x까지의 경로에 팬클럽 곰곰이가 없음을 알 수 있다
		// 즉, 팬클럽 곰곰이를 만나지 않고 이동하는 방법을 찾은 것이므로 'yes' 종료
		if len(edges[x]) == 0 {
			fmt.Fprintln(writer, "yes")
			return
		}

		for _, next := range edges[x] {
			if !visited[next] {
				visited[next] = true
				q = append(q, next)
			}
		}
	}

	fmt.Fprintln(writer, "Yes") // 어떻게 가도 팬클럽 곰곰이를 만나게 된다
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
