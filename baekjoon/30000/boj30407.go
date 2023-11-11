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
	N, H, D, K int
	punch      [19]int
	dp         [19][101][2]int
)

const INF = 987654321

// 30407번: 나비의 간식을 훔쳐먹은 춘배
// https://www.acmicpc.net/problem/30407
// 난이도: 골드 4
// 메모리: 900 KB
// 시간: 4 ms
// 분류: 다이나믹 프로그래밍, 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	H, D, K = scanInt(), scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		punch[i] = scanInt()
	}
}

func Solve() {
	ans := H - rec(0, D, 0)
	if ans <= 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func rec(turn, dist, surprise int) int {
	if turn == N {
		return 0
	}

	ret := &dp[turn][dist][surprise]
	if *ret != 0 {
		return *ret
	}

	*ret = INF

	// 1번 웅크리기
	damage := notNegative(punch[turn+1] - dist)
	*ret = min(*ret, rec(turn+1, dist, surprise)+damage/2)

	// 2번 네발로 걷기
	damage = notNegative(punch[turn+1] - (dist + K))
	*ret = min(*ret, rec(turn+1, dist+K, surprise)+damage)

	// 3번 깜짝 놀라게 하기
	if surprise == 0 && turn < N-1 {
		damage := notNegative(punch[turn+1] - dist)
		// 깜짝 놀라게 한 경우, 다음 턴 나비의 행동을 스킵할 수 있으므로
		// 그 때 네발로 걷기를 사용하여 거리를 벌리는 것이 최선의 선택이다.
		*ret = min(*ret, rec(turn+2, dist+K, 1)+damage)
	}

	return *ret
}

func notNegative(a int) int {
	if a < 0 {
		return 0
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
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
