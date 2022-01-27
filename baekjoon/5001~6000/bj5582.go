package bj5582

import (
  "fmt"
  "os"
  "bufio"
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
  // 똑같은 코드인데 문자열을 []byte로 가져오면 틀렸다고 한다
  // 왜 그러는지는 몰라서 오래 헤맸다
  scanner.Scan()
  s1 := strings.Split(scanner.Text(), "")
  scanner.Scan()
  s2 := strings.Split(scanner.Text(), "")
  dp := make([][]int, len(s1) + 1)
  for i := 0; i <= len(s1); i++ {
	dp[i] = make([]int, len(s2) + 1)
  }
  result := 0
  for i := 1; i <= len(s1); i++ {
	for j := 1; j <= len(s2); j++ {
	  if s1[i - 1] == s2[j - 1] { // 비교 대상이 같은 경우
		// 대각선 방향에 있는 값을 가져와 1을 더한 값을 부여
		dp[i][j] = dp[i - 1][j - 1] + 1
		if dp[i][j] > result { // 최댓값 비교
		  result = dp[i][j]
		}
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