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
	dp      [3][1001][1001]int
	milk    [1001][1001]int
	N       int
)

// 메모리: 55028KB
// 시간: 200ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			milk[i][j] = scanInt()
		}
	}

	fmt.Fprintln(writer, drink(1, 1, 0))
}

// 재귀 함수
func drink(x, y, target int) int {
	// 기저 사례: 탐색 범위 이탈
	if x > N || y > N {
		return 0
	}

	ret := &dp[target][x][y]
	if *ret != 0 {
		return *ret
	}

	// 1. x,y 위치에 마시고자하는 종류의 우유가 배치된 경우
	if milk[x][y] == target {
		// 1-1. 우유를 마시지 않고 지나갈 경우
		notDrninkAndGo := max(drink(x, y+1, target), drink(x+1, y, target))
		// 1-2. 우유를 마시고 지나갈 경우
		drinkAndGo := max(drink(x, y+1, (target+1)%3), drink(x+1, y, (target+1)%3)) + 1
		*ret = max(notDrninkAndGo, drinkAndGo)
	} else {
		// 2. 마시고자하는 우유가 없는 경우
		*ret = max(drink(x, y+1, target), drink(x+1, y, target))
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
