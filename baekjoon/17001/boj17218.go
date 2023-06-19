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

	A, B string
	dp   [][]int
)

// 난이도: Gold 5
// 메모리: 952KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	A, B = scanString(), scanString()
	dp = make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
}

func Solve() {
	// 가장 긴 공통 부분 수열의 길이를 구한다. (연속된 문자열이 아님에 주의)
 	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				continue
			}
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
		}
	}

	// 가장 긴 공통 부분 수열을 구한다.
	result := make([]byte, 0, dp[len(A)][len(B)])
	i, j := len(A), len(B)
	for i > 0 && j > 0 {
		// 현재 위치의 값이 왼쪽, 위쪽 값과 같으면 이전 위치로 이동한다.
		if dp[i][j] == dp[i-1][j] {
			i--
			continue
		}
		if dp[i][j] == dp[i][j-1] {
			j--
			continue
		}

		// 현재 위치의 값이 왼쪽, 위쪽 값과 다르면 문자열에 추가하고 대각선으로 이동한다.
		result = append(result, A[i-1])
		i--
		j--
	}

	// 문자열을 뒤집어서 출력한다.
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, "%c", result[i])
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

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
