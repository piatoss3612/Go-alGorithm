package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, M, K int
	hero    []int
	boss    [][2]int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

// 29792번: 규칙적인 보스돌이
// https://www.acmicpc.net/problem/29792
// 난이도: 골드 5
// 메모리: 860 KB
// 시간: 4 ms
// 분류: 다이나믹 프로그래밍, 배낭 문제, 브루트포스 알고리즘
func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()
	hero = make([]int, N)
	for i := 0; i < N; i++ {
		hero[i] = scanInt()
	}

	boss = make([][2]int, K)
	for i := 0; i < K; i++ {
		boss[i][0], boss[i][1] = scanInt(), scanInt()
	}
}

func Solve() {
	sort.Slice(hero, func(i, j int) bool {
		return hero[i] > hero[j]
	})

	total := 0

	for i := 0; i < M; i++ {
		dp := make([]int, 901)
		dp[0] = 1
		maxEarn := 0

		for j := 0; j < K; j++ {
			taken := boss[j][0] / hero[i]
			if boss[j][0]%hero[i] != 0 {
				taken += 1
			}

			for k := 900; k >= taken; k-- {
				if dp[k-taken] > 0 {
					dp[k] = max(dp[k], dp[k-taken]+boss[j][1])
					maxEarn = max(maxEarn, dp[k])
				}
			}
		}

		if maxEarn > 0 {
			total += maxEarn - 1
		}
	}

	fmt.Fprintln(writer, total)
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
