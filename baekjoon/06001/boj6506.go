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
	nums    []int
	dp      [][]int
)

// 난이도: Gold 5
// 메모리: 2032KB
// 시간: 760ms
// 분류: 다이나믹 프로그래밍
func main() {
	// defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for {
		Input()
		Solve()
	}
}

func Input() {
	N, K = scanInt(), scanInt()
	if N == 0 && K == 0 {
		writer.Flush()
		os.Exit(0)
	}

	nums = make([]int, N+1)
	dp = make([][]int, N+1)
	dp[0] = make([]int, K+1)

	for i := 1; i <= N; i++ {
		dp[i] = make([]int, K+1)
		nums[i] = scanInt()
	}
}

func Solve() {
	ans := 0
	for i := 1; i <= N-K+1; i++ {
		ans += rec(i, 1)
	}
	fmt.Fprintln(writer, ans)
}

func rec(index, length int) int {
	if length == K {
		return 1
	}

	ret := &dp[index][length]
	if *ret != 0 {
		return *ret
	}

	for i := index + 1; i <= N; i++ {
		if nums[i] > nums[index] {
			*ret += rec(i, length+1)
		}
	}

	return *ret
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
