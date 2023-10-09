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
	height  []int
	dec     []int // dec[i]: 첫번째 기둥에서부터 i번째 기둥까지 가장 긴 내리막 구간의 길이
	inc     []int // inc[i]: i번째 기둥부터 N번째 기둥까지 가장 긴 오르막 구간의 길이
)

// 난이도: Gold 4
// 메모리: 940KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	height = make([]int, N+1)
	dec = make([]int, N+1)
	inc = make([]int, N+1)
	for i := 1; i <= N; i++ {
		height[i] = scanInt()
	}
}

func Solve() {
	// 1. 첫번째 기둥에서부터 i번째 기둥까지 가장 긴 내리막 구간의 길이 구하기
	dec[1] = 1
	for i := 2; i <= N; i++ {
		dec[i] = 1
		for j := 1; j < i; j++ {
			if height[j] > height[i] {
				dec[i] = max(dec[i], dec[j]+1)
			}
		}
	}

	// 2. i번째 기둥부터 N번째 기둥까지 가장 긴 오르막 구간의 길이 구하기
	inc[N] = 1
	for i := N - 1; i >= 1; i-- {
		inc[i] = 1
		for j := N; j > i; j-- {
			if height[j] > height[i] {
				inc[i] = max(inc[i], inc[j]+1)
			}
		}
	}

	ans := 0

	// 가장 긴 하이라이트 구간은 (inc[i]+dec[i]-1)의 최댓값
	// i번째 기둥 하나가 중복되어 있으므로 1을 빼준다
	for i := 1; i <= N; i++ {
		ans = max(ans, inc[i]+dec[i]-1)
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
