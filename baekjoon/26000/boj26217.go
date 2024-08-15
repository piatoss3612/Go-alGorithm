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

// 26217번: 별꽃의 세레나데 (Easy)
// hhttps://www.acmicpc.net/problem/26217
// 난이도: 골드 5
// 메모리: 868 KB
// 시간: 4 ms
// 분류: 수학, 확률론, 기댓값의 선형성
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt() // 씨앗의 개수
}

func Solve() {
	if N == 1 {
		fmt.Fprintln(writer, 1)
		return
	}

	// 조화급수의 합
	// 1 + 1/2 + 1/3 + 1/4 + ... + 1/N
	var sum float64 = 0
	for i := 1; i <= N; i++ {
		sum += 1 / float64(i)
	}

	sum *= float64(N) // 필요한 씨앗 개수의 기댓값 = N * 조화급수의 합

	fmt.Fprintln(writer, sum)
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
