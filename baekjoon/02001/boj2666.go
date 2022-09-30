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
	n       int
	a, b    int
	m       int
	order   [21]int         // 사용할 벽장의 순서
	dp      [21][21][21]int // [i][a][b]: i번째로 사용할 벽장에 대해서 a, b 벽장문이 열려있을 때 벽장문을 이동하는 횟수의 최솟값
)

const INF = 987654321

// 메모리: 1052KB
// 시간: 4ms
// 다이나믹 프로그래밍, 브루트포스
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	a, b = scanInt(), scanInt()
	m = scanInt()
	for i := 1; i <= m; i++ {
		order[i] = scanInt()
	}

	ans := solve(1, a, b)
	fmt.Fprintln(writer, ans)
}

func solve(turn, a, b int) int {
	// 기저 사례: 턴이 m보다 커지면 재귀 호출 종료
	if turn > m {
		return 0
	}

	ret := &dp[turn][a][b]

	if *ret != 0 {
		return *ret
	}

	*ret = INF // 최솟값 비교를 위해 충분히 큰 값으로 초기화

	pos := order[turn] // 현재 사용할 벽장의 위치

	// 브루트포스: 가능한 모든 경우의 수 고려

	// 1. a와 b의 위치가 현재 사용할 벽장의 위치보다 큰 경우
	// a와 b중 더 가까운 위치에 있는 위치로만 이동
	if a > pos && b > pos {
		if a > b {
			*ret = min(*ret, solve(turn+1, a, pos)+abs(pos-b))
		} else {
			*ret = min(*ret, solve(turn+1, pos, b)+abs(pos-a))
		}
		// 2. a와 b의 위치가 현재 사용할 벽장의 위치보다 작은 경우
		// a와 b중 더 가까운 위치에 있는 위치로만 이동
	} else if a < pos && b < pos {
		if a > b {
			*ret = min(*ret, solve(turn+1, pos, b)+abs(pos-a))
		} else {
			*ret = min(*ret, solve(turn+1, a, pos)+abs(pos-b))
		}
		// 3. a와 b의 위치가 현재 사용할 벽장의 위치와 동일하거나 좌우로 분포되어 있는 경우
		// a와 b 양쪽으로 이동하는 경우를 모두 고려
	} else {
		*ret = min(*ret, solve(turn+1, pos, b)+abs(pos-a))
		*ret = min(*ret, solve(turn+1, a, pos)+abs(pos-b))
	}

	return *ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
