package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	input := []string{""}
	input = append(input, strings.Split(scanString(), "")...)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if i == 1 {
			for j := i + 1; j <= n; j++ {
				if input[j] == "O" {
					dp[j] = getMin(dp[j], dp[i]+(j-i)*(j-i))
				}
			}
		} else {
			if dp[i] > 0 {
				if input[i] == "B" {
					for j := i + 1; j <= n; j++ {
						if input[j] == "O" {
							dp[j] = getMin(dp[j], dp[i]+(j-i)*(j-i))
						}
					}
				} else if input[i] == "O" {
					for j := i + 1; j <= n; j++ {
						if input[j] == "J" {
							dp[j] = getMin(dp[j], dp[i]+(j-i)*(j-i))
						}
					}
				} else {
					for j := i + 1; j <= n; j++ {
						if input[j] == "B" {
							dp[j] = getMin(dp[j], dp[i]+(j-i)*(j-i))
						}
					}
				}
			}
		}
	}
	if dp[n] == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, dp[n])
	}
}

func getMin(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
