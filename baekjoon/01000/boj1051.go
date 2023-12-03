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
	N, M    int
	square  [][]byte
)

// 1051번: 숫자 정사각형
// https://www.acmicpc.net/problem/1051
// 난이도: 실버 3
// 메모리: 868 KB
// 시간: 4 ms
// 분류: 브루트포스 알고리즘, 구현
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	square = make([][]byte, N+1)
	for i := 1; i <= N; i++ {
		square[i] = append([]byte{'0'}, scanBytes()...)
	}
}

func Solve() {
	maxLen := min(N, M)
	for l := maxLen; l >= 1; l-- {
		for i := 1; i <= N-l+1; i++ {
			for j := 1; j <= M-l+1; j++ {
				if square[i][j] == square[i][j+l-1] && square[i][j] == square[i+l-1][j] && square[i][j] == square[i+l-1][j+l-1] {
					fmt.Fprintln(writer, l*l)
					return
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
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
