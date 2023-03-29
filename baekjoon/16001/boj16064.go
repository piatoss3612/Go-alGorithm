package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N, M     int
	slopes   [][]Slope // slopes[i]: i번째 포인트와 경사로 연결된 포인트들
	inDegree []int     // 진입차수
	dp       []int     // dp[i]: i번째 포인트까지 오는데 얼마나 어썸한지 (최대값)
)

type Slope struct {
	to, condition int
}

// 난이도: Gold 3
// 메모리: 1284KB
// 시간: 8ms
// 분류: 위상정렬, 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	slopes = make([][]Slope, N+1)
	inDegree = make([]int, N+1)
	dp = make([]int, N+1)

	for i := 1; i <= M; i++ {
		s, t, c := scanInt(), scanInt(), scanInt()
		inDegree[t] += 1                           // t로의 진입차수 증가
		slopes[s] = append(slopes[s], Slope{t, c}) // s -> t 단방향
	}
}

func Solve() {
	q := make([]int, 0, N+1)

	// 진입차수가 0인 포인트들 큐에 삽입
	for i := 1; i <= N; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	// 위상정렬
	for len(q) > 0 {
		from := q[0]
		q = q[1:]

		// from -> to 경사로를 타고 이동
		for _, s := range slopes[from] {
			to := s.to
			condition := s.condition

			dp[to] = max(dp[to], dp[from]+condition) // dp 갱신
			inDegree[to] -= 1                        // to로의 진입차수 감소

			// to로의 진입차수가 0이 되면 큐에 삽입
			if inDegree[to] == 0 {
				q = append(q, to)
			}
		}
	}

	// dp 배열에서 최대값 찾기
	ans := 0
	for i := 1; i <= N; i++ {
		ans = max(ans, dp[i])
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
