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
	S       string
	N       int
	dp      [3001][2]int
)

const INF = 987654321

// 난이도: Silver 1
// 메모리: 1424KB
// 시간: 8ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	S = scanString()
	N = len(S)
}

func Solve() {
	fmt.Fprintln(writer, rec(0, 0))
}

func rec(idx, cap int) int {
	if idx == N {
		return 0
	}

	ret := &dp[idx][cap]
	if *ret != 0 {
		return *ret
	}

	*ret = INF

	// 대문자인지 소문자인지 확인
	if S[idx] >= 'a' && S[idx] <= 'z' {
		// 소문자인 경우
		// 마름모가 꺼져있는 경우
		if cap == 0 {
			*ret = min(*ret, rec(idx+1, 0)+1) // 소문자 입력
		} else {
			// 마름모가 켜져있는 경우
			*ret = min(*ret, rec(idx+1, 0)+2) // 마름모를 끄고 소문자 입력 
			*ret = min(*ret, rec(idx+1, 1)+2) // 소문자를 입력하고 별 버튼으로 대문자로 변경
		}
	} else {
		// 대문자인 경우
		// 마름모가 꺼져있는 경우
		if cap == 0 {
			*ret = min(*ret, rec(idx+1, 1)+2) // 마름모를 켜고 대문자 입력
			*ret = min(*ret, rec(idx+1, 0)+2) // 대문자를 입력하고 별 버튼으로 소문자로 변경
		} else {
			// 마름모가 켜져있는 경우
			*ret = min(*ret, rec(idx+1, 1)+1) // 대문자 입력
		}
	}
	return *ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
