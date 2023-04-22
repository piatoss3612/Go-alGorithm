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

	N, M int
	arr  [21]int
	dp   [21][21][21]int
)

const INF = -987654321

// 난이도: Gold 5
// 메모리: 1052KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()

	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	ans := rec(0, 0, 0)
	fmt.Fprintln(writer, ans)
}

func rec(here, cnt, group int) int {
	if group == M || here == N && group == M-1 && cnt > 0 {
		return 0
	}

	if here == N {
		return INF
	}

	ret := &dp[here][cnt][group]
	if *ret != 0 {
		return *ret
	}

	*ret = INF

	if cnt == 0 {
		for i := here + 1; i <= N; i++ {
			*ret = max(*ret, rec(i, cnt+1, group)+arr[i])
		}
	} else {
		*ret = max(*ret, rec(here+1, cnt+1, group)+arr[here+1])
		*ret = max(*ret, rec(here, 0, group+1))
	}

	return *ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
