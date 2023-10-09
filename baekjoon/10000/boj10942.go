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
	graph   [2001]int
	dp      [2001][2001]int
)

// 메모리: 28044KB
// 시간: 316ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	for i := 1; i <= n; i++ {
		graph[i] = scanInt()
		for j := i; j >= 1; j-- {
			// 길이가 1이면 항상 팰린드롬
			if i == j {
				dp[i][j] = 1
				continue
			}

			// 길이가 2인 팰린드롬 체크
			if i-j == 1 && graph[i] == graph[j] {
				dp[j][i] = 1
				continue
			}

			// 길이가 3이상인 팰린드롬 체크
			if graph[i] == graph[j] && dp[j+1][i-1] == 1 {
				dp[j][i] = 1
			}
		}
	}

	/*
		예제 입력:
		7
		1 2 1 3 1 2 1
		4
		1 3
		2 5
		3 3
		5 7

		dp:
		1 0 1 0 0 0 1
		0 1 0 0 0 1 0
		0 0 1 0 1 0 0
		0 0 0 1 0 0 0
		0 0 0 0 1 0 1
		0 0 0 0 0 1 0
		0 0 0 0 0 0 1

		예제 출력:
		1
		0
		1
		1
	*/

	m := scanInt()
	var s, e int
	for i := 1; i <= m; i++ {
		s, e = scanInt(), scanInt()
		fmt.Fprintln(writer, dp[s][e])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
