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
	a, b    int
)

// 31804번: Chance!
// hhttps://www.acmicpc.net/problem/31804
// 난이도: 골드 4
// 메모리: 32576 KB
// 시간: 152 ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	a, b = scanInt(), scanInt()
	dp := [1000001][2]int{} // dp[i][j]: a에서 i를 만드는데 필요한 마법 사용의 최소 횟수 (j는 0, 1로 chance!를 사용했는지 여부)
	dp[a][0] = 1            // 시작 위치를 1로 표시

	for i := a; i <= b; i++ {
		magic := func(i, target, from, to int) {
			if target <= b {
				if dp[target][to] == 0 {
					dp[target][to] = dp[i][from] + 1
				} else {
					dp[target][to] = min(dp[target][to], dp[i][from]+1)
				}
			}
		}

		if dp[i][0] > 0 {
			// 물주기
			magic(i, i+1, 0, 0)

			// 밥 주기
			magic(i, i*2, 0, 0)

			// chance!
			magic(i, i*10, 0, 1)
		}

		if dp[i][1] > 0 {
			// 물주기
			magic(i, i+1, 1, 1)

			// 밥 주기
			magic(i, i*2, 1, 1)

			// chance!는 이미 사용해서 안됨
		}
	}

	answer := 987654321

	if dp[b][0] != 0 {
		answer = min(answer, dp[b][0])
	}

	if dp[b][1] != 0 {
		answer = min(answer, dp[b][1])
	}

	fmt.Fprintln(writer, answer-1)
}

func Solve() {
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
