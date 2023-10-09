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

	N  int
	dp [1001]int
)

// 난이도: Silver 1
// 메모리: 912KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
}

func Solve() {
	sum := (N * (N + 1)) / 2 // 1부터 N까지의 합

	// 합이 홀수면 동일한 합을 가진 두 집합을 만들 수 없음
	if sum%2 == 1 {
		fmt.Fprintln(writer, 0)
		return
	}

	dp[0] = 1 // 0을 만들 수 있는 경우의 수는 1개

	// 1부터 N까지의 숫자를 하나씩 배낭에 넣어보면서 합이 sum/2가 되는 경우의 수를 구함
	for i := 1; i <= N; i++ {
		for j := sum / 2; j >= i; j-- {
			dp[j] = (dp[j] + dp[j-i])
		}
	}

	// 중복되는 경우의 수를 제외하기 위해 2로 나눔
	fmt.Fprintln(writer, dp[sum/2]/2) 
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
