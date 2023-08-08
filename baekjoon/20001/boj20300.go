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

	N   int
	arr []int
)

// 난이도: Silver 3
// 메모리: 1576KB
// 시간: 12ms
// 분류: 그리디 알고리즘, 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 0; i < N; i++ {
		arr = append(arr, scanInt())
	}
	sort.Ints(arr)
}

func Solve() {
	N -= N % 2
	ans := 0
	for i := 0; i < N/2; i++ {
		ans = max(ans, arr[i]+arr[N-i-1])
	}
	fmt.Fprintln(writer, ans)
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
