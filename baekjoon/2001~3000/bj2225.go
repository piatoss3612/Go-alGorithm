package main

import (
	"bufio"
	_ "bytes"
	"fmt"
	_ "io/ioutil"
	_ "math"
	_ "math/big"
	"os"
	_ "sort"
	"strconv"
	_ "strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, k := scanInt(), scanInt()
	dp := make([][]int, n+1)
	dp[0] = make([]int, k+1)
	// 0도 여러 개의 합으로 표현할 수 있는 경우가 1개씩 있으므로 1로 초기화
	for i := 1; i <= k; i++ {
		dp[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
		for j := 1; j <= k; j++ {
			// 정수 1개로 만들 수 있는 경우는 반드시 1가지이므로 1로 초기화
			if j == 1 {
				dp[i][j] = 1
				continue
			}
			// 점화식이 왜 이렇게 나왔는지는 나도 잘 이해가 안되지만
			// 테이블로 정리하다보니 규칙을 찾았다는 것이 포인트
			dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % 1000000000
		}
	}
	/*
		ex)
		입력:
		6 4

		dp:
		[0 1 1 1 1]
		[0 1 2 3 4]
		[0 1 3 6 10]
		[0 1 4 10 20]
		[0 1 5 15 35]
		[0 1 6 21 56]
		[0 1 7 28 84]

		출력:
		84
	*/
	fmt.Fprintln(writer, dp[n][k])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
