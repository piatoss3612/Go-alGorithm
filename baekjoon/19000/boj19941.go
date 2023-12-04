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
	N, K    int
	table   []byte
	taken   []bool
)

// 19941번: 햄버거 분배
// https://www.acmicpc.net/problem/19941
// 난이도: 실버 3
// 메모리: 964 KB
// 시간: 8 ms
// 분류: 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	table = append([]byte{' '}, scanBytes()...)
	taken = make([]bool, N+1)
}

func Solve() {
	for i := 1; i <= N; i++ {
		if table[i] == 'H' {
			for j := i - K; j <= i+K; j++ {
				if j < 1 || j > N {
					continue
				}

				if table[j] == 'P' && !taken[j] {
					taken[j] = true
					break
				}
			}
		}
	}

	cnt := 0
	for i := 1; i <= N; i++ {
		if taken[i] {
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt)
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
