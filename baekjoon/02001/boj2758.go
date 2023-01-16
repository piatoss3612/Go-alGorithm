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
	T, n, m int
	dp      [11][2001]int // dp[i][j]: i개의 숫자를 선택했을 때 마지막 숫자가 j인 경우의 수
)

// 난이도: Gold 4
// 메모리: 1196KB
// 시간: 12ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Solve()
	Input()
}

func Solve() {
	// 1개의 숫자를 선택했을 때 마지막 숫자가 i인 경우의 수는 1
	for i := 1; i <= 2000; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= 10; i++ {
		for j := 0; j <= 2000; j++ {
			// i-1개를 선택했을 때 마지막 숫자가 j인 경우가 있다면
			if dp[i-1][j] != 0 {
				// i번째 숫자로 이전에 고른 j보다 적어도 2배가 되는 숫자가 오는 경우의 수 갱신
				for k := j * 2; k <= 2000; k++ {
					dp[i][k] += dp[i-1][j]
				}
			}
		}
	}
}

func Input() {
	T = scanInt()
	for i := 1; i <= T; i++ {
		n, m = scanInt(), scanInt()
		ans := 0

		/*
			1에서 m사이의 숫자들 중에 n개의 숫자를 고를 경우 복권을 구매하는 개수 X는

			n개의 숫자를 선택했을 때 마지막 숫자가 m인 경우의 수를 F(n, m)이라고 했을 때,

			X = F(n, 0) + F(n, 1) + F(n, 2) + ... + F(n, m-2) + F(n, m-1) + F(n, m)
		*/
		for j := 0; j <= m; j++ {
			ans += dp[n][j]
		}
		fmt.Fprintln(writer, ans)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
