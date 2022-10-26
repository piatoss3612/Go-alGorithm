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
	dp      [100][100][100]int // dp[i][j][k]: i번째 줄의 j번 타일에 기웅이가, k번 타일에 민수가 서있을 때 방을 탈출할 수 있는 경우의 수
)

const MOD = 10007

// 난이도: Gold 4
// 메모리: 9340KB
// 시간: 16ms
// 분류: 다이나믹 프로그래밍
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
	ans := rec(1, 1, 1)
	fmt.Fprintln(writer, ans)
}

func rec(row, kiwoong, minsu int) int {
	// 기저 사례: 방을 무사히 탈출한 경우
	if row == N {
		return 1
	}

	ret := &dp[row][kiwoong][minsu]
	if *ret != 0 {
		return *ret
	}

	// 기웅이와 민수가 동시에 같은 방향으로 이동하는 경우 2가지
	if kiwoong != minsu {
		*ret += rec(row+1, kiwoong, minsu)
		*ret += rec(row+1, kiwoong+1, minsu+1)
	}

	// 민수만 대각선으로 이동하는 경우
	if kiwoong != minsu+1 {
		*ret += rec(row+1, kiwoong, minsu+1)
	}

	// 기웅이만 대각선으로 이동하는 경우
	if kiwoong+1 != minsu {
		*ret += rec(row+1, kiwoong+1, minsu)
	}

	*ret %= MOD // 모듈러 연산

	return *ret
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
