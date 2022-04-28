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

// 메모리: 2180KB
// 시간: 16ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	var dp [32768][5]int // dp[x][y]: x를 y개의 제곱수로 표현하는 경우의 수
	// 문제에 제시된 라그랑주의 네 제곱수 정리에 따라 1개에서 최대 4개의 제곱수로 표현되는 경우의 수를 저장

	for i := 1; i*i < 32768; i++ {
		dp[i*i][1] = 1 // i의 제곱은 자기 자신으로만 표현되는 제곱수

		// i*i를 사용해 4개 이하의 제곱수로 j를 표현할 수 있는 경우의 수를 찾는다
		for j := i * i; j < 32768; j++ {
			dp[j][2] += dp[j-i*i][1]
			dp[j][3] += dp[j-i*i][2]
			dp[j][4] += dp[j-i*i][3]
		}
	}

	for {
		n := scanInt()
		if n == 0 {
			return
		}

		ans := 0
		for i := 0; i < 5; i++ {
			ans += dp[n][i]
		}
		fmt.Fprintln(writer, ans)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
