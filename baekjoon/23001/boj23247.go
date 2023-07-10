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

	m, n int
	psum [301][301]int
)

// 난이도: Silver 1
// 메모리: 2320KB
// 시간: 68ms
// 분류: 누적합, 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	m, n = scanInt(), scanInt()
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			psum[i][j] = scanInt() + psum[i-1][j] + psum[i][j-1] - psum[i-1][j-1] // 누적합
		}
	}
}

func Solve() {
	cnt := 0 // 누적합이 10인 경우의 수

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// (1, 1)에서 (i, j)까지의 누적합이 10이상인 경우만 탐색
			for k := i; k <= m; k++ {
				for l := j; l <= n; l++ {
					temp := psum[k][l] - psum[i-1][l] - psum[k][j-1] + psum[i-1][j-1]
					// 누적합이 10이상인 경우에는 더 탐색할 필요가 없으므로 break
					if temp >= 10 {
						// 누적합이 정확히 10인 경우 cnt++
						if temp == 10 {
							cnt++
						}
						break
					}
				}
			}
		}
	}
	fmt.Fprintln(writer, cnt)
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
