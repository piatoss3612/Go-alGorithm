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
	M       int
	N       int
	Messi1  = "Messi"
	Messi2  = "Messi Gimossi Messi"
	dp      []int
)

// 17297번: Messi Gimossi
// hhttps://www.acmicpc.net/problem/17297
// 난이도: 골드 4
// 메모리: 908 KB
// 시간: 4 ms
// 분류: 다이나믹 프로그래밍, 분할 정복
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	M = scanInt()
	dp = make([]int, 0, 100000)
	dp = append(dp, 0)

	N = 0

	for dp[N] < M {
		N++
		if N == 1 {
			dp = append(dp, 5)
			continue
		}

		if N == 2 {
			dp = append(dp, 13)
			continue
		}

		dp = append(dp, dp[N-1]+1+dp[N-2])
	}
}

func Solve() {
	divideAndConquer(N, M)
}

func divideAndConquer(n, m int) {
	if n == 1 {
		fmt.Fprintln(writer, string(Messi1[m-1]))
		return
	}

	if n == 2 {
		s := Messi2[m-1]
		if s == ' ' {
			fmt.Fprintln(writer, "Messi Messi Gimossi")
		} else {
			fmt.Fprintln(writer, string(s))
		}
		return
	}

	if m <= dp[n-1] {
		divideAndConquer(n-1, m)
		return
	}

	if m == dp[n-1]+1 {
		fmt.Fprintln(writer, "Messi Messi Gimossi")
		return
	}

	divideAndConquer(n-2, m-dp[n-1]-1)
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
