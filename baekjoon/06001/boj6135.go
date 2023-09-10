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
	N, M, T int
	path    [301][301]int
)

const INF = 987654321

// 난이도: Gold 4
// 메모리: 2236KB
// 시간: 56ms
// 분류: 플로이드 와샬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M, T = scanInt(), scanInt(), scanInt()

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			path[i][j] = INF
		}
	}

	for i := 1; i <= M; i++ {
		s, e, h := scanInt(), scanInt(), scanInt()
		path[s][e] = h
	}
}

func Solve() {
	// 소들은 허들을 몇 개든 뛰어넘을 수 있지만 허들의 높이가 높아질수록 스트레스를 많이 받는다.
	// 따라서 플로이드 와샬 알고리즘을 사용하여 허들의 개수에 상관없이 허들의 높이의 최댓값이 최소가 되는 경로를 찾으면 된다.
	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if i == j {
					continue
				}

				path[i][j] = min(path[i][j], max(path[i][k], path[k][j])) // 허들의 높이의 최댓값이 최소가 되는 경로를 찾는다.
			}
		}
	}

	for i := 1; i <= T; i++ {
		s, e := scanInt(), scanInt()
		if path[s][e] == INF { // 경로가 없는 경우
			fmt.Fprintln(writer, -1)
		} else { // 경로가 있는 경우
			fmt.Fprintln(writer, path[s][e])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
