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
	visited []bool
	arr2    []int
	ans     int
)

// 1239번: 차트
// hhttps://www.acmicpc.net/problem/1239
// 난이도: 골드 5
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
	visited = make([]bool, N)
	arr2 = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	bruteForce(0, 0)
	fmt.Fprintln(writer, ans)
}

func bruteForce(n int, idx int) {
	arr2[n] = arr[idx]
	visited[idx] = true

	defer func() {
		visited[idx] = false
		arr2[n] = 0
	}()

	if n == N-1 {
		l, r := 0, 0
		sum := 0
		cnt := 0
		for r < N {
			sum += arr2[r]

			for sum > 50 {
				sum -= arr2[l]
				l++
			}

			if sum == 50 {
				cnt++
			}

			r++
		}

		if N > 2 {
			i := 0

			for l < N {
				sum += arr2[i]

				for sum > 50 {
					sum -= arr2[l]
					l++
					if l == N {
						break
					}
				}

				if sum == 50 {
					cnt++
				}

				i++
			}
		}

		ans = max(ans, cnt/2)
		return
	}

	for i := 0; i < N; i++ {
		if !visited[i] {
			bruteForce(n+1, i)
		}
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
