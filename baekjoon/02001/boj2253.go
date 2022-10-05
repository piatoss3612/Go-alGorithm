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
	N, M    int
	small   [10001]bool     // 작은 돌인지 여부
	dp      [10001][142]int // dp[i][j]: i번 돌에서 j만큼 점프할 수 있는 경우, N번 돌까지 점프 횟수의 최솟값

	// 점프하는 거리 K는 계속해서 가속 점프를 하는 경우 j(j+1)/2 >= N (2<=N<=10,000)을 만족하는 최솟값 141+1로 설정
)

const INF = 987654321

// 메모리: 79932KB -> 24100KB
// 시간: 72ms -> 48ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	// 1. 작은 돌 번호 입력
	for i := 1; i <= M; i++ {
		small[scanInt()] = true
	}

	// 2. 재귀 함수 호출, 1번 돌에서 시작하여 N번 돌까지 점프해 갈 때 필요한 점프 횟수의 최솟값
	ans := solve(1, 0)

	// 3. N번 돌까지 점프해갈 수 없는 경우 / 가능한 경우
	if ans == INF {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func solve(rock, jump int) int {
	// 기저 사례: N번째 돌까지 점프한 경우, 더 이상 이동할 수 없으므로 0반환
	if rock == N {
		return 0
	}

	ret := &dp[rock][jump]
	if *ret != 0 {
		return *ret
	}

	*ret = INF // 최솟값 비교를 위해 INF로 초기화

	// -1: 감속, 0: 유지, +1: 가속
	for i := -1; i <= 1; i++ {
		// 조건: 점프는 반드시 한 칸 이상해야 한다
		if jump+i >= 1 {
			// next: 이동할 돌 번호, next가 N 이하이면서 작은 돌이 아닌 경우에만 이동
			if next := rock + jump + i; next <= N && !small[next] {
				*ret = min(*ret, solve(next, jump+i)+1)
			}
		}
	}

	return *ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
