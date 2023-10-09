package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, M, S, E int
	path       [10001][]Path // path[i]: i(1<=i<=N)번 섬과 연결된 섬의 번호와 중량 제한 정보들
)

type Path struct {
	conn, limit int
}

// 난이도: Gold 3
// 메모리: 12196KB
// 시간: 80ms
// 분류: 이분 탐색, 너비 우선 탐색
// 메모리 제한 128MB에 주의
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()

	var a, b, c int
	for i := 1; i <= M; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()
		path[a] = append(path[a], Path{b, c})
		path[b] = append(path[b], Path{a, c})
	}

	S, E = scanInt(), scanInt()
}

func Solve() {
	fmt.Fprintln(writer, BinarySearch())
}

// 이분 탐색을 통해 S에서 E로 한 번에 옮길 수 있는 중량의 최댓값(upper bound)을 구한다
func BinarySearch() int {
	l, r := 1, 1000000000
	for l <= r {
		m := (l + r) / 2 // 예상되는 중량
		// m만큼의 무게를 S에서 E로 한 번에 옮길 수 있는 경우
		if BFS(m) {
			l = m + 1 // 무게를 늘려서 시도
		} else {
			r = m - 1 // 무게를 줄여서 시도
		}
	}

	return r
}

// 너비 우선 탐색을 통해 S에서 E로 expected만큼의 중량을 한 번에 옮길 수 있는지 확인한다
func BFS(expected int) bool {
	var visited [10001]bool
	q := []int{S} // S에서 시작
	visited[S] = true

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		// E로 expected만큼의 중량을 한 번에 옮긴 경우: 탐색 종료
		if front == E {
			return true
		}

		for _, next := range path[front] {
			// 연결되어 있는 섬을 아직 방문하지 않았고 다리의 중량 제한이 expected보다 크거나 같은 경우
			if !visited[next.conn] && next.limit >= expected {
				visited[next.conn] = true
				q = append(q, next.conn)
			}
		}
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
