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
	N, B     int
	recovery [3001]int          // 각 구간에서 잠을 잤을 때 피로 회복량
	dp       [3001][3001][2]int // dp[i][j][k]: i개의 구간을 선택한 상태에서 j번째 구간에서 시작하여 k(취침:0 또는 기상:1) 상태일 때 피로 회복량의 최댓값
)

const INF = -987654321

// 난이도: Gold 4
// 메모리: 141400KB
// 시간: 248ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, B = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		recovery[i] = scanInt()
	}
}

func Solve() {
	ans := rec(0, 1, 1)
	fmt.Fprintln(writer, ans)
}

func rec(choosen, section, awaken int) int {
	// 기저 사례1: B개의 구간을 이미 선택한 경우
	if choosen == B {
		return 0
	}

	// 기저 사례2: 전체 구간의 범위를 넘어간 경우
	if section > N {
		return INF
	}

	ret := &dp[choosen][section][awaken]
	if *ret != 0 {
		return *ret
	}

	*ret = INF

	if awaken == 1 {
		// 1. 동혁이가 깨어 있는 상태

		// 1-1. 계속 깨어 있기
		*ret = max(*ret, rec(choosen, section+1, 1))
		// 1-2. 낮잠 자기: 준비 구간
		*ret = max(*ret, rec(choosen+1, section+1, 0))
	} else {
		// 2. 동혁이가 낮잠을 자고 있는 상태

		// 2-1. 연속해서 낮잠 자기
		*ret = max(*ret, rec(choosen+1, section+1, 0)+recovery[section])
		// 2-2. 잠에서 깨어나기
		*ret = max(*ret, rec(choosen, section+1, 1))
	}

	return *ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
