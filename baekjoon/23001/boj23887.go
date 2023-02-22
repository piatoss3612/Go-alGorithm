package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, M, K, S int
	class      [501][501]int
	visited    [501][501]bool
	children   [250001][]int // 1 <= K <= N*M (500 * 500)
	dp         [250001]int
	dx         = []int{-1, -1, -1, +0, +1, +1, +1, +0}
	dy         = []int{-1, +0, +1, +1, +1, +0, -1, -1}
)

// 난이도: Gold 3
// 메모리: 36660KB
// 시간: 240ms
// 분류: 다이나믹 프로그래밍, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()
	for i := 1; i <= K; i++ {
		x, y := scanInt(), scanInt()
		class[x][y] = i
	}
	S = scanInt()
}

func Solve() {
	var sx, sy int
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if class[i][j] == S {
				sx, sy = i, j
				break
			}
		}
	}

	BFS(sx, sy) // S번 학생의 위치에서 너비 우선 탐색 실행

	// 프린트를 받지 못한 학생이 있는 경우 찾기
	for i := 1; i <= K; i++ {
		// dp[i]가 1이 아니라면 프린트를 받지 못한 것이다
		if dp[i] < 1 {
			fmt.Fprintln(writer, -1)
			return
		}
	}

	rec(S) // 각 학생이 받아야 하는 프린트의 개수를 구하기 위한 재귀 함수 실행

	for i := 1; i <= K; i++ {
		fmt.Fprintf(writer, "%d ", dp[i])
	}
	fmt.Fprintln(writer)
}

type Student struct {
	x, y   int
	number int
}

func BFS(x, y int) {
	var q []Student

	// S번 학생에게 프린트 전달
	visited[x][y] = true
	q = append(q, Student{x, y, class[x][y]})
	dp[class[x][y]] = 1

	for len(q) > 0 {
		// 문제 조건: 어떤 학생이 두 명 이상의 학생에게 동시에 프린트를 받을 수 있다면 항상 번호가 가장 작은 학생에게만 받는다
		// 번호가 작은 학생 위치에서부터 깊이 우선 탐색을 실행하기 위해 학생 번호를 기준으로 오름차순 정렬
		sort.Slice(q, func(i, j int) bool {
			return q[i].number < q[j].number
		})

		candidates := []Student{} // 큐에 들어 있는 모든 학생들이 동시에 탐색을 실행하므로 임시로 탐색 결과를 담을 슬라이스를 초기화

		for len(q) > 0 {
			h := q[0]
			q = q[1:]

			for i := 0; i < 8; i++ {
				nx, ny := h.x+dx[i], h.y+dy[i]
				if isInBound(nx, ny) && !visited[nx][ny] && class[nx][ny] != 0 {
					visited[nx][ny] = true
					child := class[nx][ny]
					candidates = append(candidates, Student{nx, ny, child})
					dp[child] = 1
					children[h.number] = append(children[h.number], child) // h.number는 child의 유일한 부모
				}
			}
		}

		q = append(q, candidates...) // 한 차례 탐색을 실행한 뒤에 구한 탐색 결과들을 다시 큐에 집어넣는다
	}
}

func rec(x int) {
	// x는 child의 유일한 부모이므로 방문 처리 불필요
	for _, child := range children[x] {
		rec(child)
		dp[x] += dp[child] // child에서 다른 모든 자식 노드들로 프린트를 전달하기 위해 필요한 프린트의 개수 누적
	}
}

func isInBound(x, y int) bool {
	return x >= 1 && x <= N && y >= 1 && y <= M
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
