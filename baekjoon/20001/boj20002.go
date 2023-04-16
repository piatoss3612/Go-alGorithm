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

	N    int
	psum [301][301]int
)

// 난이도: Gold 5
// 메모리: 1916KB
// 시간: 44ms
// 분류: 브루트포스 알고리즘, 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			psum[i][j] = scanInt()
			psum[i][j] += psum[i-1][j] + psum[i][j-1] - psum[i-1][j-1]
		}
	}
}

func Solve() {
	ans := -987654321

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			for k := 1; k <= min(i, j); k++ {
				temp := psum[i][j] - psum[i-k][j] - psum[i][j-k] + psum[i-k][j-k]
				ans = max(ans, temp)
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
