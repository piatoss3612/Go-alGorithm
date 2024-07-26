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
	arr     []int
	op      [4]int
)

// 14888번: 연산자 끼워넣기
// hhttps://www.acmicpc.net/problem/14888
// 난이도: 실버 1
// 메모리: 864 KB
// 시간: 4 ms
// 분류: 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	arr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}
	for i := 0; i < 4; i++ {
		op[i] = scanInt()
	}
}

var minGlobal = 1000000000
var maxGlobal = -1000000000

func Solve() {
	dfs(arr[0], 1, 1000000000, -1000000000)
	fmt.Fprintln(writer, maxGlobal)
	fmt.Fprintln(writer, minGlobal)
}

func dfs(sum, idx int, min, max int) {
	if idx == N {
		if sum < minGlobal {
			minGlobal = sum
		}
		if sum > maxGlobal {
			maxGlobal = sum
		}
		return
	}

	for i := 0; i < 4; i++ {
		if op[i] == 0 {
			continue
		}

		op[i]--

		var nextSum int

		switch i {
		case 0:
			nextSum = sum + arr[idx]
		case 1:
			nextSum = sum - arr[idx]
		case 2:
			nextSum = sum * arr[idx]
		case 3:
			nextSum = sum / arr[idx]
		}

		dfs(nextSum, idx+1, min, max)

		op[i]++
	}
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
