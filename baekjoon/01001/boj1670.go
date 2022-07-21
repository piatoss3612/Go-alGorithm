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
	dp      []int
)

const MOD = 987654321

// 메모리: 1512KB
// 시간: 84ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	dp = make([]int, n+1)
	dp[0] = 1 // 0명의 팔이 교차하지 않고 악수할 수 있는 경우 수
	dp[2] = 1 // 2명의 팔이 교차하지 않고 악수할 수 있는 경우 수
	fmt.Fprintln(writer, rec(n))
}

// x명의 팔이 교차하지 않고 악수할 수 있는 경우 수
func rec(x int) int {
	ret := &dp[x]
	if *ret != 0 {
		return *ret
	}

	/*
		x가 10인 경우, 반시계방향으로 진행

		12시 방향의 대표 A가 왼쪽 사람과 악수하는 경우: dp[0] * dp[8]
		악수하는 사람과 대표 A를 제외하고 8명의 팔이 교차하지 않고 악수할 수 있는 경우 수를 구한다

		대표 A가 2명을 건너뛰고 3번째 사람과 악수하는 경우: dp[2] * dp[6]
		악수하는 사람과의 라인을 기준으로 2명, 6명으로 갈라지므로, 2명이 악수하는 경우의 수 * 6명이 악수하는 경우의 수를 구한다

		대표 A가 4명을 건너뛰고 5번째 사람과 악수하는 경우: dp[4] * dp[4]
		악수하는 사람과의 라인을 기준으로 4명, 4명으로 갈라지므로, 4명이 악수하는 경우의 수 * 4명이 악수하는 경우의 수를 구한다

		대표 A가 6명을 건너뛰고 7번째 사람과 악수하는 경우: dp[6] * dp[2]
		악수하는 사람과의 라인을 기준으로 6명, 2명으로 갈라지므로, 6명이 악수하는 경우의 수 * 2명이 악수하는 경우의 수를 구한다

		대표 A가 8명을 건너뛰고 9번째 사람, 즉 오른쪽 사람과 악수하는 경우: dp[8] * dp[0]
		악수하는 사람과 대표 A를 제외하고 8명의 팔이 교차하지 않고 악수할 수 있는 경우 수를 구한다
	*/

	for i := 2; i <= x; i += 2 {
		*ret += dp[i-2] * rec(x-i)
		*ret %= MOD
	}
	return *ret
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
