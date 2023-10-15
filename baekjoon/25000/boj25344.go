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

	N int
)

// 25344번: 행성 정렬
// https://www.acmicpc.net/problem/25344
// 난이도: 실버 4
// 메모리:864 KB
// 시간: 20 ms
// 분류: 수학, 정수론, 유클리드 호제법
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
  ans := scanInt()
  for i := 2; i <= N-2; i++ {
    x := scanInt()
    ans = lcm(ans, x)
  }
  fmt.Fprintln(writer, ans)
}

func lcm(a, b int) int {
  return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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