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
	inp     []int
	N       int
)

// 메모리: 2468KB
// 시간: 24ms
// 분할 정복, 6549번 문제와 같은 풀이
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()

	// 예외 처리: N이 0인 경우 알바를 할 수 없다
	if N == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	inp = make([]int, N+1)

	for i := 1; i <= N; i++ {
		inp[i] = scanInt()
	}

	fmt.Fprintln(writer, DAC(1, N))
}

// 분할 정복
func DAC(left, right int) int {
	// 0. 1일치 알바를 하는 경우의 급여
	if left == right {
		return inp[left]
	}

	// 1. mid를 기준으로 왼쪽, 오른쪽 구간을 나누어 급여의 최대치 구하기
	mid := (left + right) / 2
	ret := max(DAC(left, mid), DAC(mid+1, right))

	// 2. mid를 기준으로 좌우로 분포되어 있는 경우의 급여의 최대치 구하기
	segL, segR := mid, mid+1
	workingDays := min(inp[segL], inp[segR]) // 급여는 일급이 가장 작은 때를 기준
	ret = max(ret, workingDays*2)

	for left < segL || segR < right {
		if segR < right && (segL == left || inp[segL-1] < inp[segR+1]) {
			segR += 1
			workingDays = min(workingDays, inp[segR])
		} else {
			segL -= 1
			workingDays = min(workingDays, inp[segL])
		}

		ret = max(ret, (segR-segL+1)*workingDays)
	}
	return ret
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
