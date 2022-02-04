package bj5525-1

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

var (
  scanner = bufio.NewScanner(os.Stdin)
  writer  = bufio.NewWriter(os.Stdout)
)

func main() {
  defer writer.Flush()
  scanner.Split(bufio.ScanWords)
  n, m, s := scanInt(), scanInt(), scanBytes()
  
  dp := make([]int, m)
  result := 0

  for i := 2; i < m; i++ {
	if s[i] == 73 {
	  if s[i - 2] == 73 && s[i - 1] == 79 {
		dp[i] = dp[i - 2] + 1
	  }
	  if dp[i] == n {
		dp[i] -= 1
		result += 1
	  }
	}
  }
  fmt.Fprintln(writer, result)
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