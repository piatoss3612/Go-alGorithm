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

	N int
)

const MOD = 1000000

// 난이도: Gold 5
// 메모리: 928KB
// 시간: 2180ms
// 분류: 수학, 정수론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
}

func Solve() {
	ans := 0

	// i+1부터 N까지 i로 나누어 떨어지는 수의 개수와 i를 곱한 값을 누적해서 더한다.
	for i := 2; i <= N-1; i++ {
		ans += (N/i - 1) * i
		ans %= MOD
	}

	fmt.Fprintln(writer, ans)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
