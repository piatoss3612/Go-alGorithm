package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	isP     [2500][2500]int // 문자열의 특정 구간이 팰린드롬인지 체크
	dp      [2500]int       // 팰린드롬 분할 개수의 최솟값
)

// 메모리: 37516KB
// 시간: 80ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	b := scanBytes()
	n := len(b)

	// 팰린드롬 판별
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			// 부분 문자열 길이가 1인 경우는 반드시 팰린드롬
			if i == j {
				isP[i][j] = 1
				continue
			}

			// 부분 문자열 길이가 2인 경우
			if i-j == 1 && b[i] == b[j] {
				isP[j][i] = 1
				continue
			}

			// 부분 문자열 길이가 3이상인 경우
			if b[i] == b[j] && isP[j+1][i-1] == 1 {
				isP[j][i] = 1
			}
		}
	}

	// 팰린드롬 분할 개수의 최솟값 찾기
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isP[i][j] == 1 { // i~j구간이 팰린드롬일 경우

				if i == 0 { // 0에서 시작하는 팰린드롬인 경우 예외 처리
					dp[j] = 1
					continue
				}

				if dp[j] == 0 { // dp[j] 값이 처음 갱신되는 경우
					dp[j] = dp[i-1] + 1 // i이전의 팰린드롬 분할 개수 최솟값 + i~j구간
				} else {
					dp[j] = min(dp[j], dp[i-1]+1) // 현재 dp[j]값과 dp[i-1]+1 중 최솟값으로 갱신
				}
			}
		}
	}
	fmt.Fprintln(writer, dp[n-1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
