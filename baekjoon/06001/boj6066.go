package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	N, H      int
	suppliers [101]Supplier
	dp        [55001]int
	// 최소의 상한이 50000이고 그 이상을 사면서 가격이 더 싸게 나오는 경우가 있을 수 있으므로
	// 상한을 50000+5000=55000 (5000은 하나의 패키지에 들어갈 수 있는 건초 무게의 최댓값)으로 설정
)

type Supplier struct {
	Amount int
	Cost   int
}

const INF = 987654321

// 난이도: Gold 4
// 메모리: 1332KB
// 시간: 12ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, H = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		suppliers[i] = Supplier{scanInt(), scanInt()}
	}
	// 최솟값을 찾기 위해 INF로 초기화
	for i := 1; i <= H+5000; i++ {
		dp[i] = INF
	}

	// 0만큼의 건초를 사는데 드는 비용은 0
	// dp[0] = 0
}

func Solve() {
	// dp[i] = i만큼의 건초를 사는데 드는 최소 비용
	for i := 1; i <= N; i++ {
		for j := suppliers[i].Amount; j <= H+5000; j++ {
			dp[j] = min(dp[j], dp[j-suppliers[i].Amount]+suppliers[i].Cost)
		}
	}

	ans := INF
	// H부터 H+5000까지의 비용 중 최소값을 찾는다
	for i := H; i <= H+5000; i++ {
		ans = min(ans, dp[i])
	}

	fmt.Fprintln(writer, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
