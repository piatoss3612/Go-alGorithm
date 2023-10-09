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
	statue  [100001]int // 돌상의 방향
	dpLeft  [100001]int // 왼쪽을 바라보는 돌상의 개수에서 오른쪽을 바라보는 돌상의 개수를 빼는 경우
	dpRight [100001]int // 오른쪽을 바라보는 돌상의 개수에서 왼쪽을 바라보는 돌상의 개수를 빼는 경우
)

// 난이도: Gold 5
// 메모리: 3268KB
// 시간: 12ms
// 분류: 다이나믹 프로그래밍, 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		statue[i] = scanInt()
	}
}

func Solve() {
	l, r := 0, 0
	for i := 1; i <= N; i++ {
		switch statue[i] {
		// i번째 돌상이 왼쪽을 바라보고 있는 경우
		case 1:
			// 오른쪽을 기준으로 하는 구간은 깨달음의 양 1 감소
			dpRight[i] = dpRight[i-1] - 1

			// 왼쪽을 바라보는 돌상의 개수가 오른쪽을 바라보는 돌상의 개수보다 적어졌을 경우
			// 이전 구간을 계속해서 이어나가는 것은 더 이상 의미가 없으므로
			// 이전 구간을 끊어내고 i부터 연속되는 새로운 구간으로 시작
			if dpLeft[i-1] < 0 {
				dpLeft[i] = 1
			} else {
				dpLeft[i] = dpLeft[i-1] + 1
			}
			l = max(l, dpLeft[i]) // 왼쪽 기준 깨달음의 양의 최댓값 갱신

		// i번째 돌상이 오른쪽을 바라보고 있는 경우
		case 2:
			// 왼쪽을 기준으로 하는 구간은 깨달음의 양 1 감소
			dpLeft[i] = dpLeft[i-1] - 1

			// 오른쪽을 바라보는 돌상의 개수가 왼쪽을 바라보는 돌상의 개수보다 적어졌을 경우
			// 이전 구간을 끊어내고 i부터 연속되는 새로운 구간으로 시작
			if dpRight[i-1] < 0 {
				dpRight[i] = 1
			} else {
				dpRight[i] = dpRight[i-1] + 1
			}
			r = max(r, dpRight[i]) // 오른쪽 기준 깨달음의 양의 최댓값 갱신
		}
	}
	fmt.Fprintln(writer, max(l, r)) // 깨달음의 양의 최댓값 출력
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
