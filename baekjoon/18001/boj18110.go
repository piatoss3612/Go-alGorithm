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

	N   int
	arr []int
)

// 난이도: Silver 4
// 메모리: 3948KB
// 시간: 52ms
// 분류: 수학, 구현, 정렬
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
	sort.Ints(arr)
}

func Solve() {
	// N이 0이면 0 출력 (나누기 0 방지)
	if N == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	// 15%에서 반올림
	cutoff := func() int {
		temp := float64(N) * 0.15
		return int(temp + 0.5)
	}()

	// 앞에서 15%, 뒤에서 15%를 제외한 합
	sum := 0
	for i := cutoff; i < N-cutoff; i++ {
		sum += arr[i]
	}

	// 평균을 구하고 반올림
	avg := func() int {
		temp := float64(sum) / float64(N-2*cutoff)
		return int(temp + 0.5)
	}()

	fmt.Fprintln(writer, avg)
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
