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
	// dp: [남은 약의 시작 위치][남은 약의 마지막 위치]일 때 먹을 수 있는 약의 최대 개수
	//N일치 약이므로 날짜는 셀 필요 없음
	dp        [1501][1501]int
	medicines []byte
)

// 메모리: 7392KB
// 시간: 8ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	medicines = scanBytes()

	ans := solve(0, len(medicines)-1, 'B') // 시작위치 0, 마지막 위치 len(medicines)-1, 약은 아침부터 먹기 시작
	fmt.Fprintln(writer, ans)
}

func solve(left, right int, when byte) int {
	// 기저 사례: 모든 약을 먹은 경우
	if left > right {
		return 0
	}

	ret := &dp[left][right]
	if *ret != 0 {
		return *ret
	}

	// 약은 식후 30분 후에
	switch when {
	// 1. 아침약 먹기
	case 'B':
		if medicines[left] == when {
			*ret = max(*ret, solve(left+1, right, 'L')+1)
		}
		if medicines[right] == when {
			*ret = max(*ret, solve(left, right-1, 'L')+1)
		}
	// 2. 점심약 먹기
	case 'L':
		if medicines[left] == when {
			*ret = max(*ret, solve(left+1, right, 'D')+1)
		}
		if medicines[right] == when {
			*ret = max(*ret, solve(left, right-1, 'D')+1)
		}
	// 3. 저녁약 먹기
	case 'D':
		if medicines[left] == when {
			*ret = max(*ret, solve(left+1, right, 'B')+1)
		}
		if medicines[right] == when {
			*ret = max(*ret, solve(left, right-1, 'B')+1)
		}
	}

	return *ret // 최댓값 반환
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

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
