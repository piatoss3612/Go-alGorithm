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

	N, A int
	arr  []int
)

// 난이도: Gold 5
// 메모리: 916KB
// 시간: 4ms
// 분류: 그리디 알고리즘, 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	arr = make([]int, N+1)
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
	A = scanInt()
}

func Solve() {
	for i := 1; i < N; i++ {
		if A == 0 {
			break
		}

		max := arr[i]
		pos := i

		for j := i + 1; j <= N; j++ {
			if j-i > A {
				break
			}

			if max < arr[j] {
				max = arr[j]
				pos = j
			}
		}

		for j := pos; j > i; j-- {
			arr[j] = arr[j-1]
		}
		arr[i] = max
		A -= pos - i
	}

	for i := 1; i <= N; i++ {
		fmt.Fprintf(writer, "%d ", arr[i])
	}
	fmt.Fprintln(writer)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
