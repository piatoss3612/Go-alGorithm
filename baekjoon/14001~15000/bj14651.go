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
	n := scanInt()
	dp := make([]int, 33334)

	dp[2] = 2

	/*
		0,1,2로 만들 수 있는
		1자리 3의 배수: 0
		2자리 3의 배수:
			1로 시작: 12
			2로 시작: 21
		3자리 3의 배수:
			1로 시작: 102, 111, 120
			2로 시작: 201, 210, 222
		4 자리 3의 배수:
			1로 시작: 1011, 1101, 1110, 1002, 1020, 1200, 1122, 1212, 1221
			2로 시작: 2001, 2010, 2100, 2022, 2202, 2220, 2112, 2121, 2211

		점화식: dp[n] = dp[n - 1] * 3
	*/

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] * 3) % 1000000009
	}

	fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
