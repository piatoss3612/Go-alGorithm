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
	s1, s2 := scanBytes(), scanBytes() // 대문자 알파벳 비교를 위해 문자열을 바이트 슬라이스로 가져옴
	dp := make([][]int, len(s2)+1)     // 행의 갯수: s2의 길이 + 1
	dp[0] = make([]int, len(s1)+1)     // 열의 갯수: s1의 길이 + 1
	for i := 1; i <= len(s2); i++ {
		dp[i] = make([]int, len(s1)+1)
		for j := 1; j <= len(s1); j++ {
			if dp[i-1][j] > dp[i][j-1] { // 이전에 비교한 두 개의 값 중 큰 값으로 dp[i][j]를 초기화
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
			if s1[j-1] == s2[i-1] { // s1의 i번째 문자와 s2의 j번째 문자가 같은 경우
				dp[i][j] = dp[i-1][j-1] + 1 // 대각선있는 값에 1을 더해준 값을 dp[i][j]에 초기화
			}
		}
	}
	/*
		입력:
		ACAYKP
		CAPCAK

		dp:
		     A C A Y K P
		  [0 0 0 0 0 0 0]
		C [0 0 1 1 1 1 1]
		A [0 1 1 2 2 2 2]
		P [0 1 1 2 2 2 3]
		C [0 1 2 2 2 2 3]
		A [0 1 2 3 3 3 3]
		K [0 1 2 3 3 4 4]

		출력:
		4
	*/
	fmt.Fprintln(writer, dp[len(s2)][len(s1)])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
