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

	N     int
	check [2000001]bool
	arr   [21]int
)

// 난이도: Silver 1
// 메모리: 1420KB
// 시간: 16ms
// 분류: 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	rec(1, 0) // 가능한 모든 합을 구함

	// 1부터 2000000까지 순회하면서 check가 false인 값을 출력하고 종료
	for i := 1; i <= 2000000; i++ {
		if !check[i] {
			fmt.Fprintln(writer, i)
			return
		}
	}
}

func rec(n, sum int) {
	check[sum] = true

	if n == N+1 {
		return
	}

	rec(n+1, sum)
	rec(n+1, sum+arr[n])
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
