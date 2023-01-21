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
	dp      [100001]int
	// dp[i]가 INF인 경우: i개의 컴퓨터를 연결할 수 없다
	// dp[i]가 INF 이외의 값인 경우: i개의 컴퓨터를 연결할 수 있는 최소 설치 비용
)

const INF = 1000000000*300 + 1 // 비용이 10**9인 스위치를 300개 연결하는 경우보다 1 큰 수

// 난이도: Gold 3
// 메모리: 1708KB
// 시간: 56ms
// 분류: 다이나믹 프로그래밍, 배낭 문제

// '은규는 깔끔한 연결을 좋아하기 때문에 스위치에 남는 포트가 없도록 연결하려고 합니다.' 처음에 이 조건을 인지하지 못하고
// '10개의 컴퓨터를 연결하려고 할 때 10개의 컴퓨터를 연결하는 비용보다 11개의 컴퓨터를 연결하는 비용이 더 작으면 11개의 컴퓨터를 연결하는 비용이 정답아닌가?'
// 라고 생각했다가 틀렸다. 이 글을 보시는 분들은 주의하셨으면 한다.
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 0; i <= 100000; i++ {
		dp[i] = INF // 최솟값 비교를 위해 INF로 초기화
	}
	dp[1] = 0 // 스위치 없이 랜선에 직접 연결하는 경우는 비용이 0
}

func Solve() {
	var a, b int
	// 배낭 문제: 스위치를 배낭에 담기
	for i := 1; i <= N; i++ {
		a, b = scanInt(), scanInt() // i번째 스위치의 포트 수 a, 설치 비용 b
		for j := 100000; j >= 1; j-- {
			// j개의 컴퓨터를 연결할 수 있는 경우
			if dp[j] != INF {
				// i번째 스위치를 설치하려면 스위치끼리 연결에 필요한 2개의 포트를 제외해야 한다
				if j-2+a <= 100000 {
					// j-2+a개의 컴퓨터를 연결할 수 있는 경우의 설치 비용의 최솟값 갱신
					dp[j-2+a] = min(dp[j-2+a], dp[j]+b)
				}
			}
		}
	}

	M = scanInt()
	ans := dp[M]

	if ans == INF {
		// M개의 컴퓨터를 연결할 수 없는 경우
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
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
