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

// 메모리: 2148KB
// 시간: 16ms
// 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	inp = make([]int, N+1)
	for i := 1; i <= N; i++ {
		inp[i] = scanInt()
	}

	s, e := 1, N // 양끝에서 시작

	ans := 0

	for s <= e {
		temp := (e - s - 1) * min(inp[s], inp[e]) // s와 e 사이의 개발자의 수 * s,e 중 능력치의 최솟값
		ans = max(ans, temp)

		// 더 작은 값에 해당하는 부분이 이동
		if inp[s] < inp[e] {
			s++
		} else {
			e--
		}
	}

	fmt.Fprintln(writer, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
