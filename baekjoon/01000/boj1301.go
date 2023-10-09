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
	N       int    // 구슬의 종류
	beads   [6]int // 각 종류의 구슬의 개수

	// dp[i][j][a][b][c][d][e]: 전전에 i번 구슬을 선택하고 전에 j번 구슬을 선택한 중에
	// 구슬이 종류별로 각각 a, b, c, d, e개 남았을 때, 연속된 3개의 구슬의 색이 모두 다르게 목걸이를 만드는 경우의 수
	dp [6][6][11][11][11][11][11]int
)

// 목걸이를 만드는 경우의 수가 0이 될 수 있으므로
// 재귀 함수 내부에서 이미 탐색한 구간인지를 판별하려면 dp값을 -1로 초기화 해야 한다
func init() {
	for a := 0; a <= 5; a++ {
		for b := 0; b <= 5; b++ {
			for c := 0; c <= 10; c++ {
				for d := 0; d <= 10; d++ {
					for e := 0; e <= 10; e++ {
						for f := 0; f <= 10; f++ {
							for g := 0; g <= 10; g++ {
								dp[a][b][c][d][e][f][g] = -1
							}
						}
					}
				}
			}
		}
	}
}

// 난이도: Gold 3
// 메모리: 46200KB
// 시간: 56ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	for i := 1; i <= N; i++ {
		beads[i] = scanInt()
	}
}

func Solve() {
	ans := rec(0, 0)
	fmt.Fprintln(writer, ans)
}

func rec(prevprev, prev int) int {
	// 기저 사례: 남은 구슬의 개수가 0인 경우는 조건에 맞게 목걸이를 완성한 것이므로 1을 반환
	if beads[1]+beads[2]+beads[3]+beads[4]+beads[5] == 0 {
		return 1
	}
	ret := &dp[prevprev][prev][beads[1]][beads[2]][beads[3]][beads[4]][beads[5]]
	if *ret != -1 {
		return *ret
	}

	*ret = 0
	for i := 1; i <= N; i++ {
		// 연속하는 3개의 구슬의 색이 모든 다른 경우에만 목걸이를 이어서 만들 수 있다
		if beads[i] > 0 && i != prevprev && i != prev {
			beads[i]--
			*ret += rec(prev, i)
			beads[i]++
		}
	}
	return *ret
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
