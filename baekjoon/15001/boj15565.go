package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N, K     int
	puppet   []int  // N개의 인형들
	included [3]int // 집합에 포함된 각 인형의 개수
)

const INF = 987654321

// 난이도: Silver 1
// 메모리: 9288KB
// 시간: 56ms
// 분류: 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, K = scanInt(), scanInt()
	puppet = make([]int, N+1)
	for i := 1; i <= N; i++ {
		puppet[i] = scanInt()
	}
}

func Solve() {
	l, r := 1, 1
	ans := INF
	for r <= N {
		if included[1] < K {
			included[puppet[r]]++
			r++
		}

		for l < r {
			if included[1] >= K {
				ans = min(ans, r-l)
				included[puppet[l]]--
				l++
			} else {
				break
			}
		}
	}

	if ans == INF {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
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
