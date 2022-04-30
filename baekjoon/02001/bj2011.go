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

// 메모리: 1036KB
// 시간: 4ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	b := scanBytes()
	b = append([]byte{'0'}, b...)
	var dp [5001][3]int

	// 2591번 문제에서 시작 문자가 0인 경우를 예외 처리해줘야 한다
	if int(b[1]-'0') != 0 {
		dp[1][1] = 1
	}

	n := len(b) - 1

	for i := 2; i <= n; i++ {
		prev := int(b[i-1] - '0')
		cur := int(b[i] - '0')

		if cur != 0 {
			dp[i][1] += (dp[i-1][1] + dp[i-1][2]) % 1000000
		}

		if prev*10+cur <= 26 {
			dp[i][2] += dp[i-1][1] % 1000000
		}
	}
	fmt.Fprintln(writer, (dp[n][1]+dp[n][2])%1000000)
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
