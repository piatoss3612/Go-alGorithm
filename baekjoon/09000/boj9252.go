package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	s1 := []rune(scanString())
	s2 := []rune(scanString())

	l1, l2 := len(s1), len(s2)

	dp := make([][]int, l1+1)
	dp[0] = make([]int, l2+1)

	for i := 1; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		for j := 1; j <= l2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	/*
		예제 입력:
		ACAYKP
		CAPCAK

		dp:
		     A C A Y K P
		  [0 0 0 0 0 0 0]
		C [0 0 1 1 1 1 1]
		A [0 1 1 1 2 2 2]
		P [0 1 2 2 2 3 3]
		C [0 1 2 2 2 3 3]
		A [0 1 2 2 2 3 4]
		K [0 1 2 3 3 3 4]

		예제 출력:
		4
		ACAK
	*/

	total := dp[l1][l2]
	res := []rune{}

	// 역추적
	for dp[l1][l2] != 0 {
		// dp[i][j] = max(dp[i-1][j], dp[i][j-1])를 실행한 경우를 먼저 탐색하고 뒤로 이동
		// 다음으로 대각선을 체크하여 같은 문자열이 나와서 +1이 된 경우를 찾는다
		if dp[l1][l2] == dp[l1][l2-1] {
			l2 -= 1
		} else if dp[l1][l2] == dp[l1-1][l2] {
			l1 -= 1
		} else if dp[l1][l2]-1 == dp[l1-1][l2-1] {
			res = append(res, s1[l1-1])
			l1 -= 1
			l2 -= 1
		}
	}

	fmt.Fprintln(writer, total)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Fprint(writer, string(res[i]))
	}
	fmt.Fprintln(writer)
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
