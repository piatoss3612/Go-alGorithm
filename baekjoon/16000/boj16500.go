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
	words   []string
	dp      []int
)

// 16500번: 문자열 판별
// hhttps://www.acmicpc.net/problem/16500
// 난이도: 골드 5
// 메모리: 892 KB
// 시간: 4 ms
// 분류: 다이나믹 프로그래밍, 문자열
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	S = scanString()
	N = scanInt()
	words = make([]string, N)
	for i := 0; i < N; i++ {
		words[i] = scanString()
	}
}

func Solve() {
	l := len(S)
	dp = make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = -1
	}

	if rec(0) > 0 {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, 0)
	}
}

func rec(idx int) int {
	if idx == len(S) {
		return 1
	}

	ret := dp[idx]
	if ret != -1 {
		return ret
	}

	ret = 0

	for _, word := range words {
		if idx+len(word) > len(S) {
			continue
		}

		if S[idx:idx+len(word)] == word {
			ret += rec(idx + len(word))
		}
	}

	dp[idx] = ret
	return ret
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
