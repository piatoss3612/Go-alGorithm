package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, K    int
)

// 메모리: 1004KB
// 시간: 12ms
// 그리디 알고리즘, 정렬
// 오늘은 예비군 이슈로 쉬어갑니다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()

	inp := make([]int, N)
	for i := 0; i < N; i++ {
		inp[i] = scanInt()
	}

	sort.Slice(inp, func(i, j int) bool {
		return inp[i] > inp[j]
	})

	sum := 0
	for i := 0; i < K; i++ {
		sum += inp[i]
	}

	fmt.Fprintln(writer, sum-(K*(K-1)/2))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
