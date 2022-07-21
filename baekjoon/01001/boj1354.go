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
	dp      map[int]int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, p, q, x, y := scanInt(), scanInt(), scanInt(), scanInt(), scanInt()

	if n == 0 {
		fmt.Fprintln(writer, 1)
		return
	}

	dp = make(map[int]int)
	dp[0] = 1

	solve(n, p, q, x, y)
	fmt.Fprintln(writer, dp[n])
}

// 메모리: 628484KB
// 시간: 3076ms
func solve(n, p, q, x, y int) int {
	if dp[n] > 0 {
		return dp[n]
	}

	var leftValue, rightValue int

	leftIdx := n/p - x
	if leftIdx <= 0 {
		leftValue = dp[0]
	} else {
		if dp[leftIdx] > 0 {
			leftValue = dp[leftIdx]
		} else {
			leftValue = solve(leftIdx, p, q, x, y)
		}
	}

	rightIdx := n/q - y
	if rightIdx <= 0 {
		rightValue = dp[0]
	} else {
		if dp[rightIdx] > 0 {
			rightValue = dp[rightIdx]
		} else {
			rightValue = solve(rightIdx, p, q, x, y)
		}
	}

	dp[n] = leftValue + rightValue
	return dp[n]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
