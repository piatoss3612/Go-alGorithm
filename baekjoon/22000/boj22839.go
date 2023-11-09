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
	dp      [301]int // less than 300이라고 해놓고 왜 300까지 받는지 모르겠음
)

// 22839번: Square Coins
// https://www.acmicpc.net/problem/22839
// 난이도: 골드 5
// 메모리: 860 KB
// 시간: 4 ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	dp[0] = 1
	for i := 1; i*i <= 300; i++ {
		x := i * i
		for j := x; j <= 300; j++ {
			dp[j] += dp[j-x]
		}
	}
}

func Solve() {
	for {
		n := scanInt()
		if n == 0 {
			break
		}

		fmt.Fprintln(writer, dp[n])
	}
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
