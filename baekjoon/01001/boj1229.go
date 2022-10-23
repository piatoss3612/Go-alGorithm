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
	N       int
	hex     [708]int     // 육각수
	dp      [1000001]int // dp[N]: 육각수를 구성하는 점의 개수의 합이 N이 되는 육각수 개수의 최솟값
)

func init() {
	hex[1] = 1
	// 입력값 N의 최댓값 1백만보다 작은 육각수는 707번째 육각수이며 점의 개수는 998991개 이다
	for i := 2; i <= 707; i++ {
		// 육각수 점화식
		hex[i] = hex[i-1] + hex[1] + 4*(i-1)
	}

	// 최솟값 비교를 위해 dp값 초기화
	for i := 1; i <= 1000000; i++ {
		dp[i] = 987654321
	}
}

// 난이도: Gold 4
// 메모리: 8732KB
// 시간: 1456ms
// 분류: 다이나믹 프로그래밍, 런타임 전의 전처리
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N = scanInt()
}

func Solve() {
	for i := 1; i <= N; i++ {
		for j := 1; j <= 707; j++ {
			// j번째 육각수를 구성하는 점의 개수가 i보다 많은 경우
			if i-hex[j] < 0 {
				break
			}

			// 합이 i가 되는 육각수 개수의 최솟값 갱신
			dp[i] = min(dp[i], dp[i-hex[j]]+1)
		}
	}

	fmt.Fprintln(writer, dp[N])
}

func min(a, b int) int {
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
