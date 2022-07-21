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
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t, w := scanInt(), scanInt()
	var plums [1001]int // 런타임 오류 발생 원인 31 -> 1001로 수정
	for i := 1; i <= t; i++ {
		plums[i] = scanInt()
	}

	var dp [1001][31][3]int // t초에 w번 움직였을 때 자두의 위치 1 또는 2를 저장하는 배열

	// 시작 위치는 1로 고정
	if plums[1] == 1 {
		dp[1][0][1] = 1
	} else {
		dp[1][1][2] = 1
	}

	for i := 2; i <= t; i++ {

		// 움직이지 않는 경우
		if plums[i] == 1 {
			dp[i][0][1] = dp[i-1][0][1] + 1
		} else {
			dp[i][0][1] = dp[i-1][0][1]
		}

		// 1~w번 움직이는 경우
		for j := 1; j <= w; j++ {
			if plums[i] == 1 {
				// 자두가 떨어지는 위치가 1인 경우

				// i초에 j번 움직였을 때 1번 위치에서 가질 수 있는 자두의 최대 개수
				// 2-> 1번으로 1번 이동하는 경우와 1번에 가만히 있는 경우 중 최댓값 + 1
				dp[i][j][1] = max(dp[i-1][j-1][2], dp[i-1][j][1]) + 1

				// i초에 j번 움직였을 때 2번 위치에서 가질 수 있는 자두의 최대 개수
				// 1-> 2번으로 1번 이동하는 경우와 1번에 가만히 있는 경우 중 최댓값
				dp[i][j][2] = max(dp[i-1][j-1][1], dp[i-1][j][2])
			} else {
				// 자두가 떨어지는 위치가 2인 경우
				dp[i][j][1] = max(dp[i-1][j-1][2], dp[i-1][j][1])
				dp[i][j][2] = max(dp[i-1][j-1][1], dp[i-1][j][2]) + 1
			}
		}
	}

	res := 0

	// w번 이내로 움직인 범위에서 가질 수 있는 최댓값 찾기
	for i := 0; i <= w; i++ {
		res = max(res, max(dp[t][i][1], dp[t][i][2]))
	}

	fmt.Fprintln(writer, res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
