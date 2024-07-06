package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N          int
	brightness [200001]int
	isOn       [200001]int
)

// 25634번:
// hhttps://www.acmicpc.net/problem/25634
// 난이도: 골드 5
// 메모리: 5588 KB
// 시간: 44 ms
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
		brightness[i] = scanInt()
	}
	for i := 1; i <= N; i++ {
		isOn[i] = scanInt()
	}
}

func Solve() {
	sum := 0                   // 현 상태에서 켜져있는 전구의 밝기 합
	change := make([]int, N+1) // i번째 전구의 상태가 바뀌었을 때의 밝기 변화량

	for i := 1; i <= N; i++ {
		if isOn[i] == 1 {
			sum += brightness[i]
			change[i] = -brightness[i]
		} else {
			change[i] = brightness[i]
		}
	}

	// 카데인 알고리즘
	maxEndingHere := change[1] // i번째 전구까지 고려했을 때, i번째 전구를 마지막으로 하는 연속된 부분 배열의 최대 합
	maxSoFar := change[1]      // 부분 배열의 합 중 최대값

	for i := 2; i <= N; i++ {
		// i번째 전구의 상태가 바뀌었을 때의 밝기 변화량을 고려
		// 시작점은 정확히 알 수 없으나, 연속된 부분 배열이므로 i번째 전구를 마지막으로 하는 것만 고려하면 된다.
		// 따라서, i번째 전구의 상태가 바뀌었을 때의 밝기 변화량과 i-1번째 전구까지 고려했을 때의 최대 합을 더한다.
		// 만약, i-1번째 전구까지 고려했을 때의 최대 합이 change[i]보다 작다면, i번째 전구를 시작점으로 하는 새로운 부분 배열을 시작한다.
		// 이는 i번째 전구의 상태가 바뀌었을 때의 밝기 변화량만 고려하는 것이 기존의 부분 배열의 합보다 크다는 의미
		if maxEndingHere+change[i] > change[i] {
			maxEndingHere += change[i]
		} else {
			maxEndingHere = change[i]
		}

		// 부분 배열의 합 중 최대값을 갱신
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}
	}

	fmt.Fprintln(writer, sum+maxSoFar)
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
