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

	N   int
	arr []int
	ans int
)

// 난이도: Silver 1
// 메모리: 3240KB
// 시간: 16ms
// 분류: 브루트포스 알고리즘, 백트래킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	arr = make([]int, 0, N)
	for i := 0; i < N; i++ {
		arr = append(arr, scanInt())
	}
}

func Solve() {
	rec(0, 0)
	fmt.Fprintln(writer, ans)
}

func rec(n, sum int) {
	if n == N-2 {
		ans = max(ans, sum)
		return
	}

	for i := 1; i <= len(arr)-2; i++ {
		temp := arr[i-1] * arr[i+1]
		num := arr[i]
		arr = append(arr[:i], arr[i+1:]...)
		rec(n+1, sum+temp)
		arr = append(arr[:i], append([]int{num}, arr[i:]...)...)
	}
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
