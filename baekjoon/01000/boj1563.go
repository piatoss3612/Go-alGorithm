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
	N       int
	dp      [1001][3][4]int // dp[i][j][k]: i일에 j번 지각 또는 k번 연속 결석했을 경우, 개근상을 받을 수 있는 경우의 수
)

const MOD = 1000000

// 메모리: 1228KB
// 시간: 4ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	ans := solve(1, 0, 0)
	fmt.Fprintln(writer, ans)
}

func solve(day, late, absent int) int {
	// 기저 사례1: 지각을 2회 했거나 3일 연속 결석인 경우
	if late == 2 || absent == 3 {
		return 0
	}

	// 기저 사례2: 개근상을 받을 수 있는 경우
	if day > N {
		return 1
	}

	ret := &dp[day][late][absent]
	if *ret != 0 {
		return *ret
	}

	// 결석은 연속된 값이어야 하므로 출석이나 지각을 하는 경우에는 다시 0으로 초기화된다

	*ret += solve(day+1, late, 0) % MOD        // 출석
	*ret += solve(day+1, late+1, 0) % MOD      // 지각
	*ret += solve(day+1, late, absent+1) % MOD // 결석
	*ret %= MOD

	return *ret
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
