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
	ddeok   [1001][10]bool // i번째 날에 들고 가는 만든 떡
	visited [1001][10]bool // i번째 날에 호랑이에게 어떤 떡을 줬는지 여부
	ans     []int          // 호랑이에게 떡을 주는 여러 가지 경우 중 하나
)

// 메모리: 1064KB
// 시간: 4ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()

	// 1. i번째 날에 들고 가는 떡의 종류 입력
	for i := 1; i <= N; i++ {
		m := scanInt()
		for j := 1; j <= m; j++ {
			ddeok[i][scanInt()] = true
		}
	}

	ans = make([]int, N+1)

	// 2. N일 동안 호랑이에게 잡아먹히지 않는 것이 가능한지 여부 확인
	possible := solve(0, 0)
	if possible {
		// 가능, N일 동안 호랑이에게 준 떡 종류를 출력
		for i := 1; i <= N; i++ {
			fmt.Fprintln(writer, ans[i])
		}
	} else {
		// 불가능
		fmt.Fprintln(writer, -1)
	}
}

func solve(day, prev int) bool {
	// 기저 사례: N일까지 무사히 살아남은 경우
	if day == N {
		return true
	}

	// day+1일에 호랑이에게 줄 수 있는 떡 종류 탐색
	for i := 1; i <= 9; i++ {
		next := ddeok[day+1][i]

		// 호랑이에게 전날과 다른 종류의 떡을 주고 다음날로 이동
		if next && prev != i && !visited[day+1][i] {
			visited[day+1][i] = true
			ans[day+1] = i
			if solve(day+1, i) {
				return true
			}
		}
	}

	// 호랑이에게 전날과 같은 종류의 떡을 줄 수 밖에 없는 경우
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
