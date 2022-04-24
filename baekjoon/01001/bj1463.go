package bj1463

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

  dp := make([]int, n + 1)

  for i := 2; i <= n; i++ {
    dp[i] = dp[i - 1] + 1
    if i % 3 == 0 {
      if dp[i] > dp[i / 3] + 1 {
        dp[i] = dp[i / 3] + 1
      }
    }
    if i % 2 == 0 {
      if dp[i] > dp[i / 2] + 1 {
        dp[i] = dp[i / 2] + 1
      }
    }
  }
  fmt.Fprintln(writer, dp[n])
}

func scanInt() int {
  scanner.Scan()
  n, _ := strconv.Atoi(scanner.Text())
  return n
}