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
	N       int
	arr     [100001]int
	dp      [100001]int
)

// 28422번: XOR 카드 게임
// https://www.acmicpc.net/problem/28422
// 난이도: Gold 5
// 메모리: 9328 KB
// 시간: 28 ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
		dp[i] = -1
	}
	dp[0] = -1
}

func Solve() {
	ans := rec(0)
	if ans < 0 {
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func rec(x int) int {
	if N-x < 2 {
		if x == N {
			return 0
		}
		return -987654321
	}

	ret := &dp[x]
	if *ret != -1 {
		return *ret
	}

	*ret = 0

	cnt := 0
	xor := arr[x+1] ^ arr[x+2]
	for i := 1; i <= xor; i = i << 1 {
		if i&xor > 0 {
			cnt += 1
		}
	}

	*ret = max(*ret, rec(x+2)+cnt)

	if x+3 <= N {
		cnt := 0
		xor := xor ^ arr[x+3]
		for i := 1; i <= xor; i = i << 1 {
			if i&xor > 0 {
				cnt += 1
			}
		}
		*ret = max(*ret, rec(x+3)+cnt)
	}

	return *ret
}

func max(a, b int) int {
	if a > b {
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
