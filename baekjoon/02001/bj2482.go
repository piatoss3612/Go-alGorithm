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
	dp      [1001][1001]int
)

/*
손으로 dp를 작성하다 운좋게 발견한 점화식...

N 색상환에서 어떤 인접한 두 색도 동시에 선택하지 않고 고를 수 있는 색의 최대 개수는
N이 짝수인 경우: N/2
N이 홀수인 경우: int(N/2)

k가 1인 경우의 수는 반드시 N
N이 짝수이면서 k가 N/2인 경우의 수는 반드시 2번

그 외에는 점화식 dp[i][j] = dp[i-1][j] + dp[i-2][j-1]으로 계산한다
*/

/*
점화식 dp[i][j] = dp[i-1][j] + dp[i-2][j-1]는 왜 성립하는가?

dp[5][2]를 생각해보자.
dp[5][2] = dp[4][2] + dp[3][1] = 2 + 3 = 5이다.


dp[3][1]의 경우의 수 (x는 선택된 색상의 번호):

x 2 3
1 x 3
1 2 x

dp[4][2]의 경우의 수:

x 2 x 4
1 x 3 x


temp1 = dp[4][2]의 경우의 수에 5번째 색을 추가하는 경우의 수:

x 2 x 4 5
1 x 3 x 5

temp2 = dp[3][1]의 경우의 수에 4번째, 5번째 색을 추가하여 temp1과 중복되지 않는 경우의 수:

x 2 3 x 5
1 x 3 4 x
1 2 x 4 x

이처럼 경우의 수가 중복되지 않고 최적 부분구조가 성립한다.
따라서 점화식 dp[i][j] = dp[i-1][j] + dp[i-2][j-1]가 성립하는 것이다.
*/

func init() {
	for i := 1; i <= 1000; i++ {
		// dp[i][1]은 반드시 i개의 경우의 수를 가진다
		dp[i][1] = i

		// 점화식
		for j := 2; j <= i/2; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i-2][j-1]) % 1000000003
		}

		// i가 짝수인 경우, dp[i][i/2]는 반드시 2가 되어야 한다
		if i%2 == 0 {
			dp[i][i/2] = 2
		}
	}
}

// 메모리: 6860KB
// 시간: 8ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, k := scanInt(), scanInt()
	fmt.Fprintln(writer, dp[n][k])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
