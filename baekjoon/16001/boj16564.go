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
	N, K    int
	team    []int
)

// 난이도: Silver 1
// 메모리: 3256KB
// 시간: 28ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	team = make([]int, N+1)
	for i := 1; i <= N; i++ {
		team[i] = scanInt()
	}
}

func Solve() {
	l, h := 1, 2000000000 // 목표 레벨 하한, 목표 레벨 상한
	ans := 0
	for l <= h {
		m := (l + h) / 2 // 목표 레벨
		sum := 0         // 목표 레벨 m을 달성하기 위해 팀원들이 올려야 하는 레벨의 합

		for i := 1; i <= N; i++ {
			sum += max(0, m-team[i])
			// 올릴 수 있는 레벨의 범위를 넘어선 경우는 반복문 종료
			if sum > K {
				break
			}
		}

		if sum <= K {
			// 올려야 하는 레벨의 합을 K 이하로 조절하여 모든 팀원이 목표 레벨을 달성할 수 있는 경우
			ans = max(ans, m)
			l = m + 1
		} else {
			h = m - 1
		}
	}

	fmt.Fprintln(writer, ans)
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
