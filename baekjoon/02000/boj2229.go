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
	input := make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if input[i] > input[j] {
				dp[i] = getMax(dp[i], input[i]-input[j]+dp[j-1])
			} else {
				dp[i] = getMax(dp[i], input[j]-input[i]+dp[j-1])
			}
		}
	}
	/*
		예제 입력:
		10
		2 5 7 1 3 4 8 6 9 3

		예제 프로세스:
		i = 1: [0 0 0 0 0 0 0 0 0 0 0]
		i = 2: [0 0 3 0 0 0 0 0 0 0 0]
		i = 3: [0 0 3 5 0 0 0 0 0 0 0]
		i = 4: [0 0 3 5 9 0 0 0 0 0 0]
		i = 5: [0 0 3 5 9 9 0 0 0 0 0]
		i = 6: [0 0 3 5 9 9 10 0 0 0 0]
		i = 7: [0 0 3 5 9 9 10 14 0 0 0]
		i = 8: [0 0 3 5 9 9 10 14 14 0 0]
		i = 9: [0 0 3 5 9 9 10 14 14 17 0]
		i = 10: [0 0 3 5 9 9 10 14 14 17 20]

		예제 출력:
		20
	*/

	fmt.Fprintln(writer, dp[n])
}

func getMax(a, b int) int {
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
