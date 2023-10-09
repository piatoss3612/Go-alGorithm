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

	N     int
	works []int
	dp    [7]bool
)

// 난이도: Silver 1
// 메모리: 1024KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	works = make([]int, N+1)
	for i := 1; i <= N; i++ {
		works[i] = scanInt()
	}
}

func Solve() {
	dp[0] = true
	for i := 1; i <= N; i++ {
		candidates := []int{} // 중복된 값들이 여러번 등장할 수 있으므로, 중복을 제거하기 위해 후보자를 저장
		for j := 0; j <= 6; j++ {
			if dp[j] {
				candidates = append(candidates, (j+works[i])%7)
			}
		}
		for _, candidate := range candidates {
			dp[candidate] = true
		}
	}

	// 금요일에 일을 끝마치는 시점이 존재하면 YES, 아니면 NO
	if dp[4] {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
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
