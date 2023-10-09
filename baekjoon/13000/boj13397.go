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
	N, M    int
	arr     []int
)

// 난이도: Gold 4
// 메모리: 996KB
// 시간: 8ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	arr = make([]int, N+1)
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	l, r := 0, 9999
	for l <= r {
		m := (l + r) / 2
		if ParametricSearch(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	fmt.Fprintln(writer, l) // 최댓값의 최솟값(lower bound) 출력
}

// 구간의 점수가 expected일 때 M개의 구간으로 나눌 수 있는지 여부를 확인한다
func ParametricSearch(expected int) bool {
	maxNum := arr[1]
	minNum := arr[1]
	cnt := 0

	for i := 2; i <= N; i++ {
		maxNum = max(maxNum, arr[i])
		minNum = min(minNum, arr[i])

		// 구간의 점수가 expected보다 큰 경우
		// 구간을 끊고 i부터 새로 시작
		if maxNum-minNum > expected {
			cnt += 1
			maxNum = arr[i]
			minNum = arr[i]
		}
	}

	cnt++ // 마지막에 생성한 구간을 포함

	// 구간의 개수가 M개 이하인 경우
	// 생성된 구간에서 다시 임의로 구간을 나누어 M개의 구간 개수를 맞추면 된다
	return cnt <= M
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
