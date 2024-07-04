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
	t       int
	n       int
	sx, sy  int      // 상근이와 친구들의 위치
	cu      [][2]int // 편의점의 위치
	px, py  int      // 페스티벌의 위치
)

// 9205번: 맥주 마시면서 걸어가기
// hhttps://www.acmicpc.net/problem/9205
// 난이도: 골드 5
// 메모리: 1012 KB
// 시간: 8 ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	t = scanInt()
	for i := 0; i < t; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	n = scanInt()
	sx, sy = scanInt(), scanInt()
	cu = make([][2]int, n)
	for i := 0; i < n; i++ {
		cu[i] = [2]int{scanInt(), scanInt()}
	}
	px, py = scanInt(), scanInt()
}

func Solve() {
	visited := make([]bool, n) // 방문 여부

	q := [][2]int{{sx, sy}} // 상근이 위치부터 시작

	// 왜 너비 우선 탐색인가?
	// 현 위치에서 최대로 움직일 수 있는 거리는 1000이다.
	// 현 위치에서 갈 수 있는 편의점을 모두 방문했을 때의 방문한 편의점들을 A, B, C라고 하면
	// 이 편의점들 간의 거리가 1000 이하인 경우: 서로 굳이 연결할 필요가 없다
	// 이 편의점들 간의 거리가 1000 초과인 경우: 이 경우도 서로 굳이 연결할 필요가 없다
	// 따라서, 깊이 우선 탐색을 사용할 필요 없이 너비 우선 탐색으로 최단 경로만 찾으면 된다.

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]

		// 페스티벌에 도착했으면 happy 출력 후 종료
		if canMove(x, y, px, py) {
			fmt.Fprintln(writer, "happy")
			return
		}

		// 현 위치에서 갈 수 있는 편의점을 방문
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			if canMove(x, y, cu[i][0], cu[i][1]) {
				q = append(q, cu[i])
				visited[i] = true
			}
		}
	}

	// 페스티벌에 도착하지 못했으면 sad 출력
	fmt.Fprintln(writer, "sad")
}

// (x1, y1)에서 (x2, y2)로 이동할 수 있는지 확인
func canMove(x1, y1, x2, y2 int) bool {
	return abs(x1-x2)+abs(y1-y2) <= 1000
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
