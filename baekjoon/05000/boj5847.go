package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner     = bufio.NewScanner(os.Stdin)
	writer      = bufio.NewWriter(os.Stdout)
	N, M        int
	time        []int   // i번째 소의 우유를 짜는데 걸리는 시간
	inDegree    []int   // 진입 차수
	constraints [][]int // constraints[i]: i번째 소의 우유를 짜고 다음 차례의 소들의 번호
	dp          []int   // i번째 소의 우유를 짜는데까지 걸린 시간의 최댓값 메모이제이션
	ans         int
)

// 난이도: Gold 3
// 메모리: 3380KB
// 시간: 20ms
// 분류: 다이나믹 프로그래밍, 위상 정렬
// 회고: 문제가 영어로 작성되어 있어 약간의 거부감이 들지만 지극히 평범한 위상 정렬 문제.
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	time = make([]int, N+1)
	inDegree = make([]int, N+1)
	constraints = make([][]int, N+1)
	dp = make([]int, N+1)

	for i := 1; i <= N; i++ {
		time[i] = scanInt()
	}

	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		inDegree[b]++
		constraints[a] = append(constraints[a], b)
	}
}

func Solve() {
	TopologicalSort()
	fmt.Fprintln(writer, ans)
}

func TopologicalSort() {
	q := []int{}
	for i := 1; i <= N; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
			dp[i] = time[i]
		}
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]

		for _, next := range constraints[x] {
			dp[next] = max(dp[next], dp[x]+time[next])
			ans = max(dp[next], ans)
			inDegree[next]--
			if inDegree[next] == 0 {
				q = append(q, next)
			}
		}
	}
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
