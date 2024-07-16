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
)

// 11729번: 하노이 탑 이동 순서
// hhttps://www.acmicpc.net/problem/11729
// 난이도: 골드 5
// 메모리: 864 KB
// 시간: 124 ms
// 분류: 재귀
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
}

func Solve() {
	fmt.Fprintln(writer, 1<<N-1)
	towerOfHanoi(N, 1, 3, 2)
}

func towerOfHanoi(n int, from, to, via int) {
	if n == 1 {
		fmt.Fprintln(writer, from, to)
		return
	}

	towerOfHanoi(n-1, from, via, to)
	fmt.Fprintln(writer, from, to)
	towerOfHanoi(n-1, via, to, from)
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
