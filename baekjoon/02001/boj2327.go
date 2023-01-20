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
	H, N    int
	dp      [100001]int // dp[i]: 팀 구성원들의 키의 합이 i일 때, 구성원 중 가장 느린 사람의 달리기 속도의 최댓값
)

const INF = 987654321

// 난이도: Gold 3
// 메모리: 1684KB
// 시간: 56ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	H, N = scanInt(), scanInt()
	for i := 0; i <= 100000; i++ {
		dp[i] = -1
	}
	dp[0] = INF // 구성원들의 키의 합이 0인 경우는 없지만, 학생 한 명으로만 구성된 팀이 존재할 수 있으므로 dp[0]을 최솟값 비교를 위해 INF로 초기화
}

func Solve() {
	for i := 1; i <= N; i++ {
		h, s := scanInt(), scanInt() // i번째 학생의 키와 달리기 속도
		for j := H - 1; j >= 0; j-- {
			// 구성원들의 키의 합이 j인 팀이 존재하고, i번째 학생의 키 h를 더했을 때 경곗값 H보다 작거나 같은 경우에만
			if dp[j] != -1 && j+h <= H {
				// 구성원들의 키의 합이 j인 팀에 i번째 학생이 추가되었을 경우의
				// 구성원 중 가장 느린 사람의 달리기 속도의 최댓값을 갱신

				// min(dp[j], s): i번째 학생을 팀에 추가했을 때의 가장 느린 사람의 달리기 속도
				// max(dp[j+h], min(dp[j], s)): 구성원들의 키의 합이 j+h인 팀의 구성원 중 가장 느린 사람의 달리기 속도의 최댓값
				dp[j+h] = max(dp[j+h], min(dp[j], s))
			}
		}
	}

	fmt.Fprintln(writer, dp[H])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
