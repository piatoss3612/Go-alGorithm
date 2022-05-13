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
	INF     = 1<<63 - 1
	inp     [501]mtx
	dp      [501][501]int // 메모리를 동적으로 할당하는 과정하고 메모리와 시간 차이가 큰 편
)

// 행렬의 행과 열의 크기를 가지는 구조체
type mtx struct {
	r, c int
}

// 메모리: 4836KB
// 시간: 96ms
// 시간복잡도 O(n^3)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	for i := 1; i <= n; i++ {
		inp[i] = mtx{scanInt(), scanInt()}
	}

	/*
		예제 입력:
		3
		5 3
		3 2
		2 6

		dp:
		0 30 90
		0 0 36
		0 0 0

		풀이 과정:

		1. i=1, j=1, k=1
			dp[1][2] = inp[1].r*inp[1].c*inp[2].c = 30

		2. i=1, j=2, k=2
			dp[2][3] = inp[2].r*inp[2].c*inp[3].c = 36

		3. i=2, j=1, k=1
			dp[1][3] = dp[2][3] + inp[1].r*inp[1].c*inp[3].c = 36 + 90 = 126

		4. i=2, j=1, k=2
			dp[1][3] = dp[1][2] + inp[1].r*inp[2].c*inp[3].c = 30 + 60 = 90

		예제 출력:
		90
	*/

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			dp[j][i+j] = INF
			for k := j; k < i+j; k++ {
				dp[j][i+j] = min(dp[j][i+j], dp[j][k]+dp[k+1][i+j]+inp[j].r*inp[k].c*inp[i+j].c)
			}
		}
	}
	fmt.Fprintln(writer, dp[1][n]) // 전체 구간의 최솟값을 출력
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
