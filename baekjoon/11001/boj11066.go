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
	INF     = 1<<63 - 1
	n       int
	inp     []int // 입력값
	sum     []int // 구간합을 구하기 위해 입력값의 누적합을 저장
	dp      [][]int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 1; i <= t; i++ {
		// testCase1()
		testCase2()
	}
}

// 재귀적 동적 계획법
// 메모리: 11376KB
// 시간: 856ms
func testCase1() {
	n = scanInt()
	inp = make([]int, n+1)
	sum = make([]int, n+1)
	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
		sum[i] = sum[i-1] + inp[i]
	}

	dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	// 인접한 파일들 끼리만 합칠 수 있으므로 허프만 코드 알고리즘은 사용할 수 없다
	// 그러나 허프만 코드의 트리 형태로 파일들을 2개의 구간으로 반복해서 나누어 최솟값을 구할 수 있다

	res := INF
	for i := 1; i < n; i++ {
		// i를 기준으로
		// a: 1~i구간에서 임시파일을 만드는 비용 + 임시 파일의 비용
		// b: i+1~n 구간에서 임시파일을 만드는 비용 + 임시 파일의 비용
		// a+b의 최솟값을 찾는다
		res = min(res, rec(1, i)+rec(i+1, n))
	}

	fmt.Fprintln(writer, res)
}

func rec(a, b int) int {
	if a == b {
		return inp[a]
	}

	ret := &dp[a][b]
	if *ret != 0 {
		return *ret
	}

	*ret = INF
	preSum := sum[b] - sum[a-1] //a~b의 구간합 = b까지의 누적합 - a-1까지의 누적합 = 임시파일을 만드는 비용
	for i := a; i < b; i++ {
		// (a~i 구간의 최솟값 + i+1~b구간의 최솟값 + 임시파일을 만드는 비용)
		// 즉 임시파일의 비용 + 임시파일을 만드는 비용의 최솟값 찾기
		*ret = min(*ret, rec(a, i)+rec(i+1, b)+preSum)
	}

	return *ret
}

// 반복적 동적 계획법
// 메모리: 13052KB
// 시간: 604ms
// 참고: https://cocoon1787.tistory.com/317
func testCase2() {
	n = scanInt()
	inp = make([]int, n+1)
	sum = make([]int, n+1)
	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
		sum[i] = sum[i-1] + inp[i]
	}

	dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	var ret *int
	var preSum int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			ret = &dp[j][i+j]
			*ret = INF
			preSum = sum[i+j] - sum[j-1]
			for k := j; k < i+j; k++ {
				*ret = min(*ret, dp[j][k]+dp[k+1][i+j]+preSum)
			}
		}
	}
	fmt.Fprintln(writer, dp[1][n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
