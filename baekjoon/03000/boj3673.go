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
	C       int
	D, N    int
	sum     int
	pmod    []int
)

// 난이도: Gold 3
// 메모리: 37552KB
// 시간: 124ms
// 분류: 누적 합
// 비슷한 문제: boj10986
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	C = scanInt()
	for i := 1; i <= C; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	D, N = scanInt(), scanInt()
	pmod = make([]int, D)
	sum = 0
	for i := 1; i <= N; i++ {
		sum += scanInt()
		pmod[sum%D] += 1
	}
}

func Solve() {
	ans := pmod[0]
	for i := 0; i < D; i++ {
		if pmod[i] >= 2 {
			ans += (pmod[i] * (pmod[i] - 1)) / 2
		}
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
