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

	N  int
	dp [1001][3][3][2]int 
	// dp[here][prev][prev2][appeared]: here번째 지점에서 prev, prev2를 지나고 높이가 2인 선인장이 등장했는지 (appeared: 0 or 1) 여부에 따라 맵을 만들 수 있는 경우의 수
)

const MOD = 1000000007

// 난이도: Gold 4
// 메모리: 1392KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
}

func Solve() {
	fmt.Fprintln(writer, rec(1, 0, 0, 0))
}

func rec(here, prev, prev2, appeared int) int {
	// N번째 지점까지 왔을 때
	if here == N {
		// 높이가 2인 선인장이 최소 하나라도 등장했는지 여부에 따라 경우의 수를 반환
		if appeared == 0 {
			return 0
		}
		return 1
	}

	ret := &dp[here][prev][prev2][appeared]
	if *ret != 0 {
		return *ret
	}

	// 이전 지점에 선인장이 없었을 때
	if prev == 0 {
		*ret = (*ret + rec(here+1, 0, prev, appeared)) % MOD // 다음 지점에도 선인장이 없는 경우
		*ret = (*ret + rec(here+1, 1, prev, appeared)) % MOD // 다음 지점에 높이가 1인 선인장이 있는 경우
		*ret = (*ret + rec(here+1, 2, prev, 1)) % MOD // 다음 지점에 높이가 2인 선인장이 있는 경우
	}

	// 이전 지점에 높이가 1인 선인장이 있었을 때
	if prev == 1 {
		// 이전 이전 지점에 선인장이 없었을 때
		if prev2 == 0 {
			*ret = (*ret + rec(here+1, 0, prev, appeared)) % MOD // 다음 지점에 선인장이 없는 경우
			*ret = (*ret + rec(here+1, 1, prev, appeared)) % MOD // 다음 지점에 높이가 1인 선인장이 있는 경우
			*ret = (*ret + rec(here+1, 2, prev, 1)) % MOD // 다음 지점에 높이가 2인 선인장이 있는 경우
		} else {
			*ret = (*ret + rec(here+1, 0, prev, appeared)) % MOD // 선인장 2개를 뛰어넘었으므로 다음 지점은 선인장이 없는 경우만 가능
		}
	}

	// 이전 지점에 높이가 2인 선인장이 있었을 때
	if prev == 2 {
		// 이전 이전 지점에 선인장이 없었을 때
		if prev2 == 0 {
			*ret = (*ret + rec(here+1, 0, prev, appeared)) % MOD // 다음 지점에 선인장이 없는 경우
			*ret = (*ret + rec(here+1, 1, prev, appeared)) % MOD // 다음 지점에 높이가 1인 선인장이 있는 경우
			// 다음 지점에 높이가 2인 선인장이 있는 경우는 연속하는 선인장의 높이의 합이 4이상이 되므로 불가능
		} else {
			*ret = (*ret + rec(here+1, 0, prev, appeared)) % MOD // 선인장 2개를 뛰어넘었으므로 다음 지점은 선인장이 없는 경우만 가능
		}
	}

	return *ret
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
