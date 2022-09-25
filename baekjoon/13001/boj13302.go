package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N, M     int
	dp       [101][51]int // 행: 리조트를 방문하는 날, 열: 쿠폰 개수 (쿠폰은 아무리 많아봐야 40개 언저리)
	notVisit [101]bool    // 리조트를 방문하지 않는 경우
)

const INF = 987654321

// 메모리: 1028KB
// 시간: 8ms
// 다이나믹 프로그래밍
// dp를 방문 날만 염두하여 1차원 배열로 선언하면 올바른 결과를 얻을 수 없으니 주의
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	for i := 1; i <= M; i++ {
		notVisit[scanInt()] = true
	}

	ans := solve(1, 0) // 1일에 시작하여 쿠폰 0개를 가지고 리조트를 이용하기 위해 지불하는 최솟값 구하기
	fmt.Fprintln(writer, ans)
}

func solve(day, coupon int) int {
	// 기저 사례: day가 N을 초과한 경우
	if day > N {
		return 0
	}

	ret := &dp[day][coupon]

	if *ret != 0 {
		return *ret
	}

	// 리조트를 방문할 수 없는 날인 경우
	if notVisit[day] {
		*ret = solve(day+1, coupon) // 다음날로 이동
		return *ret
	}

	*ret = INF

	// 쿠폰을 사용할 수 있어서 사용하는 경우
	if coupon >= 3 {
		*ret = min(*ret, solve(day+1, coupon-3)) // 쿠폰을 사용하고 다음날로 이동한 경우의 최솟값과 비교
	}

	// 1일권을 구매하고 다음날로 이동한 경우의 최솟값과 비교
	*ret = min(*ret, solve(day+1, coupon)+10000)

	// 3일권을 구매하고 3일 뒤로 이동한 경우의 최솟값과 비교
	*ret = min(*ret, solve(day+3, coupon+1)+25000)

	// 5일권을 구매하고 5일 뒤로 이동한 경우의 최솟값과 비교
	*ret = min(*ret, solve(day+5, coupon+2)+37000)

	return *ret
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
