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
	N, M, K int
	startw  [2001][2001]int // startw[i][j] = (1, 1)부터 (i, j)까지의 체스판이 흰색으로 시작할 때 변경해야 하는 칸의 수의 합
	startb  [2001][2001]int // startb[i][j] = (1, 1)부터 (i, j)까지의 체스판이 검은색으로 시작할 때 변경해야 하는 칸의 수의 합
)

// 난이도: Gold 5
// 메모리: 68496KB
// 시간: 144ms
// 분류: 누적합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()

	for i := 1; i <= N; i++ {
		line := scanString()
		for j := 1; j <= M; j++ {
			if (i+j)%2 == 0 { // 행과 열의 합이 짝수일 때 흰색으로 시작하는 체스판의 (i, j) 칸은 흰색이어야 하고 검은색으로 시작하는 체스판의 (i, j) 칸은 검은색이어야 한다.
				if line[j-1] == 'B' { // (i, j) 칸이 검은색이면 흰색으로 시작하는 체스판에서 색을 변경해야 한다.
					startw[i][j] = 1
				} else { // (i, j) 칸이 흰색이면 검은색으로 시작하는 체스판에서 색을 변경해야 한다.
					startb[i][j] = 1
				}
			} else { // 행과 열의 합이 홀수일 때 흰색으로 시작하는 체스판의 (i, j) 칸은 검은색이어야 하고 검은색으로 시작하는 체스판의 (i, j) 칸은 흰색이어야 한다.
				if line[j-1] == 'B' { // (i, j) 칸이 검은색이면 검은색으로 시작하는 체스판에서 색을 변경해야 한다.
					startb[i][j] = 1
				} else { // (i, j) 칸이 흰색이면 흰색으로 시작하는 체스판에서 색을 변경해야 한다.
					startw[i][j] = 1
				}
			}

			// 누적합
			startw[i][j] += startw[i-1][j] + startw[i][j-1] - startw[i-1][j-1]
			startb[i][j] += startb[i-1][j] + startb[i][j-1] - startb[i-1][j-1]
		}
	}
}

func Solve() {
	ans := 987654321

	// (i, j)부터 (i+K-1, j+K-1)까지의 체스판에서 변경해야 하는 칸의 수의 최솟값을 구한다.
	for i := 1; i <= N-K+1; i++ {
		for j := 1; j <= M-K+1; j++ {
			w := startw[i+K-1][j+K-1] - startw[i+K-1][j-1] - startw[i-1][j+K-1] + startw[i-1][j-1]
			b := startb[i+K-1][j+K-1] - startb[i+K-1][j-1] - startb[i-1][j+K-1] + startb[i-1][j-1]
			if w < ans {
				ans = w
			}
			if b < ans {
				ans = b
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", ans)
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
