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
	N         int
	happiness []int
	dp        [2001][2001]int
	// dp [2001][2001][2001]int 며칠이 지났는지 까지 따져보려면 런타임에 메모리 초과 오류 발생
)

// 난이도: Gold 3
// 메모리: 46368KB
// 시간: 44ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	happiness = make([]int, N+1)
	for i := 1; i <= N; i++ {
		happiness[i] = scanInt()
	}
}

func Solve() {
	ans := rec(1, N)
	fmt.Fprintln(writer, ans)
}

// left번째 여물부터 right번째 여물까지 남아있을 때 소의 행복도의 합의 최댓값을 구한다
func rec(left, right int) int {
	// N일 동안 N개의 여물을 모두 소비해야 하므로
	// 여물을 구매한 뒤 지난 날짜는 (N - 남은 여물의 개수)로 구할 수 있다
	day := N - (right - left)

	// 기저 사례: N번째 여물을 소에게 주는 날
	if left == right {
		return day * happiness[left]
	}

	ret := &dp[left][right]
	if *ret != 0 {
		return *ret
	}

	*ret = max(*ret, rec(left+1, right)+day*happiness[left])  // 왼쪽 여물을 소에게 먹이는 경우
	*ret = max(*ret, rec(left, right-1)+day*happiness[right]) // 오른쪽 여물을 소에게 먹이는 경우

	return *ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
