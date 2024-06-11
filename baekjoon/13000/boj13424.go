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
	T       int
	N, M, K int
	board   [101][101]int
	friends []int
)

const INF = 1 << 31

// 13424번: 비밀 모임
// hhttps://www.acmicpc.net/problem/13424
// 난이도: Gold 4
// 메모리: 960 KB
// 시간: 16 ms
// 분류:
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	T = scanInt()
	for t := 0; t < T; t++ {
		Setup()
		Solve()
	}
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			board[i][j] = INF
		}
		board[i][i] = 0
	}

	for i := 1; i <= M; i++ {
		u, v, w := scanInt(), scanInt(), scanInt()
		board[u][v] = w
		board[v][u] = w
	}
	K = scanInt()
	friends = make([]int, K)
	for i := 0; i < K; i++ {
		friends[i] = scanInt()
	}
}

func Solve() {
	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if board[i][j] > board[i][k]+board[k][j] {
					board[i][j] = board[i][k] + board[k][j]
				}
			}
		}
	}

	minTotalDist := INF
	minRoom := 0
	for i := 1; i <= N; i++ {
		totalDist := 0
		for _, f := range friends {
			totalDist += board[i][f]
		}

		if totalDist < minTotalDist {
			minTotalDist = totalDist
			minRoom = i
		}
	}

	fmt.Fprintln(writer, minRoom)
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
