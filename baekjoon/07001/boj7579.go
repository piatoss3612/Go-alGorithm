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
	apps    []App
	dp      [10001]int
)

type App struct {
	memory int // m
	cost   int // c
}

// 난이도: Gold 3
// 메모리: 8792KB -> 1000KB
// 시간: 12ms -> 8ms
// 분류: 다이나믹 프로그래밍, 배낭 문제

// # 1차, 2차 시도
// M과 m의 크기가 굉장히 커서 dp를 인덱싱하기 위해 사용하는 것은 어렵지만
// N과 c의 범위가 작아 dp의 행을 N으로 잡고 열을 비용의 최댓값(100*100)으로 잡았다

// # 3차 시도
// dp를 비용에 해당하는 메모리의 최댓값만 저장하도록 1차원 배열로 초기화하여 시도
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	apps = make([]App, N+1)
	for i := 1; i <= N; i++ {
		apps[i].memory = scanInt()
	}
	for i := 1; i <= N; i++ {
		apps[i].cost = scanInt()
	}

	// 비용이 0인 경우를 고려하여 dp를 -1로 초기화한다
	for i := 1; i <= 10000; i++ {
		dp[i] = -1
	}
	dp[0] = 0 // 아무 앱도 종료되지 않은 상태의 메모리 크기를 0으로 초기화
}

func Solve() {
	ans := 987654321

	// 배낭 문제: i(1~N)번 앱을 비활성화(배낭에 담기)하는 경우
	// 비활성화된 앱들의 비용이 j(0~10000)일 때의 확보된 메모리의 최대 크기 구하기
	for i := 1; i <= N; i++ {
		// 중복 연산을 피하기 위해 j를 10000에서부터 역순으로 진행
		for j := 10000; j >= 0; j-- {
			// i번째 앱을 비활성화 하는 경우
			// 이전 단계의 앱을 비활성화 한 기록이 있거나 아무 앱도 종료되지 않은 상태여야 한다
			if dp[j] >= 0 {
				// i번째 앱을 비활성화 시켰을 경우의 비용(j+app[i].cost)와 매칭되는 여유 메모리의 크기를 최댓값으로 갱신
				dp[j+apps[i].cost] = max(dp[j+apps[i].cost], dp[j]+apps[i].memory)

				// 확보된 메모리의 크기가 M보다 크거나 같은 경우
				// 비용의 최솟값 갱신
				if dp[j+apps[i].cost] >= M {
					ans = min(ans, j+apps[i].cost)
				}
			}
		}
	}

	fmt.Fprintln(writer, ans)
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
